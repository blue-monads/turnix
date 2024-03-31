package app

import (
	"net/http"
	"time"

	"github.com/bornjre/trunis/backend/database"
	"github.com/bornjre/trunis/backend/services/signer"
	"github.com/bornjre/trunis/backend/xtypes/xproject"
	"github.com/bwmarrin/snowflake"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type App struct {
	db        *database.DB
	signer    *signer.Signer
	flakeNode *snowflake.Node
	globalJS  []byte
}

type Options struct {
	DB     *database.DB
	Signer *signer.Signer

	ProjectTypes []*xproject.TypeDefination
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
	}
}

func (a *App) Run() error {

	r := gin.Default()

	r.Use(cors.New(cors.Config{
		AllowOrigins:     []string{"*"},
		AllowMethods:     []string{http.MethodGet, http.MethodPost, http.MethodPut, http.MethodPatch, http.MethodDelete},
		AllowHeaders:     []string{"Origin", "Authorization"},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		MaxAge:           12 * time.Hour,
		AllowOriginFunc: func(origin string) bool {
			return true
		},
	}))

	a.bindRoutes(r)

	return r.Run(":7777")
}
