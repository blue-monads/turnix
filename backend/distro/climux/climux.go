package climux

import "sync"

var (
	DefaultCLI = "help"
)

type CLiFunc func(args []string) error

type Action struct {
	Name string
	Help string
	Func func(ctx Context) error
}

type Registery map[string]*Action

var (
	cliRegistery = map[string]*Action{}
	cLock        sync.Mutex
)

func Register(a *Action) {
	cLock.Lock()
	cliRegistery[a.Name] = a
	cLock.Unlock()
}

func GetRegistry() map[string]*Action {

	cLock.Lock()
	resp := map[string]*Action{}
	cLock.Unlock()

	for k, v := range cliRegistery {
		resp[k] = v
	}

	return cliRegistery
}
