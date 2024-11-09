package engine

import (
	"sync"

	pluginengine "github.com/bornjre/turnix/backend/engine/pluginAction"
	"github.com/bornjre/turnix/backend/engine/pool"

	"github.com/bornjre/turnix/backend/services/database"
	"github.com/bornjre/turnix/backend/services/signer"
	"github.com/bornjre/turnix/backend/xtypes/xengine"
	"github.com/rs/zerolog"
)

type Engine struct {
	hLock        sync.RWMutex
	pluginEngine *pluginengine.PluginAction
}

type handlerRef struct {
	handler  xengine.EventHandler
	priority int16
}

func New(db *database.DB, signer *signer.Signer, logger zerolog.Logger) *Engine {

	gpool := pool.New(logger)

	return &Engine{
		hLock:        sync.RWMutex{},
		pluginEngine: pluginengine.New(gpool),
	}

}

func (e *Engine) Init() error {

	return nil
}

func (e *Engine) Invalidate(pid int64) error {
	return nil
}

func (e *Engine) Stop(force bool) error {
	return nil
}
