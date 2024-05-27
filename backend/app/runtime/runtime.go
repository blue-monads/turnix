package runtime

import "github.com/bornjre/trunis/backend/xtypes/models"

type Runtime struct {
	hooksCache map[int64][]models.ProjectHook
}

func New() *Runtime {
	return &Runtime{
		hooksCache: map[int64][]models.ProjectHook{},
	}
}

type AppRuntime interface {
	EmitEvent(name string, data any) error
}
