package lazydata

import (
	"encoding/json"

	"github.com/dop251/goja"
	"github.com/temphia/temphia/code/backend/xtypes"
)

var (
	_ xtypes.LazyData = (*GojaData)(nil)
)

type GojaData struct {
	rt    *goja.Runtime
	value goja.Value
}

func NewGojaData(rt *goja.Runtime, value goja.Value) *GojaData {
	return &GojaData{
		rt:    rt,
		value: value,
	}
}

func (g *GojaData) AsJsonBytes() ([]byte, error) {
	var i any
	g.rt.ExportTo(g.value, &i)
	return json.Marshal(i)
}

func (g *GojaData) AsObject(target any) error {
	return g.rt.ExportTo(g.value, target)
}

func (g *GojaData) IsJsonBytes() bool { return false }
func (g *GojaData) IsObject() bool    { return false }
func (g *GojaData) Inner() any        { return nil }
