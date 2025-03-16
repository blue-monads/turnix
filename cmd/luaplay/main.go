package main

import lua "github.com/yuin/gopher-lua"

const code = `
local handle = require("handle")
print("hello from lua")
print(tb.name)
print(handle.hello())
local j, b, c = handle.return_two_things()
print(j, b, c)
`

func main() {
	L := lua.NewState()

	// Create the table for global variable
	tb := L.NewTable()
	L.SetField(tb, "name", lua.LString("john"))
	L.SetGlobal("tb", tb)

	// Create handle module table
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

	// Register the module loader
	L.PreloadModule("handle", func(L *lua.LState) int {
		// Create a copy of the handleT table to return
		mod := L.NewTable()
		L.SetFuncs(mod, map[string]lua.LGFunction{
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
		L.Push(mod)
		return 1
	})

	defer L.Close()

	err := L.DoString(code)
	if err != nil {
		panic(err)
	}
}
