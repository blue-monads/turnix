package main

import (
	"github.com/bornjre/turnix/backend/distro"
	"github.com/k0kubun/pp"

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

	defer app.Stop()

	mig, err := app.NeedsMigrate()
	handle(err)

	if mig {
		err = app.RunTestSeed()
		handle(err)
	}

	go runTest(app)

	err = app.Start()
	handle(err)

}

func handle(err error) {
	if err != nil {
		pp.Println(err.Error())
		panic(err)
	}
}
