package xtypes

import (
	"github.com/bornjre/turnix/backend/xtypes/services/xdatabase"
	"github.com/bornjre/turnix/backend/xtypes/services/xsockd"
	"github.com/bornjre/turnix/backend/xtypes/xbus"

	"github.com/gin-gonic/gin"
)

type ApiHandler func(ctx ContextPlus) (any, error)

type App interface {
	Start(string) error
	Stop() error

	AuthMiddleware(fn ApiHandler) gin.HandlerFunc
	GetDatabase() xdatabase.Database
	GetSockd() xsockd.Sockd

	GetEventBus() xbus.EventBus

	NewId() int64
}
