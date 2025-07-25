package autocdc

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log/slog"
	"strings"
	"text/template"

	"github.com/mattn/go-sqlite3"
)

// AutoCapture captures changes in sqlite database tables.

type TableSchema struct {
	Type     string  `json:"type" db:"type,omitempty"`
	Name     string  `json:"name" db:"name"`
	TblName  string  `json:"tbl_name" db:"tbl_name"`
	RootPage int64   `json:"rootpage" db:"rootpage"`
	Sql      *string `json:"sql" db:"sql"`
}

type CaptureOptions struct {
	Logger                 slog.Logger
	CustomSqliteDriverName string
}

type AutoCapture struct {
	db     *sql.DB
	logger slog.Logger
}

func NewAutoCapture(logger slog.Logger) *AutoCapture {
	return &AutoCapture{
		db:     nil,
		logger: logger,
	}
}

func (ac *AutoCapture) SetDB(db *sql.DB) {
	ac.db = db
}

func (ac *AutoCapture) RegisterCallback() func(conn *sqlite3.SQLiteConn) error {
	return func(conn *sqlite3.SQLiteConn) error {
		ac.logger.Info("Registering capture callback")
		err := conn.RegisterFunc("cdc_row_to_json", ac.cdcRowToJson, true)
		if err != nil {
			return err
		}

		return nil
	}
}

// POST regsiter

var triggerTemplate = template.Must(template.New("cdc").Parse(`

	CREATE TABLE IF NOT EXISTS zcdc_log_{{.TableName}} (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			event_data TEXT NOT NULL,
			event_type TEXT NOT NULL,
			row_id INTEGER NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
	);


	CREATE TRIGGER {{.TableName}}_insert_trigger
		AFTER INSERT ON {{.TableName}}
		FOR EACH ROW
		BEGIN
			INSERT INTO zcdc_log_{{.TableName}} (event_data, event_type, row_id) 
		   VALUES (cdc_row_to_json('{{.TableName}}', 1, NEW.id), 'insert', NEW.id);
	END

	CREATE TRIGGER {{.TableName}}_update_trigger
		AFTER UPDATE ON {{.TableName}}
		FOR EACH ROW
		BEGIN
			INSERT INTO zcdc_log_{{.TableName}} (event_data, event_type, row_id) 
		   VALUES (cdc_row_to_json('{{.TableName}}', 2, NEW.id), 'update', NEW.id);
	END

	CREATE TRIGGER {{.TableName}}_delete_trigger
		AFTER DELETE ON {{.TableName}}
		FOR EACH ROW
		BEGIN
			INSERT INTO zcdc_log_{{.TableName}} (event_data, event_type, row_id) 
		   VALUES (cdc_row_to_json('{{.TableName}}', 3, OLD.id), 'delete', OLD.id);
	END

	`))

func (ac *AutoCapture) Init() error {

	tables, err := ac.getTableSchemas()
	if err != nil {
		ac.logger.Error("Failed to get table schemas", "error", err)
		return err
	}

	for _, table := range tables {
		ac.logger.Info("Found table schema", "table", table.Name)

		if strings.HasPrefix(table.TblName, "sqlite_") {
			ac.logger.Debug("Skipping sqlite internal table", "table", table.TblName)
			continue
		}

		if strings.HasPrefix(table.TblName, "zcdc_") {
			ac.logger.Debug("Skipping existing internal cdc table", "table", table.TblName)
			continue
		}

		if strings.HasPrefix(table.TblName, "zbuddy_") {
			ac.logger.Debug("Skipping buddy table", "table", table.TblName)
			continue
		}

		// Register hooks for each table

	}

	return nil
}

// main cdc functions to handle inserts, updates, and deletes

func (ac *AutoCapture) cdcRowToJson(tableName string, etype int8, id int64) (string, error) {
	ac.logger.Info("Capturing %s for table: %s, id: %d", etype, tableName, id)

	if etype == 3 {
		return "{}", nil
	}

	rowData, err := ac.queryRowData(tableName, id)
	if err != nil {
		return "", fmt.Errorf("failed to query row data: %v", err)
	}

	jsonData, err := json.Marshal(rowData)
	if err != nil {
		return "", fmt.Errorf("failed to marshal row data: %v", err)
	}

	return string(jsonData), nil
}

// private

func (ac *AutoCapture) getTableSchemas() ([]TableSchema, error) {
	rows, err := ac.db.Query("SELECT type, name, tbl_name, rootpage, sql FROM sqlite_master WHERE type IN ('table')")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var schemas []TableSchema
	for rows.Next() {
		var schema TableSchema
		if err := rows.Scan(&schema.Type, &schema.Name, &schema.TblName, &schema.RootPage, &schema.Sql); err != nil {
			return nil, err
		}
		schemas = append(schemas, schema)
	}

	if err := rows.Err(); err != nil {
		return nil, err
	}

	return schemas, nil
}

func (ac *AutoCapture) queryRowData(tableName string, id int64) (map[string]any, error) {
	selectQuery := fmt.Sprintf("SELECT * FROM %s WHERE id = ?", tableName)
	rows, err := ac.db.Query(selectQuery, id)
	if err != nil {
		return nil, fmt.Errorf("failed to query row: %v", err)
	}
	defer rows.Close()

	// Get column names from the result set
	columns, err := rows.Columns()
	if err != nil {
		return nil, fmt.Errorf("failed to get columns: %v", err)
	}

	// Check if we have a row
	if !rows.Next() {
		return nil, fmt.Errorf("no row found with id %d", id)
	}

	// Prepare interface slice for scanning
	values := make([]any, len(columns))
	valuePtrs := make([]any, len(columns))
	for i := range values {
		valuePtrs[i] = &values[i]
	}

	// Scan the row
	err = rows.Scan(valuePtrs...)
	if err != nil {
		return nil, fmt.Errorf("failed to scan row: %v", err)
	}

	// Convert to map
	result := make(map[string]any)
	for i, col := range columns {
		result[col] = values[i]
	}

	return result, nil
}
