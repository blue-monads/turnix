package assetserve

import (
	"github.com/bornjre/turnix/backend/registry"
	"github.com/bornjre/turnix/backend/xtypes/xproject"
)

func init() {
	registry.Register(&xproject.Defination{
		Name:                "AssetServe",
		Slug:                "assetserve",
		NewFormSchemaFields: []xproject.PTypeField{},
		Perminssions:        []string{},
		EventTypes:          []string{},
		Builder:             New,
	})

}

type AssetServe struct {
}

func (a *AssetServe) Init(pid int64) error   { return nil }
func (a *AssetServe) DeInit(pid int64) error { return nil }

func New(opt xproject.BuilderOption) (xproject.ProjectType, error) {

	as := &AssetServe{}

	return xproject.NewAdapter(as), nil
}
