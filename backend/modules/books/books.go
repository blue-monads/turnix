package books

import (
	_ "embed"
	"fmt"
	"strings"

	"github.com/bornjre/turnix/backend/xtypes"
	"github.com/bornjre/turnix/backend/xtypes/services/xdatabase"
	"github.com/k0kubun/pp"
	"github.com/upper/db/v4"
)

//go:embed books.sql
var Schema string

type BookModule struct {
	app  xtypes.App
	sess db.Session
	db   xdatabase.Database
}

func (b *BookModule) Init(pid int64) error {

	parameterizedSchema := strings.ReplaceAll(Schema, "__project__", fmt.Sprintf("_%d_", pid))

	pp.Println("@parameterizedSchema")
	fmt.Println(parameterizedSchema)

	err := b.db.RunDDL(xdatabase.DDLContext{
		DDL: parameterizedSchema,
	})
	if err != nil {
		pp.Println("@err", err.Error())
	}

	return err

}

func (b *BookModule) DeInit(pid int64) error { return nil }
