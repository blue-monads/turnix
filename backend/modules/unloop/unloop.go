package unloop

import (
	"github.com/blue-monads/turnix/backend/xtypes/xproject"
)

var def = &xproject.Defination{
	Name: "Un Loop",
	Slug: "unloop",
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
