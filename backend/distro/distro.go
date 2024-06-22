package distro

import (
	"github.com/bornjre/turnix/backend/app"
	"github.com/bornjre/turnix/backend/registry"
	"github.com/bornjre/turnix/backend/services/database"
	"github.com/bornjre/turnix/backend/services/signer"
	"github.com/bornjre/turnix/backend/xtypes"
)

type Options struct {
	MasterSecret string
	HttpPort     string
}

type DistroApp struct {
	Options Options
	App     xtypes.App
}

var defaultOption = Options{
	MasterSecret: "A_long_HARD_Token",
	HttpPort:     ":7777",
}

func NewApp() (*DistroApp, error) {
	return NewAppWithOptions(defaultOption)
}

func NewAppWithOptions(opts Options) (*DistroApp, error) {

	db, err := database.NewDB()
	if err != nil {
		return nil, err
	}

	defer db.Close()

	signer := signer.New([]byte(opts.MasterSecret))

	as := app.New(app.Options{
		DB:           db,
		Signer:       signer,
		ProjectTypes: registry.GetAll(),
	})

	return &DistroApp{
		App:     as,
		Options: opts,
	}, nil

}

func (d *DistroApp) Run() error {
	return d.App.Start(d.Options.HttpPort)
}

func (d *DistroApp) NeedsMigrate() (bool, error) {

	db := d.App.GetDatabase()

	usrs, err := db.ListUser()
	if err != nil {
		return false, err
	}

	return len(usrs) == 0, nil

}
