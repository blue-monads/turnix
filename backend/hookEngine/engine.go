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
	return nil
}

func (h *HookEngine) Emit(e xtypes.HookEvent) xtypes.HookResult {

	return xtypes.HookResult{
		NoOfHooksRan: 0,
		Mutated:      false,
		Error:        nil,
	}
}

func (h *HookEngine) Stop(force bool) error {
	return nil
}
