package node

import (
	"github.com/bornjre/turnix/backend/distro"
	"github.com/bornjre/turnix/backend/distro/climux"
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
