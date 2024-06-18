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

type EventContext struct {
	ProjectId   int64             `json:"project_id"`
	EventId     int64             `json:"event_id"`
	RunasUserID int64             `json:"run_as_user"`
	ExtraMeta   map[string]string `json:"envs"`
}

type RunnerOptions struct {
	ProjectId   int64 `json:"project_id"`
	EventId     int64 `json:"event_id"`
	RunasUserID int64 `json:"run_as_user"`
	Envs        map[string]string
}
