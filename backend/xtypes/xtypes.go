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

	GetHookEngine() HookEngine

	NewId() int64
}

type HookEvent struct {
	Name      string
	UserId    int64
	ProjectId int64
	Data      map[string]any
}

type HookResult struct {
	NoOfHooksRan  int16
	Mutated       bool
	Error         error
	PreventAction bool
}

type HookEngine interface {
	Init() error
	Invalidate(pid int64) error
	Emit(e HookEvent) HookResult
	Stop(force bool) error
}
