package books

import (
	"github.com/bornjre/turnix/backend/registry"
	"github.com/bornjre/turnix/backend/services/database"
	"github.com/bornjre/turnix/backend/xtypes/xproject"
	"github.com/gin-gonic/gin"
)

var eventTypes = []string{
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
}

var perminssions = []string{"read", "write", "read/write"}

func init() {
	registry.Register("books", New)
}

func New(opt xproject.BuilderOption) (*xproject.Defination, error) {

	db := opt.App.GetDatabase()

	d := db.(*database.DB)

	mod := &BookModule{
		sess: d.GetSession(),
		db:   d,
		app:  opt.App,
	}

	def := &xproject.Defination{
		Name:                "Books",
		Slug:                "books",
		Info:                "Accounting made simpler",
		Icon:                "book-open",
		NewFormSchemaFields: []xproject.PTypeField{},
		Perminssions:        perminssions,
		EventTypes:          eventTypes,

		OnDeInit: func(pid int64) error {
			return mod.DeInit(pid)
		},

		IsInitilized: func(pid int64) (bool, error) {
			return true, nil
		},

		OnInit: func(pid int64) error {
			return mod.Init(pid)
		},

		OnAPIMount: func(rg *gin.RouterGroup) {
			mod.register(rg)
		},
	}

	return def, nil
}
