package hookengine

import (
	"sync"

	"github.com/bornjre/turnix/backend/engine/pool"
	"github.com/bornjre/turnix/backend/services/database"
	"github.com/bornjre/turnix/backend/services/signer"
	"github.com/bornjre/turnix/backend/xtypes/xengine"

	"github.com/rs/zerolog"
)

type HookEngine struct {
	db          *database.DB
	hookRunners map[int64]*hookRunner
	hrLock      sync.RWMutex
	gojaPool    *pool.GojaPool
	logger      zerolog.Logger
	signer      *signer.Signer
}

type Options struct {
	DB       *database.DB
	Signer   *signer.Signer
	Logger   zerolog.Logger
	GojaPool *pool.GojaPool
}

func New(opts Options) *HookEngine {

	hook := &HookEngine{
		db:          opts.DB,
		hookRunners: make(map[int64]*hookRunner),
		gojaPool:    pool.New(opts.Logger),
		hrLock:      sync.RWMutex{},
		logger:      opts.Logger,
		signer:      opts.Signer,
	}

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

func (h *HookEngine) Emit(e xengine.EventContext) (*xengine.EventResult, error) {

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
