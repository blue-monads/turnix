package server

import (
	"fmt"
	"io"
	"regexp"
	"strings"
	"sync"

	"github.com/bornjre/turnix/backend/engine/pool"
	"github.com/bornjre/turnix/backend/utils/libx/httpx"
	"github.com/bornjre/turnix/backend/xtypes"
	"github.com/bornjre/turnix/backend/xtypes/models"
	"github.com/dop251/goja"
	"github.com/gin-gonic/gin"
)

type Injector struct {
	server           *Server
	hashooksIndex    map[string]bool
	hasHookIndexLock sync.RWMutex

	handles map[int64]HookHandle
	hLock   sync.RWMutex

	pool *pool.GojaPool
}

type HookHandle struct {
	hooks   []models.ProjectHook
	program *goja.Program
}

var regex = regexp.MustCompile(`const handle = \((.*?)\) => {`)

func (s *Server) AsApiAction(name string, fn xtypes.ApiHandler) gin.HandlerFunc {
	return s.injector.AsApiAction(name, fn)
}

func (sh *Injector) AsApiAction(name string, fn xtypes.ApiHandler) gin.HandlerFunc {

	return func(ctx *gin.Context) {

		claim, err := sh.server.withAccess(ctx)
		if err != nil {
			return
		}

		pid := ctx.Param("pid")
		key := fmt.Sprintf("%s_%s", pid, name)

		sh.hLock.RLock()
		ok := sh.hashooksIndex[key]
		sh.hLock.RUnlock()

		ctxplus := xtypes.ContextPlus{
			Claim: claim,
			Http:  ctx,
		}

		if ok {
			resp, err := sh.runHook(ctxplus)
			_, exist := ctx.Get("exit")
			if exist {
				httpx.WriteJSON(ctx, resp, err)
				return
			}
		}

		resp, err := fn(ctxplus)

		httpx.WriteJSON(ctx, resp, err)
	}
}

func (sh *Injector) runHook(ctx xtypes.ContextPlus) (any, error) {

	pid := ctx.ProjectId()

	jsHandle := sh.pool.Get(pid, false)
	if jsHandle == nil {
		return nil, fmt.Errorf("could not accure JS runtime")
	}

	handle, ok := sh.handles[pid]

	if jsHandle.LastPid != pid {

		if !ok {

			hooks, err := sh.server.db.ListProjectHooks(pid)
			if err != nil {
				return nil, err
			}

			stringsBuf := strings.Builder{}

			for _, hook := range hooks {
				if hook.HookType != "script" {
					continue
				}

				updatedCode := regex.ReplaceAllString(hook.HookCode, fmt.Sprintf("const _handle_%d = ($1) => {", hook.ID))
				stringsBuf.WriteString(updatedCode)
			}

			code := stringsBuf.String()

			program, err := goja.Compile(fmt.Sprintf("project_hook_%d", pid), code, true)
			if err != nil {
				return nil, err
			}

			handle = HookHandle{
				hooks:   hooks,
				program: program,
			}

			sh.hLock.Lock()
			sh.handles[pid] = handle
			sh.hLock.Unlock()

		}

		jsHandle.JS.RunProgram(handle.program)

		jsHandle.Bind()
		jsHandle.JS.SetFieldNameMapper(goja.TagFieldNameMapper("json", true))

		jsHandle.LastPid = pid

	}

	for _, hook := range handle.hooks {
		if hook.HookType != "script" {
			continue
		}

		var entry func(ctx *goja.Object)

		eval := jsHandle.JS.Get(hook.Name)
		if eval == nil {
			return nil, fmt.Errorf("%s function not found in script", hook.Name)
		}

		err := jsHandle.JS.ExportTo(eval, &entry)
		if err != nil {
			return nil, err
		}

		obj := buildEventObject(jsHandle.JS, ctx.Http)

		entry(obj)
	}

	return nil, nil
}

func (sh *Injector) loadHooks() error {
	hooks, err := sh.server.db.ListAllProjectHooks()
	if err != nil {
		return err
	}

	sh.hLock.Lock()
	defer sh.hLock.Unlock()

	pool.New(sh.server.rootLogger)

	for _, hook := range hooks {
		key := fmt.Sprintf("%s_%d", hook.Event, hook.ProjectID)
		sh.hashooksIndex[key] = true
	}

	return nil

}

func buildEventObject(jsEngine *goja.Runtime, http *gin.Context) *goja.Object {
	obj := jsEngine.NewObject()

	obj.Set("dataAsObject", func() any {
		var data map[string]any
		err := http.BindJSON(&data)
		if err != nil {
			return nil
		}

		return data
	})

	obj.Set("dataAsJsonString", func() any {

		out, err := io.ReadAll(http.Request.Body)
		if err != nil {
			return nil
		}

		return string(out)
	})

	return obj
}
