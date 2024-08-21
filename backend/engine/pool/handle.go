package pool

import (
	"time"

	"github.com/dop251/goja"
	"github.com/k0kubun/pp"
)

type GojaHandle struct {
	JS      *goja.Runtime
	LastPid int64
}

type FetchRequest struct{}

type FetchResponse struct{}

func (g *GojaHandle) Bind() {

	obj := g.JS.NewObject()

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

	g.JS.Set("runtime", obj)

}
