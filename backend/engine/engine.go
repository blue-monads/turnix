package hookengine

import (
	"sync"

	"github.com/bornjre/turnix/backend/services/database"
	"github.com/bornjre/turnix/backend/xtypes/services/xhook"
	"github.com/bwmarrin/snowflake"
	"github.com/rs/zerolog"
)

type HookEngine struct {
	db          *database.DB
	hookRunners map[int64]*hookRunner
	hrLock      sync.RWMutex
	gojaPool    GojaPool
	logger      zerolog.Logger
	snowflake   *snowflake.Node
}

func New(db *database.DB, logger zerolog.Logger) *HookEngine {
	snode, err := snowflake.NewNode(1)
	if err != nil {
		panic(err)
	}

	hook := &HookEngine{
		db:          db,
		hookRunners: make(map[int64]*hookRunner),
		gojaPool:    newGojaPool(),
		hrLock:      sync.RWMutex{},
		logger:      logger,
		snowflake:   snode,
	}

	hook.gojaPool.engine = hook

	return hook
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

	e.Id = h.snowflake.Generate().Int64()

	h.logger.Info().
		Int64("pid", e.ProjectId).
		Int64("uid", e.UserId).
		Msg("Emit")

	h.logger.Debug().Any("event", e).Msg("Emit/debug")

	return h.emit(e)
}

func (h *HookEngine) Stop(force bool) error {
	h.logger.Info().Bool("force", force).Msg("Stop")

	return nil
}
