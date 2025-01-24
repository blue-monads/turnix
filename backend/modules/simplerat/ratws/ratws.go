package ratws

import (
	"sync"

	"github.com/coder/websocket"
)

type ECPWebsocket struct {
	pid              int64
	svcRooms         map[int64]*WSServiceRoom
	svcRoomIdCounter int64

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
		agentsConns:      make(map[int64]*websocket.Conn),
		svcRoomIdCounter: 0,
		acLock:           sync.RWMutex{},
	}
}
