package app

import (
	"github.com/bornjre/trunis/backend/services/database"
	"github.com/bornjre/trunis/backend/services/signer"
	"github.com/bornjre/trunis/backend/xtypes/xproject"
	"github.com/bwmarrin/snowflake"
	"github.com/gin-gonic/gin"
)

type App struct {
	db        *database.DB
	signer    *signer.Signer
	flakeNode *snowflake.Node
	globalJS  []byte
	ptypeDefs []*xproject.Defination
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
		db:        opts.DB,
		signer:    opts.Signer,
		flakeNode: node,
		globalJS:  out,
		ptypeDefs: opts.ProjectTypes,
	}
}

func (a *App) Run() error {

	r := gin.Default()

	a.bindRoutes(r)

	return r.Run(":7777")
}
