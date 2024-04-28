package books

import (
	"github.com/bornjre/trunis/backend/registry"
	"github.com/bornjre/trunis/backend/xtypes/xproject"
)

var defs = &xproject.Defination{
	Name:                "Books",
	Slug:                "books",
	Info:                "Accounting made simpler",
	Icon:                "book-open",
	NewFormSchemaFields: []xproject.PTypeField{},
	Builder:             New,
}

func init() {
	registry.Register(defs)
}

func New(opt xproject.BuilderOption) (xproject.ProjectType, error) {

	db := opt.App.GetDatabase()

	mod := &BookModule{
		sess: db.GetSession(),
		db:   db,
		app:  opt.App,
	}

	err := mod.register(opt.RouterGroup)
	if err != nil {
		return nil, err
	}

	return xproject.NewAdapter(mod), nil
}
