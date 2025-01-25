package main

import "sync"

var (
	handlers = map[string]Handler{
		"run-cmd": HandleRunCommand,
		"ping":    HandlePing,
	}
	hLock = sync.Mutex{}
)

type Handler func(*WHContext) (interface{}, error)

func RegisterHandler(mtype string, handler Handler) {
	hLock.Lock()
	defer hLock.Unlock()
	handlers[mtype] = handler
}

// WHContext is the context for the worker handler
type WHContext struct {
	Packet  *Packet
	Service *AgentService
}
