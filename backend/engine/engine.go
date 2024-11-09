package engine

import (
	pluginengine "github.com/bornjre/turnix/backend/engine/pluginAction"
	"github.com/bornjre/turnix/backend/engine/pool"

	"github.com/bornjre/turnix/backend/services/database"
	"github.com/bornjre/turnix/backend/services/signer"
	"github.com/rs/zerolog"
)

type Engine struct {
	pluginEngine *pluginengine.PluginAction
}

func New(db *database.DB, signer *signer.Signer, logger zerolog.Logger) *Engine {

	gpool := pool.New(logger)

	return &Engine{
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
