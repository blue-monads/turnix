package hookengine

import (
	"sync"

	"github.com/bornjre/turnix/backend/services/database"
	"github.com/bornjre/turnix/backend/services/signer"
	"github.com/bornjre/turnix/backend/xtypes/xbus"
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
	signer      *signer.Signer
}

func New(db *database.DB, signer *signer.Signer, logger zerolog.Logger, snowflake *snowflake.Node) *HookEngine {

	hook := &HookEngine{
		db:          db,
		hookRunners: make(map[int64]*hookRunner),
		gojaPool:    newGojaPool(),
		hrLock:      sync.RWMutex{},
		logger:      logger,
		signer:      signer,
		snowflake:   snowflake,
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

func (h *HookEngine) Emit(e xbus.EventContext) (*xbus.EventResult, error) {

	h.logger.Info().
		Int64("pid", e.Event.Project).
		Int64("uid", e.Event.UserId).
		Msg("Emit")

	h.logger.Debug().Any("event", e).Msg("Emit/debug")

	return h.emit(e)
}

func (h *HookEngine) Stop(force bool) error {
	h.logger.Info().Bool("force", force).Msg("Stop")

	return nil
}
