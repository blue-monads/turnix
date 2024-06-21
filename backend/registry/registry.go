package registry

import (
	"sync"

	"github.com/bornjre/turnix/backend/xtypes/xproject"
)

var (
	registeredProjects = []*xproject.Defination{}
	rmutex             = sync.Mutex{}
	initFinished       = false
)

func Register(projdef *xproject.Defination) {
	rmutex.Lock()
	defer rmutex.Unlock()

	if initFinished {
		panic("Registered Projects are already being initilized, too late to register")

	}

	registeredProjects = append(registeredProjects, projdef)

}

func GetAll() []*xproject.Defination {
	rmutex.Lock()
	defer rmutex.Unlock()
	initFinished = true

	return registeredProjects

}
