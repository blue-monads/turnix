package database

import (
	"database/sql"
	_ "embed"
	"log"
	"strings"

	"github.com/bornjre/turnix/backend/modules/books"
	"github.com/bornjre/turnix/backend/xtypes/services/xdatabase"
	"github.com/gobuffalo/fizz"
	"github.com/gobuffalo/fizz/translators"
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
		Database: `data.db`,
	}

	sess, err := sqlite.Open(settings)
	if err != nil {
		log.Fatalf("db.Open(): %q\n", err)
	}

	exists, _ := sess.Collection("Users").Exists()

	if !exists {
		driver := sess.Driver().(*sql.DB)

		_, err = driver.Exec(schema)
		if err != nil {
			sess.Close()
			return nil, err
		}

		_, err = driver.Exec(books.Schema)
		if err != nil {
			sess.Close()
			return nil, err
		}

	}

	return &DB{
		sess: sess,
	}, nil
}

func (db *DB) Init() error {

	return nil
}

func (db *DB) Close() error {
	return db.sess.Close()
}

func (db *DB) Vender() string {
	return "sqlite"
}

func (db *DB) RunDDL(ctx xdatabase.DDLContext) error {
	var buf strings.Builder

	t := translators.NewSQLite("")

	if ctx.DDL != "" {
		str, err := fizz.AString(ctx.DDL, t)
		if err != nil {
			return err
		}

		buf.WriteString(str)

	} else {
		for _, tbl := range ctx.Tables {
			str, err := fizz.AString(tbl.String(), t)
			if err != nil {
				return err
			}
			buf.WriteString(str)
		}
	}

	driver := db.sess.Driver().(*sql.DB)

	_, err := driver.Exec(buf.String())
	if err != nil {
		return err
	}

	return nil
}

func (db *DB) GetSession() db.Session {
	return db.sess
}

func (db *DB) Table(name string) db.Collection {
	return db.sess.Collection(name)
}
