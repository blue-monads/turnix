package app

import (
	"log"
	"os"

	"github.com/blue-monads/turnix/backend/app/server"
	"github.com/blue-monads/turnix/backend/controller"
	"github.com/blue-monads/turnix/backend/services/database"
	"github.com/blue-monads/turnix/backend/services/signer"
	"github.com/blue-monads/turnix/backend/xtypes"
	"github.com/blue-monads/turnix/backend/xtypes/services/xsockd"
	"github.com/blue-monads/turnix/backend/xtypes/xengine"
	"github.com/blue-monads/turnix/backend/xtypes/xproject"
	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"
	"github.com/rs/zerolog"
)

type App struct {
	db         *database.DB
	signer     *signer.Signer
	projects   map[string]*xproject.Defination
	rootLogger zerolog.Logger
	controller *controller.RootController
	server     *server.Server
}

type Options struct {
	DB              *database.DB
	Signer          *signer.Signer
	ProjectBuilders map[string]xproject.Builder
	LocalSocket     string
}

func New(opts Options) *App {

	rootLogger := zerolog.New(os.Stdout)

	app := &App{
		db:         opts.DB,
		signer:     opts.Signer,
		rootLogger: rootLogger,
		projects:   make(map[string]*xproject.Defination),
	}

	for pbName, pBuilder := range opts.ProjectBuilders {
		pp.Println("@buinding_ptype", pbName)

		subLogger := rootLogger.With().Str("ptype", pbName).Logger()

		proj, err := pBuilder(xproject.BuilderOption{
			App:    app,
			Logger: subLogger,
		})
		if err != nil {
			log.Fatal(err)
		}

		app.projects[pbName] = proj

	}

	app.controller = controller.New(app.db, app.projects)

	app.server = server.New(server.Options{
		DB:              app.db,
		Signer:          app.signer,
		ProjectBuilders: app.projects,
		Controller:      app.controller,
		LocalSocket:     opts.LocalSocket,
		DevMode:         os.Getenv("TURNIX_DEV_MODE") == "true",
	})

	return app
}

func (a *App) Start(port string) error {
	return a.server.Start(port)
}

func (a *App) Stop() error {
	return a.db.Close()
}

func (a *App) GetDatabase() any {
	return a.db
}

func (a *App) GetEngine() xengine.XEngine {
	return nil
}

func (a *App) GetSockd() xsockd.Sockd {
	return nil
}

func (a *App) GetController() any {
	return a.controller
}

func (a *App) GetServer() any {
	return a.server
}

func (a *App) AuthMiddleware(fn xtypes.ApiHandler) gin.HandlerFunc {
	return a.server.AuthMiddleware(fn)
}

func (a *App) AsApiAction(name string, fn xtypes.ApiHandler) gin.HandlerFunc {
	return a.server.AsApiAction(name, fn)
}
