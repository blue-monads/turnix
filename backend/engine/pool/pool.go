package pool

import (
	"sync"

	"github.com/dop251/goja"
	"github.com/rs/zerolog"
)

type GojaPool struct {
	logger zerolog.Logger
	pool   sync.Pool

	pidCacheIndex       map[int64]*GojaHandle
	currentRuntimeCount int64
	maxRuntimeCount     int64
	idexMLock           sync.Mutex
}

func New(logger zerolog.Logger) *GojaPool {

	return &GojaPool{
		logger:              logger,
		pool:                sync.Pool{},
		pidCacheIndex:       map[int64]*GojaHandle{},
		currentRuntimeCount: 0,
		maxRuntimeCount:     10,
		idexMLock:           sync.Mutex{},
	}
}

func (g *GojaPool) Get(pid int64, eid string, cacheOnly bool) *GojaHandle {

	g.logger.Info().Int64("pid", pid).Str("eid", eid).Msg("GojaPool.Get")

	g.idexMLock.Lock()
	defer g.idexMLock.Unlock()

	gh := g.pidCacheIndex[pid]
	if gh != nil || cacheOnly {
		g.logger.Info().Msg("GojaPool.Get/pidCacheIndex")
		delete(g.pidCacheIndex, pid)
		return gh
	}

	maybeGh := g.pool.Get()
	if maybeGh != nil {
		g.logger.Info().Msg("GojaPool.Get/pool")
		return maybeGh.(*GojaHandle)
	}

	if g.maxRuntimeCount <= g.currentRuntimeCount {
		g.logger.Info().Int64("current_count", g.currentRuntimeCount).Msg("GojaPool.Get/maxReached")
		return nil
	}

	g.currentRuntimeCount = g.currentRuntimeCount + 1

	g.logger.Info().Int64("current_count", g.currentRuntimeCount).Msg("GojaPool.Get/New")

	return &GojaHandle{
		JS:      goja.New(),
		LastPid: 0,
	}
}

func (g *GojaPool) Set(gh *GojaHandle) {
	g.idexMLock.Lock()
	defer g.idexMLock.Unlock()

	existing := g.pidCacheIndex[gh.LastPid]
	if existing == nil {
		g.pidCacheIndex[gh.LastPid] = gh
		return
	}

	g.pool.Put(gh)
}
