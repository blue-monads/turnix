package binds

import (
	"database/sql"
	"log"
	"os"
	"testing"

	"github.com/k0kubun/pp"
	"github.com/upper/db/v4/adapter/sqlite"
	lua "github.com/yuin/gopher-lua"
)

func TestMain(m *testing.M) {
	// remmove test.db file if it exists
	if _, err := os.Stat("test.db"); err == nil {
		os.Remove("test.db")
	}

	code := m.Run()

	os.Exit(code)

	if _, err := os.Stat("test.db"); err == nil {
		os.Remove("test.db")
	}

	os.Exit(code)
}

func TestBinds(t *testing.T) {
	pp.Println("test")

	var settings = sqlite.ConnectionURL{
		Database: "test.db",
	}

	sess, err := sqlite.Open(settings)
	if err != nil {
		log.Fatalf("db.Open(): %q\n", err)
	}

	db := sess.Driver().(*sql.DB)

	bindDb := BluaDb{
		db: db,
	}

	_, err = db.Exec(`

	CREATE TABLE IF NOT EXISTS users (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		name TEXT NOT NULL
	);

	INSERT INTO users (name) VALUES ('John Doe');

	`)

	if err != nil {
		t.Fatalf("db.Exec(): %q\n", err)
	}

	L := lua.NewState()

	// Preload the db module first
	L.PreloadModule("db", func(L *lua.LState) int {
		return bindDb.Bind(L)
	})

	// Now load and execute the Lua code
	err = L.DoString(`
	
	function test()
		local db = require("db")

		print(db.query)

		if not db then
			error("Failed to load db module")
		end
		if not db.query then
			error("db.query is not a function")
		end
		local result, err = db.query("SELECT * FROM users")
		print(result, err)
		return result
	end

	`)
	if err != nil {
		t.Fatalf("L.DoString(): %q\n", err)
	}

	tfunc := L.GetGlobal("test")
	if tfunc.Type() != lua.LTFunction {
		t.Fatalf("Expected function, got %s", tfunc.Type())
	}

	L.Push(tfunc)
	err = L.PCall(0, 1, nil)
	if err != nil {
		t.Fatalf("Error calling test(): %v", err)
	}

	pp.Println(L.Get(-1))

}
