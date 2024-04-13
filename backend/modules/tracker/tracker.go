package tracker

import (
	"github.com/bornjre/trunis/backend/registry"
	"github.com/bornjre/trunis/backend/xtypes/xproject"
)

var defs = &xproject.Defination{
	Name:                "Trackers",
	Slug:                "tracker",
	Info:                "Track everything",
	Icon:                "map-pin",
	NewFormSchemaFields: []xproject.PTypeField{},
}

func init() {
	registry.Register(defs)
}
