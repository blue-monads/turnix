package main

import (
	"github.com/bornjre/turnix/backend/app"
	"github.com/bornjre/turnix/backend/registry"
	"github.com/bornjre/turnix/backend/services/database"
	"github.com/bornjre/turnix/backend/services/signer"

	_ "github.com/mattn/go-sqlite3"

	// modules

	_ "github.com/bornjre/turnix/backend/modules/books"
	_ "github.com/bornjre/turnix/backend/modules/tracker"
	_ "github.com/bornjre/turnix/backend/modules/unloop"
)

func Run() error {

	db, err := database.NewDB()
	if err != nil {
		return err
	}

	defer db.Close()

	signer := signer.New([]byte("A_long_HARD_Token"))

	as := app.New(app.Options{
		DB:           db,
		Signer:       signer,
		ProjectTypes: registry.GetAll(),
	})

	return as.Start()

}
