package simplerat

import (
	_ "embed"

	"github.com/blue-monads/turnix/backend/modules"
	"github.com/blue-monads/turnix/backend/registry"
	"github.com/blue-monads/turnix/backend/services/database"
	"github.com/blue-monads/turnix/backend/services/signer"
	"github.com/blue-monads/turnix/backend/xtypes/xproject"
	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"
)

var perminssions = []string{"read", "write", "read/write"}

func init() {
	registry.Register("simplerat", New)
}

//go:embed simplerat.sql
var Schema string

func New(opt xproject.BuilderOption) (*xproject.Defination, error) {

	db := opt.App.GetDatabase().(*database.DB)

	mod := &ECPServer{
		db:     db,
		app:    opt.App,
		signer: opt.App.GetSigner().(*signer.Signer),
	}

	def := &xproject.Defination{
		Name:                "SimpleRAT",
		Slug:                "simplerat",
		Info:                "Remote Access Tool",
		Icon:                "book-open",
		NewFormSchemaFields: []xproject.PTypeField{},
		Perminssions:        perminssions,
		OnAPIMount: func(rg *gin.RouterGroup) {
			mod.register(rg)
		},
		OnInit: func(pid int64) error {

			schema := modules.ParamaterizedSchema(Schema, pid)

			err := db.RunDDL(schema)
			if err != nil {
				pp.Println("@err", err.Error())
			}

			return err

		},
	}

	return def, nil
}
