package app

import (
	"net/http"
	"time"

	"github.com/bornjre/trunis/backend/database"
	"github.com/bornjre/trunis/backend/token"
	"github.com/bwmarrin/snowflake"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"
)

type App struct {
	db        *database.DB
	signer    *token.TokenService
	flakeNode *snowflake.Node
}

func New(db *database.DB, signer *token.TokenService) *App {

	node, err := snowflake.NewNode(0)
	if err != nil {
		panic(err)
	}

	return &App{
		db:        db,
		signer:    signer,
		flakeNode: node,
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
		AllowOriginFunc: func(origin string) bool {
			return origin == "https://github.com"
		},
		MaxAge: 12 * time.Hour,
	}))

	a.bindRoutes(r)

	return r.Run()
}
