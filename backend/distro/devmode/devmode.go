package devmode

import (
	_ "github.com/mattn/go-sqlite3"

	// modules

	"github.com/blue-monads/turnix/backend/distro"
)

func Run() {

	app, err := distro.NewApp()
	if err != nil {
		panic(err)
	}

	defer app.Stop()

	mig, err := app.NeedsMigrate()
	if err != nil {
		panic(err)
	}

	if mig {
		err = app.RunNormalSeed()
		if err != nil {
			panic(err)
		}
	}

	err = app.Start()
	if err != nil {
		panic(err)
	}

}
