package main

import (
	"github.com/bornjre/turnix/backend/distro"

	_ "github.com/mattn/go-sqlite3"

	// modules

	_ "github.com/bornjre/turnix/backend/modules/books"
	_ "github.com/bornjre/turnix/backend/modules/tracker"
	_ "github.com/bornjre/turnix/backend/modules/unloop"
)

func main() {

	app, err := distro.NewAppWithOptions(distro.Options{
		MasterSecret: "ftdfwyguhytcyuagiuhs",
		HttpPort:     ":8777",
	})
	handle(err)

	mig, err := app.NeedsMigrate()
	handle(err)

	if mig {
		err = app.RunTestSeed()
		handle(err)
	}

	go runTest(app)

	err = app.Run()
	handle(err)

}

func handle(err error) {
	if err != nil {
		panic(err)
	}
}
