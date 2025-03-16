package luaz

import (
	"os"
	"time"

	"github.com/blue-monads/turnix/backend/xtypes"
	"github.com/blue-monads/turnix/backend/xtypes/xproject"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
	lua "github.com/yuin/gopher-lua"
)

type Luaz struct {
	pool   *LuaStatePool
	app    xtypes.App
	logger zerolog.Logger
	root   *os.Root

	// routesMap map[string]string
	// services  map[string]string
}

type Options struct {
	BuilderOpts   xproject.BuilderOption
	Code          string
	WorkingFolder string
}

func New(opts Options) *Luaz {
	lz := &Luaz{
		pool: nil,
	}

	pool := NewLuaStatePool(LuaStatePoolOptions{
		MaxSize: 100,
		Ttl:     time.Hour,
		InitFn: func() (*LuaH, error) {

			if opts.Code == "" {
				opts.Code = code
			}

			L := lua.NewState()
			err := L.DoString(opts.Code)
			if err != nil {
				return nil, err
			}

			lh := &LuaH{
				parent:  lz,
				L:       L,
				closers: []func() error{},
			}

			err = lh.registerModules()
			if err != nil {
				return nil, err
			}

			return lh, nil
		},
	})

	lz.pool = pool

	return lz
}

func (l *Luaz) Handle(ctx *gin.Context) {

	lh, err := l.pool.Get()
	if err != nil {
		return
	}

	lh.Handle(ctx)
}
