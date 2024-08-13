package main

import (
	"os"
	"sync"
	"time"

	"github.com/bornjre/turnix/backend/distro"
	"github.com/bornjre/turnix/tests/cases/assetserve"
	hooktest "github.com/bornjre/turnix/tests/cases/hook"
	"github.com/bornjre/turnix/tests/must"
	"github.com/k0kubun/pp"

	_ "github.com/mattn/go-sqlite3"

	// modules

	_ "github.com/bornjre/turnix/backend/modules/books"
	//_ "github.com/bornjre/turnix/backend/modules/tracker"
	_ "github.com/bornjre/turnix/backend/modules/unloop"
)

func main() {

	app, err := distro.NewAppWithOptions(distro.Options{
		MasterSecret: "ftdfwyguhytcyuagiuhs",
		HttpPort:     ":8777",
	})
	must.Handle(err)

	defer app.Stop()

	mig, err := app.NeedsMigrate()
	must.Handle(err)

	if mig {
		err = app.RunTestSeed()
		must.Handle(err)
	}

	go runTest(app)

	err = app.Start()
	must.Handle(err)

}

func runTest(app *distro.DistroApp) {

	pp.Println("@runTest/sleeping/before")

	time.Sleep(time.Second * 3)

	pp.Println("@runTest/sleeping/end")

	wg := &sync.WaitGroup{}

	wg.Add(1)
	go hooktest.Run(app, wg)

	wg.Add(1)
	go assetserve.Run(app, wg)

	wg.Wait()

	os.Exit(0)

}
