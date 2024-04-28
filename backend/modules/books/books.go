package books

import (
	_ "embed"

	"github.com/bornjre/trunis/backend/xtypes"
	"github.com/bornjre/trunis/backend/xtypes/services/xdatabase"
	"github.com/upper/db/v4"
)

//go:embed books.sql
var Schema string

type BookModule struct {
	app  xtypes.App
	sess db.Session
	db   xdatabase.Database
}

func (b *BookModule) Init(pid int64) error { return nil }

func (b *BookModule) DeInit(pid int64) error { return nil }
