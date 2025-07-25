package main

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

	"github.com/mattn/go-sqlite3"
	_ "github.com/mattn/go-sqlite3"
)

var globalDB *sql.DB // Global reference to access DB from custom functions

// CDCEvent represents a change data capture event
type CDCEvent struct {
	Operation string                 `json:"operation"`
	Table     string                 `json:"table"`
	Timestamp string                 `json:"timestamp"`
	OldData   map[string]interface{} `json:"old_data,omitempty"`
	NewData   map[string]interface{} `json:"new_data,omitempty"`
}

// Helper function to query row data by ID
func queryRowData(tableName string, id int64) (map[string]interface{}, error) {
	if globalDB == nil {
		return nil, fmt.Errorf("database connection not available")
	}

	// First, get column names for the table
	columnQuery := fmt.Sprintf("PRAGMA table_info(%s)", tableName)
	rows, err := globalDB.Query(columnQuery)
	if err != nil {
		return nil, fmt.Errorf("failed to get table info: %v", err)
	}
	defer rows.Close()

	var columns []string
	for rows.Next() {
		var cid int
		var name, dataType string
		var notNull, pk int
		var defaultValue interface{}

		err := rows.Scan(&cid, &name, &dataType, &notNull, &defaultValue, &pk)
		if err != nil {
			return nil, fmt.Errorf("failed to scan column info: %v", err)
		}
		columns = append(columns, name)
	}

	if len(columns) == 0 {
		return nil, fmt.Errorf("no columns found for table %s", tableName)
	}

	// Build dynamic SELECT query
	selectQuery := fmt.Sprintf("SELECT * FROM %s WHERE id = ?", tableName)
	row := globalDB.QueryRow(selectQuery, id)

	// Prepare interface slice for scanning
	values := make([]interface{}, len(columns))
	valuePtrs := make([]interface{}, len(columns))
	for i := range values {
		valuePtrs[i] = &values[i]
	}

	err = row.Scan(valuePtrs...)
	if err != nil {
		return nil, fmt.Errorf("failed to scan row: %v", err)
	}

	// Convert to map
	result := make(map[string]interface{})
	for i, col := range columns {
		result[col] = values[i]
	}

	return result, nil
}

// Custom function for INSERT operations - only needs table name and ID
func cdcInsertToJson(tableName string, id int64) string {
	newData, err := queryRowData(tableName, id)
	if err != nil {
		log.Printf("Error querying new row data: %v", err)
		return ""
	}

	event := CDCEvent{
		Operation: "INSERT",
		Table:     tableName,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		NewData:   newData,
	}

	jsonBytes, err := json.Marshal(event)
	if err != nil {
		log.Printf("Error marshaling INSERT event: %v", err)
		return ""
	}

	fmt.Printf("CDC INSERT: %s\n", string(jsonBytes))
	return string(jsonBytes)
}

// Custom function for UPDATE operations - needs table name and ID
// For updates, we capture the NEW data (since trigger is AFTER UPDATE)
func cdcUpdateToJson(tableName string, id int64) string {
	// Note: Since this is AFTER UPDATE, we can only get the NEW data
	// If you need OLD data, you'd need to capture it before the update
	// or use a BEFORE trigger to store it temporarily
	newData, err := queryRowData(tableName, id)
	if err != nil {
		log.Printf("Error querying updated row data: %v", err)
		return ""
	}

	event := CDCEvent{
		Operation: "UPDATE",
		Table:     tableName,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		NewData:   newData,
		// OldData would need to be captured in a BEFORE trigger if needed
	}

	jsonBytes, err := json.Marshal(event)
	if err != nil {
		log.Printf("Error marshaling UPDATE event: %v", err)
		return ""
	}

	fmt.Printf("CDC UPDATE: %s\n", string(jsonBytes))
	return string(jsonBytes)
}

