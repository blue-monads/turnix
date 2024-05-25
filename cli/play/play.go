package main

import (
	"time"

	"github.com/bornjre/trunis/backend/app"
	"github.com/bornjre/trunis/backend/registry"
	"github.com/bornjre/trunis/backend/services/database"
	"github.com/bornjre/trunis/backend/services/signer"
	"github.com/k0kubun/pp"

	_ "github.com/mattn/go-sqlite3"

	// modules

	_ "github.com/bornjre/trunis/backend/modules/books"
	_ "github.com/bornjre/trunis/backend/modules/tracker"
	_ "github.com/bornjre/trunis/backend/modules/unloop"
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

	go play(as)

	return as.Start()

}

func play(app *app.App) {
	time.Sleep(time.Second * 10)
	pp.Println("@play_start")

	// session := app.GetDatabase().GetSession()

}
