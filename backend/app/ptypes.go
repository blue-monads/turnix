package app

import (
	"bytes"
	"encoding/json"

	"github.com/bornjre/trunis/backend/xtypes/xproject"
)

type Def struct {
	Name       string `json:"name"`
	Slug       string `json:"slug"`
	Icon       string `json:"icon,omitempty"`
	IsExternal bool   `json:"is_external"`
}

func buildGlobalJS(projTypes []*xproject.TypeDefination) ([]byte, error) {
	data := make([]Def, 0)
	var buf bytes.Buffer

	for _, ptype := range projTypes {
		data = append(data, Def{
			Name:       ptype.Name,
			Slug:       ptype.Slug,
			Icon:       ptype.Icon,
			IsExternal: ptype.AssetData != nil,
		})
	}

	out, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	buf.Write([]byte(`window["__turnis_ptypes__"] = `))
	buf.Write(out)
	buf.WriteByte('\n')

	for _, ptype := range projTypes {
		buf.Write(ptype.GlobalJS)
		buf.WriteByte('\n')
	}

	return buf.Bytes(), nil

}
