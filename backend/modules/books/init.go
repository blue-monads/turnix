package books

import (
	"github.com/blue-monads/turnix/backend/modules/books/dbops"
	"github.com/blue-monads/turnix/backend/registry"
	"github.com/blue-monads/turnix/backend/services/database"
	"github.com/blue-monads/turnix/backend/xtypes/xproject"
	"github.com/gin-gonic/gin"
)

var perminssions = []string{"read", "write", "read/write"}

func init() {
	registry.Register("books", New)
}

func New(opt xproject.BuilderOption) (*xproject.Defination, error) {

	db := opt.App.GetDatabase()

	d := db.(*database.DB)

	mod := &BookModule{
		sess:   d.GetSession(),
		db:     d,
		app:    opt.App,
		dbOpts: dbops.New(d),
	}

	def := &xproject.Defination{
		Name:                "Books",
		Slug:                "books",
		Info:                "Accounting made simpler",
		Icon:                "book-open",
		NewFormSchemaFields: []xproject.PTypeField{},
		Perminssions:        perminssions,
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
