package books

import (
	"github.com/bornjre/trunis/backend/registry"
	"github.com/bornjre/trunis/backend/xtypes/xproject"
)

var defs = &xproject.Defination{
	Name:                "Books",
	Slug:                "books",
	Info:                "Accounting made simpler",
	Icon:                "book-open",
	NewFormSchemaFields: []xproject.PTypeField{},
}

func init() {
	registry.Register(defs)
}
