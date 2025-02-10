package xtypes

import (
	"github.com/blue-monads/turnix/backend/xtypes/services/xsockd"

	"github.com/blue-monads/turnix/backend/xtypes/xengine"

	"github.com/gin-gonic/gin"
)

type ApiHandler func(ctx ContextPlus) (any, error)

type App interface {
	Start(string) error
	Stop() error

	AuthMiddleware(fn ApiHandler) gin.HandlerFunc
	AsApiAction(name string, fn ApiHandler) gin.HandlerFunc

	GetDatabase() any
	GetSockd() xsockd.Sockd

	GetController() any
	GetServer() any
	GetSigner() any

	GetEngine() xengine.XEngine
}
