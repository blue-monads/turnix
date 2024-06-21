package app

import (
	"bytes"
	"encoding/json"
	"fmt"

	"github.com/bornjre/turnix/backend/xtypes/models"
	"github.com/bornjre/turnix/backend/xtypes/xproject"
)

func buildGlobalJS(projTypes []*xproject.Defination) ([]byte, error) {
	data := make([]models.ProjectTypes, 0)
	var buf bytes.Buffer

	for _, ptype := range projTypes {
		data = append(data, models.ProjectTypes{
			Name:       ptype.Name,
			Slug:       ptype.Slug,
			Icon:       ptype.Icon,
			Ptype:      ptype.Slug,
			Info:       ptype.Info,
			Link:       fmt.Sprintf("/z/pages/portal/projects/%s", ptype.Slug),
			IsExternal: ptype.AssetData != nil,
		})
	}

	out, err := json.Marshal(data)
	if err != nil {
		return nil, err
	}

	buf.Write([]byte(`window["__turnix_ptypes__"] = `))
	buf.Write(out)
	buf.WriteByte('\n')

	for _, ptype := range projTypes {
		buf.Write(ptype.GlobalJS)
		buf.WriteByte('\n')
	}

	return buf.Bytes(), nil

}
