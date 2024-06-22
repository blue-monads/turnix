package xhook

type Event struct {
	Name      string
	UserId    int64
	ProjectId int64
	Data      map[string]any
}

type Result struct {
	NoOfHooksRan  int16
	Mutated       bool
	PreventAction bool
	Errors        map[string]string
}

type Engine interface {
	Init() error
	Invalidate(pid int64) error
	Emit(e Event) (*Result, error)
	Stop(force bool) error
}
