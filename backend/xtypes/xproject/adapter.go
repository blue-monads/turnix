package xproject

type Adapter struct {
	core Core
}

func NewAdapter(core Core) *Adapter {
	return &Adapter{
		core: core,
	}
}

func (a *Adapter) Init(pid int64) error {
	return a.core.Init(pid)
}

func (a *Adapter) IsInitilized(pid int64) (bool, error) {
	return true, nil
}

func (a *Adapter) DeInit(pid int64) error {
	return a.core.DeInit(pid)
}

func (a *Adapter) OnFileEvent(event *FileEvent) error {
	return nil
}

func (a *Adapter) OnUserEvent(event *UserEvent) error {
	return nil
}

type Core interface {
	Init(pid int64) error
	DeInit(pid int64) error
}
