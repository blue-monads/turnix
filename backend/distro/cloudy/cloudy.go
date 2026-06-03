package cloudy

import (
	"context"
	"log"
	"log/slog"

	"github.com/blue-monads/potatoverse/backend/app"
	_ "github.com/blue-monads/potatoverse/backend/distro"
	"github.com/blue-monads/potatoverse/backend/engine/hubs/repohub"
	"github.com/blue-monads/potatoverse/backend/services/datahub/database"
	"github.com/blue-monads/potatoverse/backend/services/mailer/stdio"
	"github.com/blue-monads/potatoverse/backend/services/signer"
	"github.com/blue-monads/potatoverse/backend/xtypes"
	"github.com/k0kubun/pp"
	turso "turso.tech/database/tursogo"
)

type Config struct {
	Port           int
	WorkingDir     string
	MasterSecret   string
	TursoAuthToken string
	TursoRemoteURL string
}

type CloudyApp struct {
	rootCtx context.Context
	tursoDB *turso.TursoSyncDb
	config  *Config

	app xtypes.App

	onBuild chan struct{}
}

func New(config *Config) (*CloudyApp, error) {

	ctx := context.Background()

	bootstrap := true

	db, err := turso.NewTursoSyncDb(ctx, turso.TursoSyncDbConfig{
		Path:             "aa.db",
		RemoteUrl:        config.TursoRemoteURL,
		AuthToken:        config.TursoAuthToken,
		BootstrapIfEmpty: &bootstrap,
	})

	if err != nil {
		pp.Print("@")
		return nil, err
	}

	return &CloudyApp{
		rootCtx: ctx,
		tursoDB: db,
		config:  config,
		onBuild: make(chan struct{}),
	}, nil

}

func (a *CloudyApp) Build() error {

	db, err := a.tursoDB.Connect(a.rootCtx)
	if err != nil {
		log.Fatal(err)
	}

	logger := slog.Default()

	adb, err := database.FromSqlHandle(db, logger)
	if err != nil {
		log.Fatal(err)
	}

	appOpts := &xtypes.AppOptions{
		Port:         8080,
		WorkingDir:   a.config.WorkingDir,
		MasterSecret: a.config.MasterSecret,
		Name:         "Cloudy",
		Repos:        repohub.Default,
	}

	m := stdio.NewMailer(logger.With("module", "mailer"))

	var happ *app.App

	happ = app.New(app.Option{
		Database:          adb,
		Logger:            logger,
		Signer:            signer.New([]byte(a.config.MasterSecret)),
		AppOpts:           appOpts,
		Mailer:            m,
		WorkingFolderBase: appOpts.WorkingDir,
		BuddyHub:          nil,
		OnStart: func() {

			a.onBuild <- struct{}{}
			log.Println("Cloudy is running on port", appOpts.Port)

		},
	})

	a.app = happ

	return nil

}

func (a *CloudyApp) Run() error {
	if err := a.Build(); err != nil {
		return err
	}

	var err error

	go func() {
		err = a.app.Start()
	}()

	<-a.onBuild

	if err != nil {
		return err
	}

	return nil

}