// For DELETE, we capture data before deletion using BEFORE trigger
func cdcDeleteToJson(tableName string, id int64) string {
	// Since this is a BEFORE DELETE trigger, we can still query the data
	oldData, err := queryRowData(tableName, id)
	if err != nil {
		log.Printf("Error querying row data before delete: %v", err)
		return ""
	}

	event := CDCEvent{
		Operation: "DELETE",
		Table:     tableName,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		OldData:   oldData,
	}

	jsonBytes, err := json.Marshal(event)
	if err != nil {
		log.Printf("Error marshaling DELETE event: %v", err)
		return ""
	}

	fmt.Printf("CDC DELETE: %s\n", string(jsonBytes))
	return string(jsonBytes)
}

// For UPDATE, we need to capture OLD data before the change
// This function will be called from BEFORE UPDATE trigger
func cdcUpdateBeforeToJson(tableName string, id int64) string {
	// Capture the OLD data before update
	oldData, err := queryRowData(tableName, id)
	if err != nil {
		log.Printf("Error querying old row data: %v", err)
		return ""
	}

	// Store old data temporarily (you might want to use a proper cache/storage)
	// For now, we'll just return it to be stored in a temp table
	jsonBytes, err := json.Marshal(oldData)
	if err != nil {
		log.Printf("Error marshaling old data: %v", err)
		return ""
	}

	return string(jsonBytes)
}

// Enhanced UPDATE function that captures both OLD and NEW data
func cdcUpdateAfterToJson(tableName string, id int64) string {
	// Get NEW data after update
	newData, err := queryRowData(tableName, id)
	if err != nil {
		log.Printf("Error querying updated row data: %v", err)
		return ""
	}

	// For now, we'll just capture NEW data
	// To get OLD data, we'd need a more complex setup with temp storage
	event := CDCEvent{
		Operation: "UPDATE",
		Table:     tableName,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		NewData:   newData,
		// OldData: would need temp storage between BEFORE and AFTER triggers
	}

	jsonBytes, err := json.Marshal(event)
	if err != nil {
		log.Printf("Error marshaling UPDATE event: %v", err)
		return ""
	}

	fmt.Printf("CDC UPDATE: %s\n", string(jsonBytes))
	return string(jsonBytes)
}

func init() {

	conn, err := globalDB.Conn(context.TODO())
	if err != nil {
		log.Printf("Error getting DB connection: %v", err)
		return
	}
	defer conn.Close()

}

// Register custom functions with SQLite
func registerCustomFunctions() error {
	sql.Register("sqlite3_with_cdc_functions",
		&sqlite3.SQLiteDriver{
			ConnectHook: func(conn *sqlite3.SQLiteConn) error {
				// Register cdcInsertToJson function
				if err := conn.RegisterFunc("cdcInsertToJson", cdcInsertToJson, true); err != nil {
					return fmt.Errorf("failed to register cdcInsertToJson: %v", err)
				}

				// Register cdcUpdateToJson function (simple version)
				if err := conn.RegisterFunc("cdcUpdateToJson", cdcUpdateToJson, true); err != nil {
					return fmt.Errorf("failed to register cdcUpdateToJson: %v", err)
				}

				// Register BEFORE UPDATE function for capturing old data
				if err := conn.RegisterFunc("cdcUpdateBeforeToJson", cdcUpdateBeforeToJson, true); err != nil {
					return fmt.Errorf("failed to register cdcUpdateBeforeToJson: %v", err)
				}

				// Register AFTER UPDATE function
				if err := conn.RegisterFunc("cdcUpdateAfterToJson", cdcUpdateAfterToJson, true); err != nil {
					return fmt.Errorf("failed to register cdcUpdateAfterToJson: %v", err)
				}

				// Register cdcDeleteToJson function
				if err := conn.RegisterFunc("cdcDeleteToJson", cdcDeleteToJson, true); err != nil {
					return fmt.Errorf("failed to register cdcDeleteToJson: %v", err)
				}

				fmt.Println("Custom CDC functions registered successfully")
				return nil
			},
		})
	return nil
}

