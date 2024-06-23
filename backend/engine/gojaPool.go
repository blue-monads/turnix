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

func (g *GojaPool) Get(pid int64, cacheOnly bool) *gojaHandle {

	g.idexMLock.Lock()
	defer g.idexMLock.Unlock()

	gh := g.pidCacheIndex[pid]
	if gh != nil || cacheOnly {
		delete(g.pidCacheIndex, pid)
		return gh
	}

	maybeGh := g.pool.Get()
	if maybeGh != nil {
		return maybeGh.(*gojaHandle)
	}

	if g.maxRuntimeCount >= g.currentRuntimeCount {
		return nil
	}

	g.currentRuntimeCount = g.currentRuntimeCount + 1

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
