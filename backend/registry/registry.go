package registry

import (
	"sync"

	"github.com/blue-monads/turnix/backend/xtypes/xproject"
)

var (
	registeredProjects = map[string]xproject.Builder{}
	rmutex             = sync.Mutex{}
	initFinished       = false
)

func Register(name string, builder xproject.Builder) {
	rmutex.Lock()
	defer rmutex.Unlock()

	if initFinished {
		panic("Registered Projects are already being initilized, too late to register")

	}
	registeredProjects[name] = builder

}

func GetAll() map[string]xproject.Builder {
	rmutex.Lock()
	defer rmutex.Unlock()
	initFinished = true

	return registeredProjects

}
