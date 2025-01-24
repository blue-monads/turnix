package eagent

type EAgent struct {
}

type Platform interface {
	FSListDir(path string) ([]string, error)
	FSReadFile(path string) ([]byte, error)
	FSWriteFile(path string, data []byte) error
	FSRemove(path string) error
	FSMkdir(path string) error
	FSRename(old string, new string) error

	SystemExec(command string) ([]byte, error)
	SystemStartShell(cmd string, submitFunc func([]byte) error) (chan []byte, error)
}

type ShellHandle struct {
	Id string
}

type Shell interface {
	ShellWrite(data []byte) error
	Close() error
}

type AgentHandle struct {
	ActiveShells map[string]*ShellHandle
}
