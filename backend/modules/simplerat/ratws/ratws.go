package ratws

import (
	"sync"
	"time"

	"github.com/coder/websocket"
	"github.com/k0kubun/pp"
)

type ECPWebsocket struct {
	pid              int64
	svcRooms         map[int64]*WSServiceRoom
	svcRoomIdCounter int64
	srLock           sync.RWMutex

	agentsConns            map[int64]*websocket.Conn
	acLock                 sync.RWMutex
	pendingBrowserRequests map[int64]chan *Message
	pbLock                 sync.RWMutex
	pbCounter              int64
}

func NewECPWebsocket(pid int64) *ECPWebsocket {

	return &ECPWebsocket{
		pid:              pid,
		svcRooms:         make(map[int64]*WSServiceRoom),
		svcRoomIdCounter: 1,
		srLock:           sync.RWMutex{},

		agentsConns: make(map[int64]*websocket.Conn),

		acLock:                 sync.RWMutex{},
		pendingBrowserRequests: make(map[int64]chan *Message),
		pbLock:                 sync.RWMutex{},
		pbCounter:              1,
	}
}

func (e *ECPWebsocket) Run() {

	for {
		time.Sleep(5 * time.Second)

		pp.Println("@active rooms", len(e.svcRooms))
		pp.Println("@active agents", len(e.agentsConns))

		for conn := range e.agentsConns {
			pp.Println("@active agent", conn)
		}

	}

}
