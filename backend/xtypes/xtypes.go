package xtypes

import (
	"github.com/bornjre/trunis/backend/database"
	"github.com/bornjre/trunis/backend/services/signer"
	"github.com/gin-gonic/gin"
)

type ApiHandler func(claim *signer.AccessClaim, ctx *gin.Context) (any, error)

type App interface {
	Start() error
	Stop() error

	AuthMiddleware(fn ApiHandler) gin.HandlerFunc
	GetDatabase() *database.DB
}
