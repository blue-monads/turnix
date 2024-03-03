package database

import (
	"database/sql"
	_ "embed"
	"log"

	"github.com/upper/db/v4"
	"github.com/upper/db/v4/adapter/sqlite"
)

//go:embed schema.sql
var schema string

type DB struct {
	sess db.Session
}

func NewDB() (*DB, error) {

	var settings = sqlite.ConnectionURL{
		Database: `foo.db`, // Path to database file
	}

	sess, err := sqlite.Open(settings)
	if err != nil {
		log.Fatalf("db.Open(): %q\n", err)
	}

	driver := sess.Driver().(*sql.DB)

	_, err = driver.Exec(schema)
	if err != nil {
		sess.Close()
		return nil, err
	}

	return &DB{
		sess: sess,
	}, nil
}

func (db *DB) Close() error {
	return db.sess.Close()
}

func (db *DB) table(name string) db.Collection {
	return db.sess.Collection(name)
}
