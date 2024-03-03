package app

import (
	"github.com/bornjre/trunis/backend/database"
	"github.com/bornjre/trunis/backend/token"
	"github.com/bwmarrin/snowflake"
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

	a.bindRoutes(r)

	return r.Run()
}
