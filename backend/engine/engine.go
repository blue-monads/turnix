package hookengine

import (
	"sync"

	"github.com/bornjre/turnix/backend/services/database"
	"github.com/bornjre/turnix/backend/xtypes/services/xhook"
	"github.com/rs/zerolog"
)

type HookEngine struct {
	db          *database.DB
	hookRunners map[int64]*hookRunner
	hrLock      sync.RWMutex
	gojaPool    GojaPool
	logger      zerolog.Logger
}

func New(db *database.DB, logger zerolog.Logger) *HookEngine {
	return &HookEngine{
		db:          db,
		hookRunners: make(map[int64]*hookRunner),
		gojaPool:    newGojaPool(),
		hrLock:      sync.RWMutex{},
		logger:      logger,
	}
}

func (h *HookEngine) Init() error {

	h.logger.Info().Msg("Init")

	return nil
}

func (h *HookEngine) Invalidate(pid int64) error {

	h.logger.Info().Int64("pid", pid).Msg("Invalidate")

	h.hrLock.Lock()
	defer h.hrLock.Unlock()

	delete(h.hookRunners, pid)

	return nil
}

func (h *HookEngine) Emit(e xhook.Event) (*xhook.Result, error) {
	h.logger.Info().
		Int64("pid", e.ProjectId).
		Int64("uid", e.UserId).
		Msg("Emit")

	return h.emit(e)
}

func (h *HookEngine) Stop(force bool) error {
	return nil
}
