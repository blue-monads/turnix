package hookengine

import (
	"time"

	"github.com/dop251/goja"
	"github.com/k0kubun/pp"
)

type gojaHandle struct {
	js      *goja.Runtime
	lastPid int64
}

type FetchRequest struct{}

type FetchResponse struct{}

func (g *gojaHandle) bind() {

	obj := g.js.NewObject()

	obj.Set("sleep", func(ms int32) {
		time.Sleep(time.Millisecond * time.Duration(ms))
	})

	obj.Set("fetchResolve", func(reqs []FetchRequest) any {

		return nil

	})

	obj.Set("logInfo", func(msg string, opts map[string]any) {
		pp.Println(msg, opts)
	})

	obj.Set("logWarn", func(msg string, opts map[string]any) {
		pp.Println(msg, opts)
	})

	obj.Set("logDebug", func(msg string, opts map[string]any) {
		pp.Println(msg, opts)
	})

	g.js.Set("runtime", obj)

}
