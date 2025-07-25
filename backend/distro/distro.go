package distro

import (
	"os"
	"path"
	"strings"

	"github.com/blue-monads/turnix/backend/app"
	"github.com/blue-monads/turnix/backend/registry"
	"github.com/blue-monads/turnix/backend/services/database"
	"github.com/blue-monads/turnix/backend/services/signer"
	"github.com/blue-monads/turnix/backend/xtypes"
)

type Options struct {
	BasePath           string
	MasterSecret       string
	HttpPort           string
	LocalSocket        string
	DatabaseFile       string
	ProjectInstallPath string
}

type DistroApp struct {
	Options Options
	App     xtypes.App
}

var DefaultOption = Options{
	MasterSecret: "A_long_HARD_Token",
	HttpPort:     ":7703",
	LocalSocket:  "/tmp/turnix.sock",
	DatabaseFile: "./data.db",
}

func NewApp() (*DistroApp, error) {
	return NewAppWithOptions(DefaultOption)
}

func NewAppWithOptions(opts Options) (*DistroApp, error) {

	db, err := database.NewDB(opts.DatabaseFile)
	if err != nil {
		return nil, err
	}

	signer := signer.New([]byte(opts.MasterSecret))

	basePath := opts.BasePath
	if basePath == "" {
		wd, err := os.Getwd()
		if err != nil {
			return nil, err
		}

		basePath = path.Join(wd, ".turnix")
	}

	pappPath := opts.ProjectInstallPath
	if pappPath == "" {
		pappPath = path.Join(basePath, "papps")
	}

	if strings.HasPrefix(pappPath, "./") {
		pappPath = path.Join(basePath, pappPath)
	}

	as := app.New(app.Options{
		DB:                 db,
		Signer:             signer,
		ProjectBuilders:    registry.GetAll(),
		LocalSocket:        opts.LocalSocket,
		DevMode:            true,
		BasePath:           basePath,
		ProjectInstallPath: pappPath,
	})

	return &DistroApp{
		App:     as,
		Options: opts,
	}, nil

}

func (d *DistroApp) Stop() error {

	return d.App.Stop()
}

func (d *DistroApp) Start() error {
	return d.App.Start(d.Options.HttpPort)
}

func (d *DistroApp) NeedsMigrate() (bool, error) {

	db := d.App.GetDatabase()

	dz := db.(*database.DB)

	usrs, err := dz.ListUser()
	if err != nil {
		return false, err
	}

	return len(usrs) == 0, nil

}
