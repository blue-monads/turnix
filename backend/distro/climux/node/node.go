package node

import (
	"github.com/bornjre/turnix/backend/distro/climux"

	_ "github.com/mattn/go-sqlite3"

	// modules

	"github.com/bornjre/turnix/backend/distro"

	_ "github.com/bornjre/turnix/backend/modules/books"

	// _ "github.com/bornjre/turnix/backend/modules/tracker"
	_ "github.com/bornjre/turnix/backend/modules/unloop"
)

func init() {

	climux.Register(&climux.Action{
		Name: "node",
		Help: "start node",
		Func: RunNode,
	})

}

func RunNode(cctx climux.Context) error {

	app, err := distro.NewApp()
	if err != nil {
		return err
	}

	defer app.Stop()

	mig, err := app.NeedsMigrate()
	if err != nil {
		return err
	}

	if mig {
		err = app.RunNormalSeed()
		if err != nil {
			return err
		}
	}

	err = app.Start()
	if err != nil {
		return err
	}

	return nil
}
