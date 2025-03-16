package luaz

import (
	"time"

	"github.com/blue-monads/turnix/backend/xtypes"
	"github.com/blue-monads/turnix/backend/xtypes/xproject"
	"github.com/gin-gonic/gin"
	"github.com/rs/zerolog"
)

type Luaz struct {
	pool   *LuaStatePool
	app    xtypes.App
	logger zerolog.Logger
}

func New(opts xproject.BuilderOption) *Luaz {

	pool := NewLuaStatePool(LuaStatePoolOptions{
		MaxSize: 100,
		Ttl:     time.Hour,
		InitFn: func() (*LuaH, error) {

			return nil, nil
		},
	})

	l := &Luaz{
		pool: pool,
	}

	return l
}

func (l *Luaz) Handle(ctx *gin.Context) {

}
