package luaz

import (
	"github.com/gin-gonic/gin"
	lua "github.com/yuin/gopher-lua"
)

const code = `
function handler(ctx)
  print("Hello from lua!", ctx.type())
end

`

type LuaH struct {
	parent  *Luaz
	closers []func() error
	L       *lua.LState
}

func (l *LuaH) Close() error {
	for _, c := range l.closers {
		c()
	}

	l.closers = l.closers[:0]

	return nil
}

func (l *LuaH) Handle(ctx *gin.Context) {
	handler := l.L.GetGlobal("handler")
	ctxt := l.L.NewTable()

	l.L.SetFuncs(ctxt, map[string]lua.LGFunction{
		"request": func(l *lua.LState) int {
			return 0
		},
		"type": func(l *lua.LState) int {
			l.Push(lua.LString("http"))
			return 1
		},
	})

	l.L.Push(handler)
	l.L.Push(ctxt)
	l.L.Call(1, 0)

}
