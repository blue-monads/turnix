package engine

import (
	"github.com/bornjre/turnix/backend/engine/pool"

	"encoding/json"
	"fmt"

	"github.com/bornjre/turnix/backend/services/database"
	"github.com/dop251/goja"
)

type Engine struct {
	pool *pool.GojaPool
	db   *database.DB
}

func New(pool *pool.GojaPool) *Engine {
	return &Engine{
		pool: pool,
	}
}

func (p *Engine) Run(pid, plugId int64, name string, data map[string]any) (any, error) {

	jsHandle := p.pool.Get(pid, true)

	if jsHandle == nil {
		return nil, fmt.Errorf("no goja handle found")
	}

	defer p.pool.Set(jsHandle)

	if jsHandle.LastPid != pid {

		plugin, err := p.db.GetProjectPlugin(0, pid, plugId)
		if err != nil {
			return nil, err
		}

		pgm, err := goja.Compile(fmt.Sprintf("project_plugin_%d", pid), plugin.ServerCode, true)
		if err != nil {
			return nil, err
		}

		jsHandle.LastPid = 0
		_, err = jsHandle.JS.RunProgram(pgm)
		if err != nil {
			return nil, err
		}

		jsHandle.Bind()
		jsHandle.JS.SetFieldNameMapper(goja.TagFieldNameMapper("json", true))

		jsHandle.LastPid = pid
	}

	var entry func(ctx *goja.Object)

	eval := jsHandle.JS.Get(name)
	if eval == nil {
		return nil, fmt.Errorf("%s function not found in script", name)
	}

	err := jsHandle.JS.ExportTo(eval, &entry)
	if err != nil {
		return nil, err
	}

	obj := buildEventObject(jsHandle.JS, data)

	entry(obj)

	return nil, nil
}

func (p *Engine) LiveRun(pid int64, name, source string, data map[string]any) (any, error) {

	rt := goja.New()

	_, err := rt.RunString(source)
	if err != nil {
		return nil, err
	}

	var entry func(ctx *goja.Object)

	eval := rt.Get(name)
	if eval == nil {
		return nil, fmt.Errorf("%s function not found in script", name)
	}

	err = rt.ExportTo(eval, &entry)
	if err != nil {
		return nil, err
	}

	obj := buildEventObject(rt, data)

	entry(obj)

	return nil, nil

}

func buildEventObject(jsEngine *goja.Runtime, data map[string]any) *goja.Object {
	obj := jsEngine.NewObject()

	obj.Set("dataAsObject", func() any {

		return data
	})

	obj.Set("dataAsJsonString", func() any {

		out, err := json.Marshal(data)
		if err != nil {
			return nil
		}

		return string(out)
	})

	return obj
}
