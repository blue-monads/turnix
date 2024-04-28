package books

import (
	_ "embed"

	"github.com/bornjre/trunis/backend/xtypes"
	"github.com/bornjre/trunis/backend/xtypes/services/xdatabase"
	"github.com/bornjre/trunis/backend/xtypes/services/xsockd"
	"github.com/bornjre/trunis/backend/xtypes/xproject"
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

func (b *BookModule) IsInitilized(pid int64) (bool, error) { return false, nil }

func (b *BookModule) DeInit(pid int64) error { return nil }

func (b *BookModule) OnSockdMessage(msg *xsockd.Message) error { return nil }

func (b *BookModule) OnSockdConn(opts *xsockd.ConnectOptions) error { return nil }

func (b *BookModule) OnFileEvent(event *xproject.FileEvent) error { return nil }

func (b *BookModule) OnUserEvent(event *xproject.UserEvent) error { return nil }
