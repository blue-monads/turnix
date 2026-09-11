package cloudy

import (
	"fmt"
	"io/fs"
	"log"
	"sort"
	"strings"
)

func (a *CloudyApp) Migrate() error {
	if err := a.initStore(); err != nil {
		return err
	}
	if err := a.store.migrate(); err != nil {
		return err
	}
	if err := a.tursoDB.Push(a.rootCtx); err != nil {
		return fmt.Errorf("push migrated schema: %w", err)
	}
	log.Println("cloudy migrations applied")
	return nil
}

func (a *CloudyApp) initStore() error {
	if a.store != nil {
		return nil
	}
	sqlDB, err := a.tursoDB.Connect(a.rootCtx)
	if err != nil {
		return err
	}
	store, err := newStore(sqlDB)
	if err != nil {
		return err
	}
	a.store = store
	return nil
}

func (s *Store) migrate() error {
	if err := s.ensureCloudyUsersTable(); err != nil {
		return err
	}

	entries, err := fs.ReadDir(migrationFS, "migrations")
	if err != nil {
		return err
	}
	names := make([]string, 0, len(entries))
	for _, e := range entries {
		if e.IsDir() || !strings.HasSuffix(e.Name(), ".sql") {
			continue
		}
		names = append(names, e.Name())
	}
	sort.Strings(names)

	if err := s.execScript(`CREATE TABLE IF NOT EXISTS schema_migrations (
		name TEXT PRIMARY KEY,
		applied_at TIMESTAMP NOT NULL DEFAULT CURRENT_TIMESTAMP
	)`); err != nil {
		return err
	}

	applied, err := s.appliedMigrations()
	if err != nil {
		return err
	}

	for _, name := range names {
		if applied[name] {
			continue
		}
		body, err := fs.ReadFile(migrationFS, "migrations/"+name)
		if err != nil {
			return err
		}
		log.Println("applying migration", name)
		if err := s.execScript(string(body)); err != nil {
			return fmt.Errorf("migration %s: %w", name, err)
		}
		if _, err := s.sql.Exec(`INSERT INTO schema_migrations (name) VALUES (?)`, name); err != nil {
			return err
		}
	}

	return s.ensureCloudyUserColumns()
}

func (s *Store) appliedMigrations() (map[string]bool, error) {
	rows, err := s.sql.Query(`SELECT name FROM schema_migrations`)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	out := map[string]bool{}
	for rows.Next() {
		var name string
		if err := rows.Scan(&name); err != nil {
			return nil, err
		}
		out[name] = true
	}
	return out, rows.Err()
}

func (s *Store) execScript(script string) error {
	for _, stmt := range splitSQL(script) {
		if _, err := s.sql.Exec(stmt); err != nil {
			return err
		}
	}
	return nil
}

func (s *Store) tableExists(name string) (bool, error) {
	var n int
	err := s.sql.QueryRow(
		`SELECT COUNT(*) FROM sqlite_master WHERE type = 'table' AND name = ? COLLATE NOCASE`,
		name,
	).Scan(&n)
	return n > 0, err
}

func (s *Store) columnSet(table string) (map[string]bool, error) {
	rows, err := s.sql.Query(fmt.Sprintf(`PRAGMA table_info("%s")`, table))
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	cols := map[string]bool{}
	for rows.Next() {
		var cid int
		var name, ctype string
		var notnull, pk int
		var dflt any
		if err := rows.Scan(&cid, &name, &ctype, &notnull, &dflt, &pk); err != nil {
			return nil, err
		}
		cols[strings.ToLower(name)] = true
	}
	return cols, rows.Err()
}

func (s *Store) ensureCloudyUsersTable() error {
	hasCloudy, err := s.tableExists("CloudyUsers")
	if err != nil {
		return err
	}
	if hasCloudy {
		return nil
	}

	hasUsers, err := s.tableExists("Users")
	if err != nil {
		return err
	}
	if !hasUsers {
		return nil
	}

	cols, err := s.columnSet("Users")
	if err != nil {
		return err
	}
	if !cols["fullname"] || !cols["tenant_key"] {
		return nil
	}

	log.Println("renaming Users to CloudyUsers")
	_, err = s.sql.Exec(`ALTER TABLE Users RENAME TO CloudyUsers`)
	return err
}

func (s *Store) ensureCloudyUserColumns() error {
	exists, err := s.tableExists("CloudyUsers")
	if err != nil {
		return err
	}
	if !exists {
		return fmt.Errorf("CloudyUsers table was not created")
	}

	cols, err := s.columnSet("CloudyUsers")
	if err != nil {
		return err
	}

	adds := []struct {
		name string
		def  string
	}{
		{"fullname", "TEXT NOT NULL DEFAULT ''"},
		{"email", "TEXT"},
		{"password", "TEXT"},
		{"tenant_key", "TEXT NOT NULL DEFAULT ''"},
		{"utype", "TEXT NOT NULL DEFAULT 'normal'"},
		{"pricing_tier", "TEXT NOT NULL DEFAULT 'free'"},
		{"is_verified", "INTEGER NOT NULL DEFAULT 0"},
		{"is_lazy_loaded", "INTEGER NOT NULL DEFAULT 1"},
		{"is_disabled", "INTEGER NOT NULL DEFAULT 0"},
		{"created_at", "TIMESTAMP DEFAULT CURRENT_TIMESTAMP"},
		{"updated_at", "TIMESTAMP DEFAULT CURRENT_TIMESTAMP"},
	}

	for _, col := range adds {
		if cols[col.name] {
			continue
		}
		log.Println("adding CloudyUsers." + col.name)
		if _, err := s.sql.Exec(fmt.Sprintf(`ALTER TABLE CloudyUsers ADD COLUMN %s %s`, col.name, col.def)); err != nil {
			return fmt.Errorf("add column %s: %w", col.name, err)
		}
	}

	if _, err := s.sql.Exec(`CREATE UNIQUE INDEX IF NOT EXISTS idx_cloudy_users_email ON CloudyUsers(email)`); err != nil {
		return err
	}
	if _, err := s.sql.Exec(`CREATE UNIQUE INDEX IF NOT EXISTS idx_cloudy_users_tenant_key ON CloudyUsers(tenant_key)`); err != nil {
		return err
	}
	return nil
}

func splitSQL(script string) []string {
	var stmts []string
	for _, part := range strings.Split(script, ";") {
		stmt := strings.TrimSpace(part)
		if stmt == "" || strings.HasPrefix(stmt, "--") {
			continue
		}
		stmts = append(stmts, stmt)
	}
	return stmts
}
