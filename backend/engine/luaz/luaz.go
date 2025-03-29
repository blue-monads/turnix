package luaz

import (
	"os"
	"strings"
	"time"

	"github.com/blue-monads/turnix/backend/xtypes"
	"github.com/blue-monads/turnix/backend/xtypes/models"
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
	routes []models.Route

	// services  map[string]string
}

type Options struct {
	BuilderOpts   xproject.BuilderOption
	Code          string
	WorkingFolder string
	Manifest      *models.Manifest
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
	lz.routes = opts.Manifest.Routes

	return lz
}

func (l *Luaz) Handle(ctx *gin.Context, relativePath string) {

	lh, err := l.pool.Get()
	if err != nil {
		return
	}

	var defaultRoute *models.Route

	for _, route := range l.routes {
		if route.Name == "default" {
			defaultRoute = &route
			continue
		}

		if route.Type != "http" {
			continue
		}

		if route.Method != ctx.Request.Method {
			continue
		}

		isEqual, params := compareRoutes(route.Path, relativePath)
		if !isEqual {
			continue
		}

		lh.Handle(ctx, route.Handler, params)
		return
	}

	if defaultRoute != nil {
		lh.Handle(ctx, defaultRoute.Handler, EmptyParams)
	}

}

// /post/:post_id
// /post/:post_id/edit

var EmptyParams = map[string]string{}

func compareRoutes(rpattern string, path string) (bool, map[string]string) {

	if !strings.Contains(rpattern, ":") {
		return rpattern == path, EmptyParams
	}

	rpparts := strings.Split(rpattern, "/")
	pathparts := strings.Split(path, "/")

	if len(rpparts) != len(pathparts) {
		return false, nil
	}

	params := map[string]string{}

	for i, part := range rpparts {
		if part == pathparts[i] {
			continue
		}

		if !strings.HasPrefix(part, ":") {
			return false, nil
		}

		paramName := part[1:]
		paramValue := pathparts[i]

		if paramValue == "" {
			return false, nil
		}

		params[paramName] = paramValue
	}

	return true, params
}
