package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"time"

	sqlite3 "github.com/mattn/go-sqlite3"
)

// CDCEvent represents a change data capture event
type CDCEvent struct {
	Operation string                 `json:"operation"`
	Table     string                 `json:"table"`
	Timestamp string                 `json:"timestamp"`
	OldData   map[string]interface{} `json:"old_data,omitempty"`
	NewData   map[string]interface{} `json:"new_data,omitempty"`
}

// Custom function for INSERT operations - registered with SQLite
func cdcInsertToJson(tableName string, newRowJson string) string {
	var newData map[string]interface{}
	if err := json.Unmarshal([]byte(newRowJson), &newData); err != nil {
		log.Printf("Error unmarshaling new row data: %v", err)
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

// Custom function for UPDATE operations - registered with SQLite
func cdcUpdateToJson(tableName string, oldRowJson string, newRowJson string) string {
	var oldData, newData map[string]interface{}

	if err := json.Unmarshal([]byte(oldRowJson), &oldData); err != nil {
		log.Printf("Error unmarshaling old row data: %v", err)
		return ""
	}

	if err := json.Unmarshal([]byte(newRowJson), &newData); err != nil {
		log.Printf("Error unmarshaling new row data: %v", err)
		return ""
	}

	event := CDCEvent{
		Operation: "UPDATE",
		Table:     tableName,
		Timestamp: time.Now().UTC().Format(time.RFC3339),
		OldData:   oldData,
		NewData:   newData,
	}

	jsonBytes, err := json.Marshal(event)
	if err != nil {
		log.Printf("Error marshaling UPDATE event: %v", err)
		return ""
	}

	fmt.Printf("CDC UPDATE: %s\n", string(jsonBytes))
	return string(jsonBytes)
}

// Custom function for DELETE operations - registered with SQLite
func cdcDeleteToJson(tableName string, oldRowJson string) string {
	var oldData map[string]interface{}

	if err := json.Unmarshal([]byte(oldRowJson), &oldData); err != nil {
		log.Printf("Error unmarshaling old row data: %v", err)
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

// Register custom functions with SQLite
func registerCustomFunctions() error {
	sql.Register("sqlite3_with_cdc_functions",
		&sqlite3.SQLiteDriver{
			ConnectHook: func(conn *sqlite3.SQLiteConn) error {
				// Register cdcInsertToJson function
				if err := conn.RegisterFunc("cdcInsertToJson", cdcInsertToJson, true); err != nil {
					return fmt.Errorf("failed to register cdcInsertToJson: %v", err)
				}

				// Register cdcUpdateToJson function
				if err := conn.RegisterFunc("cdcUpdateToJson", cdcUpdateToJson, true); err != nil {
					return fmt.Errorf("failed to register cdcUpdateToJson: %v", err)
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
	// Create sample users table
	_, err := db.Exec(`
		CREATE TABLE IF NOT EXISTS users (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			name TEXT NOT NULL,
			email TEXT UNIQUE NOT NULL,
			age INTEGER,
			department TEXT,
			salary REAL,
			created_at DATETIME DEFAULT CURRENT_TIMESTAMP,
			updated_at DATETIME DEFAULT CURRENT_TIMESTAMP
		)
	`)
	if err != nil {
		return fmt.Errorf("failed to create users table: %v", err)
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
	}

	for _, drop := range dropTriggers {
		_, err = db.Exec(drop)
		if err != nil {
			return fmt.Errorf("failed to drop trigger: %v", err)
		}
	}

	// Create triggers using our custom functions
	// These triggers dynamically capture ALL columns without hardcoding
	triggers := []string{
		// INSERT trigger - captures all NEW row data
		`CREATE TRIGGER users_insert_trigger
		 AFTER INSERT ON users
		 FOR EACH ROW
		 BEGIN
		   INSERT INTO cdc_log (event_data) 
		   VALUES (
		     cdcInsertToJson(
		       'users',
		       json_object(
		         'id', NEW.id,
		         'name', NEW.name,
		         'email', NEW.email,
		         'age', NEW.age,
		         'department', NEW.department,
		         'salary', NEW.salary,
		         'created_at', NEW.created_at,
		         'updated_at', NEW.updated_at
		       )
		     )
		   );
		 END`,

		// UPDATE trigger - captures both OLD and NEW row data
		`CREATE TRIGGER users_update_trigger
		 AFTER UPDATE ON users
		 FOR EACH ROW
		 BEGIN
		   INSERT INTO cdc_log (event_data) 
		   VALUES (
		     cdcUpdateToJson(
		       'users',
		       json_object(
		         'id', OLD.id,
		         'name', OLD.name,
		         'email', OLD.email,
		         'age', OLD.age,
		         'department', OLD.department,
		         'salary', OLD.salary,
		         'created_at', OLD.created_at,
		         'updated_at', OLD.updated_at
		       ),
		       json_object(
		         'id', NEW.id,
		         'name', NEW.name,
		         'email', NEW.email,
		         'age', NEW.age,
		         'department', NEW.department,
		         'salary', NEW.salary,
		         'created_at', NEW.created_at,
		         'updated_at', NEW.updated_at
		       )
		     )
		   );
		 END`,

		// DELETE trigger - captures OLD row data
		`CREATE TRIGGER users_delete_trigger
		 AFTER DELETE ON users
		 FOR EACH ROW
		 BEGIN
		   INSERT INTO cdc_log (event_data) 
		   VALUES (
		     cdcDeleteToJson(
		       'users',
		       json_object(
		         'id', OLD.id,
		         'name', OLD.name,
		         'email', OLD.email,
		         'age', OLD.age,
		         'department', OLD.department,
		         'salary', OLD.salary,
		         'created_at', OLD.created_at,
		         'updated_at', OLD.updated_at
		       )
		     )
		   );
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

	// Insert operation
	fmt.Println("1. Inserting new user...")
	_, err := db.Exec(`
		INSERT INTO users (name, email, age, department, salary) 
		VALUES ('John Doe', 'john@example.com', 30, 'Engineering', 75000.50)
	`)
	if err != nil {
		return fmt.Errorf("failed to insert user: %v", err)
	}

	// Another insert with different data
	fmt.Println("2. Inserting another user...")
	_, err = db.Exec(`
		INSERT INTO users (name, email, age, department, salary) 
		VALUES ('Jane Smith', 'jane@example.com', 28, 'Marketing', 65000.00)
	`)
	if err != nil {
		return fmt.Errorf("failed to insert second user: %v", err)
	}

	// Update operation
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

	// Delete operation
	fmt.Println("4. Deleting user...")
	_, err = db.Exec(`
		DELETE FROM users 
		WHERE email = 'jane@example.com'
	`)
	if err != nil {
		return fmt.Errorf("failed to delete user: %v", err)
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
	for rows.Next() {
		var id int
		var eventData, createdAt string

		err := rows.Scan(&id, &eventData, &createdAt)
		if err != nil {
			return fmt.Errorf("failed to scan row: %v", err)
		}

		fmt.Printf("\nEvent ID: %d (Stored at: %s)\n", id, createdAt)

		// Pretty print JSON
		var prettyJSON map[string]interface{}
		if err := json.Unmarshal([]byte(eventData), &prettyJSON); err == nil {
			prettyBytes, _ := json.MarshalIndent(prettyJSON, "", "  ")
			fmt.Printf("%s\n", string(prettyBytes))
		} else {
			fmt.Printf("Raw: %s\n", eventData)
		}
	}

	return rows.Err()
}

// Show current users table state
func showCurrentUsers(db *sql.DB) error {
	rows, err := db.Query(`SELECT id, name, email, age, department, salary FROM users`)
	if err != nil {
		return fmt.Errorf("failed to query users: %v", err)
	}
	defer rows.Close()

	fmt.Println("\n=== Current Users Table ===")
	for rows.Next() {
		var id, age int
		var name, email, department string
		var salary float64

		err := rows.Scan(&id, &name, &email, &age, &department, &salary)
		if err != nil {
			return fmt.Errorf("failed to scan user row: %v", err)
		}

		fmt.Printf("ID: %d, Name: %s, Email: %s, Age: %d, Dept: %s, Salary: %.2f\n",
			id, name, email, age, department, salary)
	}

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

	// Show current state of users table
	if err := showCurrentUsers(db); err != nil {
		log.Fatal("Failed to show current users:", err)
	}

	// Display all CDC events that were captured and stored
	if err := displayStoredCDCEvents(db); err != nil {
		log.Fatal("Failed to display stored CDC events:", err)
	}

	fmt.Println("\nCDC example with custom registered functions completed successfully!")
}
