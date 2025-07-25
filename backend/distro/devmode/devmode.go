package devmode

import (
	"os"
	"path"

	_ "github.com/mattn/go-sqlite3"

	// modules

	"github.com/blue-monads/turnix/backend/distro"
)

func Run() {

	opts := distro.Options{
		MasterSecret:       "A_long_HARD_Token",
		HttpPort:           ":7703",
		LocalSocket:        "/tmp/turnix.sock",
		DatabaseFile:       "./data.db",
		BasePath:           "./tmp",
		ProjectInstallPath: "./papps",
	}

	wd, err := os.Getwd()
	if err != nil {
		panic(err)
	}

	appsPath := path.Join(wd, "backend/engine/papps")
	opts.ProjectInstallPath = appsPath

	app, err := distro.NewAppWithOptions(opts)
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
