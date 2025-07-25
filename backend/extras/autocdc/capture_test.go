package autocdc

import (
	"database/sql"
	"log/slog"
	"testing"

	"github.com/k0kubun/pp"
	_ "github.com/mattn/go-sqlite3"
)

func TestAutoCapture(t *testing.T) {

	t.Log("Performing test for AutoCapture")

	ac := NewAutoCapture(CaptureOptions{
		Logger: slog.Default(),
	})

	db, err := sql.Open("sqlite3", ":memory:")
	if err != nil {
		t.Fatal("Failed to open database:", err)
	}

	defer db.Close()

	_, err = db.Exec("CREATE TABLE test (id INTEGER PRIMARY KEY, name TEXT, age INTEGER, created_at DATETIME DEFAULT CURRENT_TIMESTAMP)")
	if err != nil {
		t.Fatal("Failed to create table:", err)
	}

	ac.SetDB(db)
	err = ac.Init()
	if err != nil {
		t.Fatal("Failed to initialize AutoCapture:", err)
	}

	_, err = db.Exec("INSERT INTO test (name, age) VALUES ('test1', 25)")
	if err != nil {
		t.Fatal("Failed to insert data:", err)
	}

	data, err := queryRowData(db, "test", 1)
	if err != nil {
		t.Fatal("Failed to query row data:", err)
	}

	t.Log("Captured data:", data)

	// check values

	if data["name"] != "test1" {
		t.Errorf("Expected name 'test1' and age 25, got %s and %d", data["name"], data["age"])
	}

	pp.Println(data)

	if float64(data["age"].(int64)) != float64(25) {
		t.Errorf("Expected age 25, got %d", data["age"])
	}

}