// Create sample table and triggers using custom functions
func setupDatabase(db *sql.DB) error {
	// Store global reference for custom functions
	globalDB = db

	// Create sample users table
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			email TEXT UNIQUE NOT NULL,
			age INTEGER,
			department TEXT,
			salary REAL,
			status TEXT DEFAULT 'active',
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		return fmt.Errorf("failed to create users table: %v", err)
	}

	// Create products table to show it works with any table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS products (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			price REAL,
			category TEXT,
			in_stock BOOLEAN DEFAULT 1,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		return fmt.Errorf("failed to create products table: %v", err)
	}

	// Create CDC log table
	_, err = db.Exec(`
		CREATE TABLE IF NOT EXISTS cdc_log (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			event_data TEXT NOT NULL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		return fmt.Errorf("failed to create cdc_log table: %v", err)
	}

	// Drop existing triggers if they exist
	dropTriggers := []string{
		"DROP TRIGGER IF EXISTS users_insert_trigger",
		"DROP TRIGGER IF EXISTS users_update_trigger",
		"DROP TRIGGER IF EXISTS users_delete_trigger",
		"DROP TRIGGER IF EXISTS products_insert_trigger",
		"DROP TRIGGER IF EXISTS products_update_trigger",
		"DROP TRIGGER IF EXISTS products_delete_trigger",
	}

	for _, drop := range dropTriggers {
		_, err = db.Exec(drop)
		if err != nil {
			return fmt.Errorf("failed to drop trigger: %v", err)
		}
	}

	// Create simple triggers that only pass table name and ID
	triggers := []string{
		// Users table triggers - ONLY table name and ID!
		`CREATE TRIGGER users_insert_trigger
		 AFTER INSERT ON users
		 FOR EACH ROW
		 BEGIN
		   INSERT INTO cdc_log (event_data) 
		   VALUES (cdcInsertToJson('users', NEW.id));
		 END`,

		`CREATE TRIGGER users_update_trigger
		 AFTER UPDATE ON users
		 FOR EACH ROW
		 BEGIN
		   INSERT INTO cdc_log (event_data) 
		   VALUES (cdcUpdateAfterToJson('users', NEW.id));
		 END`,

		`CREATE TRIGGER users_delete_trigger
		 BEFORE DELETE ON users
		 FOR EACH ROW
		 BEGIN
		   INSERT INTO cdc_log (event_data) 
		   VALUES (cdcDeleteToJson('users', OLD.id));
		 END`,

		// Products table triggers - same pattern, only table name and ID!
		`CREATE TRIGGER products_insert_trigger
		 AFTER INSERT ON products
		 FOR EACH ROW
		 BEGIN
		   INSERT INTO cdc_log (event_data) 
		   VALUES (cdcInsertToJson('products', NEW.id));
		 END`,

		`CREATE TRIGGER products_update_trigger
		 AFTER UPDATE ON products
		 FOR EACH ROW
		 BEGIN
		   INSERT INTO cdc_log (event_data) 
		   VALUES (cdcUpdateAfterToJson('products', NEW.id));
		 END`,

		`CREATE TRIGGER products_delete_trigger
		 BEFORE DELETE ON products
		 FOR EACH ROW
		 BEGIN
		   INSERT INTO cdc_log (event_data) 
		   VALUES (cdcDeleteToJson('products', OLD.id));
		 END`,
	}

	for _, trigger := range triggers {
		_, err = db.Exec(trigger)
		if err != nil {
			return fmt.Errorf("failed to create trigger: %v", err)
		}
	}

	fmt.Println("Database tables and triggers created successfully")
	return nil
}

// Simulate database operations to test CDC
func simulateOperations(db *sql.DB) error {
	fmt.Println("\n=== Simulating Database Operations ===")

	// Insert operations
	fmt.Println("1. Inserting users...")
	_, err := db.Exec(`
		INSERT INTO users (name, email, age, department, salary, status) 
		VALUES ('John Doe', 'john@example.com', 30, 'Engineering', 75000.50, 'active')
	`)
	if err != nil {
		return fmt.Errorf("failed to insert user: %v", err)
	}

	_, err = db.Exec(`
		INSERT INTO users (name, email, age, department, salary, status) 
		VALUES ('Jane Smith', 'jane@example.com', 28, 'Marketing', 65000.00, 'active')
	`)
	if err != nil {
		return fmt.Errorf("failed to insert second user: %v", err)
	}

	// Insert products
	fmt.Println("2. Inserting products...")
	_, err = db.Exec(`
		INSERT INTO products (name, price, category, in_stock) 
		VALUES ('Laptop', 999.99, 'Electronics', 1)
	`)
	if err != nil {
		return fmt.Errorf("failed to insert product: %v", err)
	}

	// Update operations
	fmt.Println("3. Updating user...")
	_, err = db.Exec(`
		UPDATE users 
		SET name = 'John Smith', 
		    department = 'Senior Engineering',
		    salary = 85000.00,
		    updated_at = CURRENT_TIMESTAMP 
		WHERE email = 'john@example.com'
	`)
	if err != nil {
		return fmt.Errorf("failed to update user: %v", err)
	}

	fmt.Println("4. Updating product...")
	_, err = db.Exec(`
		UPDATE products 
		SET price = 899.99, in_stock = 0 
		WHERE name = 'Laptop'
	`)
	if err != nil {
		return fmt.Errorf("failed to update product: %v", err)
	}

	// Delete operations
	fmt.Println("5. Deleting user...")
	_, err = db.Exec(`
		DELETE FROM users 
		WHERE email = 'jane@example.com'
	`)
	if err != nil {
		return fmt.Errorf("failed to delete user: %v", err)
	}

	fmt.Println("6. Deleting product...")
	_, err = db.Exec(`
		DELETE FROM products 
		WHERE name = 'Laptop'
	`)
	if err != nil {
		return fmt.Errorf("failed to delete product: %v", err)
	}

	return nil
}

// Display stored CDC events
func displayStoredCDCEvents(db *sql.DB) error {
	rows, err := db.Query(`
		SELECT id, event_data, created_at 
		FROM cdc_log 
		ORDER BY created_at
	`)
	if err != nil {
		return fmt.Errorf("failed to query CDC log: %v", err)
	}
	defer rows.Close()

	fmt.Println("\n=== Stored CDC Events ===")
	eventCount := 0
	for rows.Next() {
		var id int
		var eventData, createdAt string

		err := rows.Scan(&id, &eventData, &createdAt)
		if err != nil {
			return fmt.Errorf("failed to scan row: %v", err)
		}

		eventCount++
		fmt.Printf("\nEvent #%d (ID: %d, Stored: %s)\n", eventCount, id, createdAt)

		// Pretty print JSON
		var prettyJSON map[string]interface{}
		if err := json.Unmarshal([]byte(eventData), &prettyJSON); err == nil {
			prettyBytes, _ := json.MarshalIndent(prettyJSON, "", "  ")
			fmt.Printf("%s\n", string(prettyBytes))
		} else {
			fmt.Printf("Raw: %s\n", eventData)
		}
	}

	fmt.Printf("\nTotal CDC events captured: %d\n", eventCount)
	return rows.Err()
}

func main() {
	// Register custom functions first
	if err := registerCustomFunctions(); err != nil {
		log.Fatal("Failed to register custom functions:", err)
	}

	// Open SQLite database with our custom driver
	db, err := sql.Open("sqlite3_with_cdc_functions", "cdc_example.db")
	if err != nil {
		log.Fatal("Failed to open database:", err)
	}
	defer db.Close()

	// Test connection
	if err := db.Ping(); err != nil {
		log.Fatal("Failed to ping database:", err)
	}

	fmt.Println("Connected to SQLite database with custom CDC functions")

	// Setup database tables and triggers
	if err := setupDatabase(db); err != nil {
		log.Fatal("Failed to setup database:", err)
	}

	// Simulate database operations (this will trigger our custom functions)
	if err := simulateOperations(db); err != nil {
		log.Fatal("Failed to simulate operations:", err)
	}

	// Display all CDC events that were captured and stored
	if err := displayStoredCDCEvents(db); err != nil {
		log.Fatal("Failed to display stored CDC events:", err)
	}

	fmt.Println("\nCDC example with dynamic data querying completed successfully!")
}
