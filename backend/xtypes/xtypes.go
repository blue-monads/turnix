package xtypes

import (
	"github.com/bornjre/trunis/backend/xtypes/services/xdatabase"
	"github.com/bornjre/trunis/backend/xtypes/services/xsockd"

	"github.com/gin-gonic/gin"
)

type ApiHandler func(ctx ContextPlus) (any, error)

type App interface {
	Start() error
	Stop() error

	AuthMiddleware(fn ApiHandler) gin.HandlerFunc
	GetDatabase() xdatabase.Database
	GetSockd() xsockd.Sockd

	NewId() int64
}
