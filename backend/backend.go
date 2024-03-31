package backend

import (
	"github.com/bornjre/trunis/backend/app"
	"github.com/bornjre/trunis/backend/database"
	"github.com/bornjre/trunis/backend/services/signer"
	"github.com/bornjre/trunis/backend/xtypes/xproject"
)

func Run() error {

	db, err := database.NewDB()
	if err != nil {
		return err
	}

	defer db.Close()

	err = db.Init()
	if err != nil {
		return err
	}

	signer := signer.New([]byte("A_long_HARD_Token"))

	as := app.New(app.Options{
		DB:     db,
		Signer: signer,
		ProjectTypes: []*xproject.TypeDefination{
			{
				Name:               "On Loop",
				Slug:               "onloop",
				Info:               "Perform action with you on loop",
				Icon:               "arrow-path",
				InterceptFileEvent: false,
			},
		},
	})

	return as.Run()

}
