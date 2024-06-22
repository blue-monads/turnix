package distro

import (
	"github.com/bornjre/turnix/backend/app"
	"github.com/bornjre/turnix/backend/registry"
	"github.com/bornjre/turnix/backend/services/database"
	"github.com/bornjre/turnix/backend/services/signer"
	"github.com/bornjre/turnix/backend/xtypes"
)

type DistroApp struct {
	App xtypes.App
}

func NewApp() (*DistroApp, error) {

	db, err := database.NewDB()
	if err != nil {
		return nil, err
	}

	defer db.Close()

	signer := signer.New([]byte("A_long_HARD_Token"))

	as := app.New(app.Options{
		DB:           db,
		Signer:       signer,
		ProjectTypes: registry.GetAll(),
	})

	return &DistroApp{
		App: as,
	}, nil

}

func (d *DistroApp) Run() error {
	return d.App.Start()
}

func (d *DistroApp) NeedsMigrate() (bool, error) {

	db := d.App.GetDatabase()

	usrs, err := db.ListUser()
	if err != nil {
		return false, err
	}

	return len(usrs) == 0, nil

}
