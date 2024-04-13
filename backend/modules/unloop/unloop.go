package unloop

import (
	"github.com/bornjre/trunis/backend/xtypes/xproject"
)

var Def = xproject.Defination{
	Name: "On Loop",
	Slug: "onloop",
	Info: "Perform action with you on loop",
	Icon: "arrow-path",
	NewFormSchemaFields: []xproject.PTypeField{
		{
			Name:       "Init examples",
			KeyName:    "init_examples",
			Ftype:      "BOOL",
			IsRequired: false,
		},
	},
}
