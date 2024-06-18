package hookengine

import (
	"sync"

	"github.com/bornjre/trunis/backend/services/database"
	"github.com/bornjre/trunis/backend/xtypes"
)

type HookEngine struct {
	db          *database.DB
	hookRunners map[int64]*hookRunner
	hrLock      sync.RWMutex
	gojaPool    GojaPool
}

func New(db *database.DB) *HookEngine {
	return &HookEngine{
		db:          db,
		hookRunners: make(map[int64]*hookRunner),
		gojaPool:    newGojaPool(),
		hrLock:      sync.RWMutex{},
	}
}

func (h *HookEngine) Init() error {

	return nil
}

func (h *HookEngine) Invalidate(pid int64) error {

	h.hrLock.Lock()
	defer h.hrLock.Unlock()

	delete(h.hookRunners, pid)

	return nil
}

func (h *HookEngine) Emit(e xtypes.HookEvent) xtypes.HookResult {
	return h.emit(e)
}

func (h *HookEngine) Stop(force bool) error {
	return nil
}
