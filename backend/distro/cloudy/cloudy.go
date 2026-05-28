package cloudy

import (
	"log/slog"
	"os"
	"path/filepath"

	"github.com/blue-monads/potatoverse/backend/app"
	"github.com/blue-monads/potatoverse/backend/engine/hubs/repohub"
	"github.com/blue-monads/potatoverse/backend/services/buddyhub"
	"github.com/blue-monads/potatoverse/backend/services/datahub/database"
	"github.com/blue-monads/potatoverse/backend/services/mailer/stdio"
	"github.com/blue-monads/potatoverse/backend/services/signer"
	"github.com/blue-monads/potatoverse/backend/xtypes"
)

const (
	DefaultWorkingDir   = "./cloudy_data"
	DefaultMasterSecret = "supersecretkey"
	DefaultName         = "CloudyVerse"
)

func RunApp() {

	logger := slog.Default()

	maindbDir := filepath.Join(DefaultWorkingDir, "datadb")
	dbFile := filepath.Join(maindbDir, "main.sqlite")

	os.MkdirAll(maindbDir, 0755)

	appOpts := &xtypes.AppOptions{
		Port:         8080,
		WorkingDir:   DefaultWorkingDir,
		MasterSecret: DefaultMasterSecret,
		Name:         DefaultName,
		Repos:        repohub.Default,
	}

	bhub := buddyhub.NewBuddyHub(appOpts, logger)

	db, err := database.NewDB(dbFile, logger)
	if err != nil {
		logger.Error("Failed to initialize database", "err", err)
		return
	}

	if err := db.Init(bhub); err != nil {
		logger.Error("Failed to initialize database", "err", err)
		return
	}

	m := stdio.NewMailer(logger.With("module", "mailer"))

	var happ *app.App

	happ = app.New(app.Option{
		Database:          db,
		Logger:            logger,
		Signer:            signer.New([]byte(DefaultMasterSecret)),
		AppOpts:           appOpts,
		Mailer:            m,
		WorkingFolderBase: appOpts.WorkingDir,
		BuddyHub:          bhub,
		OnStart: func() {
			onStart(happ)
		},
	})

	err = happ.Start()
	if err != nil {
		logger.Error("Failed to start app", "err", err)
		return
	}

}

func onStart(app *app.App) {

}
