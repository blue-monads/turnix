package app

import (
	"os"

	hookengine "github.com/bornjre/trunis/backend/hookEngine"
	"github.com/bornjre/trunis/backend/services/database"
	"github.com/bornjre/trunis/backend/services/signer"
	"github.com/bornjre/trunis/backend/xtypes"
	"github.com/bornjre/trunis/backend/xtypes/services/xdatabase"
	"github.com/bornjre/trunis/backend/xtypes/services/xsockd"
	"github.com/bornjre/trunis/backend/xtypes/xproject"
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
	projects   map[string]xproject.ProjectType
	rootLogger zerolog.Logger
	hookEngine *hookengine.HookEngine
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

	return &App{
		db:         opts.DB,
		signer:     opts.Signer,
		flakeNode:  node,
		globalJS:   out,
		ptypeDefs:  opts.ProjectTypes,
		projects:   make(map[string]xproject.ProjectType),
		rootLogger: zerolog.New(os.Stdout),
		hookEngine: hookengine.New(opts.DB),
	}
}

func (a *App) Start() error {

	r := gin.Default()

	a.bindRoutes(r)

	err := a.hookEngine.Init()
	if err != nil {
		return err
	}

	return r.Run(":7777")
}

func (a *App) Stop() error {
	return nil
}

func (a *App) GetDatabase() xdatabase.Database {
	return a.db
}

func (a *App) GetHookEngine() xtypes.HookEngine {
	return a.hookEngine
}

func (a *App) GetSockd() xsockd.Sockd {
	return nil
}

func (a *App) NewId() int64 {
	return a.flakeNode.Generate().Int64()
}
