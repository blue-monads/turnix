package eventbus

import (
	"sync"

	"github.com/bornjre/turnix/backend/xtypes/xbus"
)

type EventBus struct {
	handlers map[string]map[int64][]func(ev *xbus.EventContext) error
	wLock    sync.RWMutex
}

func New() *EventBus {
	return &EventBus{}

}
