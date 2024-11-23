package node

import (
	"errors"
	"os"

	"github.com/blue-monads/turnix/backend/distro"
	"github.com/blue-monads/turnix/backend/distro/climux"
)

func init() {

	climux.Register(&climux.Action{
		Name: "node",
		Help: "start node",
		Func: RunNode,
	})

}

func RunNode(cctx climux.Context) error {

	masterSecret := os.Getenv("TURNIX_MASTER_SECRET")
	if masterSecret == "" {
		return errors.New("TURNIX_MASTER_SECRET not set")
	}

	httpPort := os.Getenv("TURNIX_HTTP_PORT")
	if httpPort == "" {
		return errors.New("TURNIX_HTTP_PORT not set")
	}

	localSocket := os.Getenv("TURNIX_LOCAL_SOCKET")
	if localSocket == "" {
		return errors.New("TURNIX_LOCAL_SOCKET not set")
	}

	databaseFile := os.Getenv("TURNIX_DATABASE_FILE")
	if databaseFile == "" {
		return errors.New("TURNIX_DATABASE_FILE not set")
	}

	app, err := distro.NewAppWithOptions(distro.Options{
		MasterSecret: masterSecret,
		HttpPort:     httpPort,
		LocalSocket:  localSocket,
		DatabaseFile: databaseFile,
	})

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
