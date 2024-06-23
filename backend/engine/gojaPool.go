package hookengine

import (
	"sync"

	"github.com/dop251/goja"
)

type GojaPool struct {
	engine *HookEngine

	pool sync.Pool

	pidCacheIndex       map[int64]*gojaHandle
	currentRuntimeCount int64
	maxRuntimeCount     int64
	idexMLock           sync.Mutex
}

func newGojaPool() GojaPool {

	return GojaPool{
		pool:                sync.Pool{},
		pidCacheIndex:       map[int64]*gojaHandle{},
		engine:              nil,
		currentRuntimeCount: 0,
		maxRuntimeCount:     10,
		idexMLock:           sync.Mutex{},
	}
}

func (g *GojaPool) Get(pid, eid int64, cacheOnly bool) *gojaHandle {

	g.engine.logger.Info().Int64("pid", pid).Int64("eid", eid).Msg("GojaPool.Get")

	g.idexMLock.Lock()
	defer g.idexMLock.Unlock()

	gh := g.pidCacheIndex[pid]
	if gh != nil || cacheOnly {
		g.engine.logger.Info().Msg("GojaPool.Get/pidCacheIndex")
		delete(g.pidCacheIndex, pid)
		return gh
	}

	maybeGh := g.pool.Get()
	if maybeGh != nil {
		g.engine.logger.Info().Msg("GojaPool.Get/pool")
		return maybeGh.(*gojaHandle)
	}

	if g.maxRuntimeCount <= g.currentRuntimeCount {
		g.engine.logger.Info().Int64("current_count", g.currentRuntimeCount).Msg("GojaPool.Get/maxReached")
		return nil
	}

	g.currentRuntimeCount = g.currentRuntimeCount + 1

	g.engine.logger.Info().Int64("current_count", g.currentRuntimeCount).Msg("GojaPool.Get/New")

	return &gojaHandle{
		js:      goja.New(),
		lastPid: 0,
	}
}

func (g *GojaPool) Set(pid int64, gh *gojaHandle) {
	g.idexMLock.Lock()
	defer g.idexMLock.Unlock()

	existing := g.pidCacheIndex[pid]
	if existing == nil {
		g.pidCacheIndex[pid] = gh
		return
	}

	g.pool.Put(gh)
}
