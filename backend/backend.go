package backend

import (
	"github.com/bornjre/trunis/backend/app"
	"github.com/bornjre/trunis/backend/registry"
	"github.com/bornjre/trunis/backend/services/database"
	"github.com/bornjre/trunis/backend/services/signer"
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
		DB:           db,
		Signer:       signer,
		ProjectTypes: registry.GetAll(),
	})

	return as.Start()

}
