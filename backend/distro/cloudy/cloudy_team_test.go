package cloudy

import (
	"database/sql"
	"os"
	"path/filepath"
	"testing"

	_ "github.com/mattn/go-sqlite3"
)

func TestTeamStructure(t *testing.T) {
	dir, err := os.MkdirTemp("", "cloudy_test_*")
	if err != nil {
		t.Fatalf("temp dir: %v", err)
	}
	defer os.RemoveAll(dir)

	dbPath := filepath.Join(dir, "test.db")
	sqlDB, err := sql.Open("sqlite3", dbPath)
	if err != nil {
		t.Fatalf("open sqlite: %v", err)
	}
	defer sqlDB.Close()

	store, err := newStore(sqlDB)
	if err != nil {
		t.Fatalf("new store: %v", err)
	}

	if err := store.migrate(); err != nil {
		t.Fatalf("migrate: %v", err)
	}

	// Verify all 4 tables exist
	tables := []string{"CloudyUsers", "CloudyTeams", "CloudyTeamMembers", "CloudyTeamPotatoInstances"}
	for _, tbl := range tables {
		var count int
		err := sqlDB.QueryRow("SELECT COUNT(*) FROM sqlite_master WHERE type='table' AND name=?", tbl).Scan(&count)
		if err != nil {
			t.Fatalf("check table %s: %v", tbl, err)
		}
		if count == 0 {
			t.Fatalf("expected table %s to exist", tbl)
		}
	}

	// Insert user with team and instance
	user, team, inst, err := store.insertUserWithTeamAndInstance("Alice Smith", "alice@example.com", "secret_hash", "acme-corp", UTypeNormal, true)
	if err != nil {
		t.Fatalf("insert user with team: %v", err)
	}

	if user.ID == 0 || user.Email != "alice@example.com" {
		t.Fatalf("unexpected user: %+v", user)
	}
	if team.ID == 0 || team.OwnerID != user.ID || team.Name != "acme-corp" {
		t.Fatalf("unexpected team: %+v", team)
	}
	if inst.ID == 0 || inst.TeamID != team.ID || inst.Slug != "acme-corp" {
		t.Fatalf("unexpected instance: %+v", inst)
	}

	// Lookup primary instance
	pInst, pTeam, err := store.getPrimaryInstanceForUser(user.ID)
	if err != nil {
		t.Fatalf("get primary instance: %v", err)
	}
	if pInst == nil || pInst.Slug != "acme-corp" {
		t.Fatalf("expected primary instance slug acme-corp, got %+v", pInst)
	}
	if pTeam == nil || pTeam.ID != team.ID {
		t.Fatalf("expected primary team id %d, got %+v", team.ID, pTeam)
	}
}
