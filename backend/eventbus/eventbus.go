package eventbus

import (
	"sync"

	hookengine "github.com/bornjre/turnix/backend/eventbus/hookEngine"
	"github.com/bornjre/turnix/backend/services/database"
	"github.com/bornjre/turnix/backend/services/signer"
	"github.com/bornjre/turnix/backend/xtypes/xbus"
	"github.com/bwmarrin/snowflake"
	"github.com/rs/zerolog"
)

type EventBus struct {
	handlers   map[string][]handlerRef
	hLock      sync.RWMutex
	hookEngine *hookengine.HookEngine
	snowflake  *snowflake.Node
}

type handlerRef struct {
	handler  xbus.EventHandler
	priority int16
}

func New(db *database.DB, signer *signer.Signer, logger zerolog.Logger) *EventBus {
	snode, err := snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}

	return &EventBus{
		handlers:   make(map[string][]handlerRef),
		hLock:      sync.RWMutex{},
		hookEngine: hookengine.New(db, signer, logger, snode),
		snowflake:  snode,
	}

}

func (e *EventBus) Init() error {
	return e.hookEngine.Init()
}

func (e *EventBus) Invalidate(pid int64) error {
	return e.hookEngine.Invalidate(pid)
}

func (e *EventBus) Emit(ev xbus.EventNew) (*xbus.EventResult, error) {

	e.hLock.RLock()

	handlers := e.handlers[ev.Type]

	e.hLock.RUnlock()

	ctx := xbus.EventContext{
		EventId:       e.snowflake.Generate().Int64(),
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

	if ev.Project == 0 {
		return &xbus.EventResult{
			PreventAction: ctx.PreventAction,
			Errors:        map[string]string{},
			Data:          ctx.Data,
		}, nil
	}

	return e.hookEngine.Emit(ctx)

}

func (e *EventBus) OnEvent(name string, handler xbus.EventHandler, priority int16) {

	e.hLock.Lock()
	defer e.hLock.Unlock()

	handlers := e.handlers[name]

	handlers = append(handlers, handlerRef{
		handler:  handler,
		priority: priority,
	})

	e.handlers[name] = handlers

}

func (e *EventBus) Stop(force bool) error {
	return nil
}
