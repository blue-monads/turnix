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
	Perminssions:        []string{"read", "write", "read/write"},
	EventTypes: []string{
		"after_txn_add",
		"after_txn_edit",
		"after_txn_remove",
		"before_txn_remove",
		"before_txn_add",
		"before_txn_edit",
		"after_account_add",
		"after_account_edit",
		"after_account_remove",
		"before_account_remove",
		"before_account_add",
		"before_account_edit",
	},
	Builder: New,
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
