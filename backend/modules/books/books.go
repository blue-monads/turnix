package books

import (
	_ "embed"
	"fmt"

	"github.com/blue-monads/turnix/backend/modules"
	"github.com/blue-monads/turnix/backend/modules/books/dbops"
	"github.com/blue-monads/turnix/backend/services/database"
	"github.com/blue-monads/turnix/backend/xtypes"
	"github.com/k0kubun/pp"
	"github.com/upper/db/v4"
)

//go:embed books.sql
var Schema string

type BookModule struct {
	app    xtypes.App
	sess   db.Session
	db     *database.DB
	dbOpts dbops.DbOps
}

func (b *BookModule) Init(pid int64) error {

	schema := modules.ParamaterizedSchema(Schema, pid)

	pp.Println("@parameterizedSchema")
	fmt.Println(schema)

	err := b.db.RunDDL(schema)
	if err != nil {
		pp.Println("@err", err.Error())
	}

	return err

}

func (b *BookModule) DeInit(pid int64) error { return nil }
