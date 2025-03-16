package main

import lua "github.com/yuin/gopher-lua"

const code = `

print("hello from lua")
print(tb.name)
print(handle.hello())
local j, b, c = handle.return_two_things()
print(j, b, c)

`

func main() {

	L := lua.NewState()
	tb := L.NewTable()

	L.SetField(tb, "name", lua.LString("john"))
	L.SetGlobal("tb", tb)

	handleT := L.NewTable()
	L.SetFuncs(handleT, map[string]lua.LGFunction{
		"hello": func(L *lua.LState) int {
			L.Push(lua.LString("hello from go"))
			return 1
		},
		"return_two_things": func(L *lua.LState) int {
			L.Push(lua.LString("jamba"))
			L.Push(lua.LString("juce"))
			return 2
		},
	})

	L.SetGlobal("handle", handleT)

	defer L.Close()

	err := L.DoString(code)
	if err != nil {
		panic(err)
	}

}
