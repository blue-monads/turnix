package app

import (
	"os"

	"github.com/bornjre/turnix/backend/services/database"
	"github.com/bornjre/turnix/backend/services/signer"
	"github.com/bornjre/turnix/backend/xtypes/services/xdatabase"
	"github.com/bornjre/turnix/backend/xtypes/services/xsockd"
	"github.com/bornjre/turnix/backend/xtypes/xbus"
	"github.com/bornjre/turnix/backend/xtypes/xproject"
	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type App struct {
	db         *database.DB
	signer     *signer.Signer
	flakeNode  *snowflake.Node
	globalJS   []byte
	ptypeDefs  []*xproject.Defination
	projects   map[string]ProjectInstance
	rootLogger zerolog.Logger
}

type ProjectInstance struct {
	Project xproject.ProjectType
	Def     *xproject.Defination
}

type Options struct {
	DB     *database.DB
	Signer *signer.Signer

	ProjectTypes []*xproject.Defination
}

func New(opts Options) *App {

	node, err := snowflake.NewNode(0)
	if err != nil {
		panic(err)
	}

	out, err := buildGlobalJS(opts.ProjectTypes)
	if err != nil {
		panic(err)
	}

	rootLogger := zerolog.New(os.Stdout)

	return &App{
		db:         opts.DB,
		signer:     opts.Signer,
		flakeNode:  node,
		globalJS:   out,
		ptypeDefs:  opts.ProjectTypes,
		rootLogger: rootLogger,
		// 		hookEngine: hookengine.New(opts.DB, opts.Signer, rootLogger.With().Str("service", "engine").Logger()),
	}
}

func (a *App) Start(port string) error {

	r := gin.Default()

	a.bindRoutes(r)

	// err := a.hookEngine.Init()
	// if err != nil {
	// 	return err
	// }

	return r.Run(port)
}

func (a *App) Stop() error {
	return a.db.Close()
}

func (a *App) GetDatabase() xdatabase.Database {
	return a.db
}

func (a *App) GetEventBus() xbus.EventBus {
	return nil
}

func (a *App) GetSockd() xsockd.Sockd {
	return nil
}

func (a *App) NewId() int64 {
	return a.flakeNode.Generate().Int64()
}
