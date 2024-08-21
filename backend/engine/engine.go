package engine

import (
	"sync"

	hookengine "github.com/bornjre/turnix/backend/engine/hookEngine"
	pluginengine "github.com/bornjre/turnix/backend/engine/pluginAction"
	"github.com/bornjre/turnix/backend/engine/pool"

	"github.com/bornjre/turnix/backend/services/database"
	"github.com/bornjre/turnix/backend/services/signer"
	"github.com/bornjre/turnix/backend/xtypes/xengine"
	"github.com/rs/xid"
	"github.com/rs/zerolog"
)

type Engine struct {
	handlers     map[string][]handlerRef
	hLock        sync.RWMutex
	hookEngine   *hookengine.HookEngine
	pluginEngine *pluginengine.PluginAction
}

type handlerRef struct {
	handler  xengine.EventHandler
	priority int16
}

func New(db *database.DB, signer *signer.Signer, logger zerolog.Logger) *Engine {

	gpool := pool.New(logger)

	return &Engine{
		handlers: make(map[string][]handlerRef),
		hLock:    sync.RWMutex{},
		hookEngine: hookengine.New(hookengine.Options{
			DB:       db,
			Signer:   signer,
			Logger:   logger,
			GojaPool: gpool,
		}),
		pluginEngine: pluginengine.New(gpool),
	}

}

func (e *Engine) Init() error {
	return e.hookEngine.Init()
}

func (e *Engine) Invalidate(pid int64) error {
	return e.hookEngine.Invalidate(pid)
}

func (e *Engine) Emit(ev xengine.EventNew) (*xengine.EventResult, error) {

	e.hLock.RLock()

	handlers := e.handlers[ev.Type]

	e.hLock.RUnlock()

	ctx := xengine.EventContext{
		EventId:       xid.New().String(),
		Event:         &ev,
		PreventAction: false,
		Data:          map[string]any{},
	}

	for _, hr := range handlers {
		err := hr.handler(ctx)
		if err != nil {
			return nil, err
		}
	}

	if ev.Project == 0 || !ev.AllowHook {
		return &xengine.EventResult{
			PreventAction: ctx.PreventAction,
			Errors:        map[string]string{},
			Data:          ctx.Data,
		}, nil
	}

	return e.hookEngine.Emit(ctx)

}

func (e *Engine) OnEvent(name string, handler xengine.EventHandler, priority int16) {

	e.hLock.Lock()
	defer e.hLock.Unlock()

	handlers := e.handlers[name]

	handlers = append(handlers, handlerRef{
		handler:  handler,
		priority: priority,
	})

	e.handlers[name] = handlers

}

func (e *Engine) Stop(force bool) error {
	return nil
}
