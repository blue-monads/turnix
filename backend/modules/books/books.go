package books

import (
	_ "embed"
	"fmt"
	"strings"

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

	parameterizedSchema := strings.ReplaceAll(Schema, "__project__", fmt.Sprintf("z_%d_", pid))

	pp.Println("@parameterizedSchema")
	fmt.Println(parameterizedSchema)

	err := b.db.RunDDL(parameterizedSchema)
	if err != nil {
		pp.Println("@err", err.Error())
	}

	return err

}

func (b *BookModule) DeInit(pid int64) error { return nil }
