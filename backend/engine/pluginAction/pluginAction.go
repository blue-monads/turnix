package pluginaction

import "github.com/bornjre/turnix/backend/engine/pool"

type PluginAction struct {
	pool *pool.GojaPool
}

func New(pool *pool.GojaPool) *PluginAction {
	return &PluginAction{
		pool: pool,
	}
}
