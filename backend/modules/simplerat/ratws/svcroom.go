package ratws

import (
	"context"
	"log/slog"
	"net/http"
	"sync"

	"github.com/coder/websocket"
	"github.com/gin-gonic/gin"
)

type WSServiceRoom struct {
	parent            *ECPWebsocket
	roomId            int64
	useId             int64
	deviceId          int64
	agentConnection   *websocket.Conn
	browserConnection *websocket.Conn
	roomClosed        bool
	browserClosed     bool
	agentClosed       bool
	lock              sync.Mutex
}

func (e *ECPWebsocket) AddRoom(userId, deviceId int64) *WSServiceRoom {
	e.srLock.Lock()
	defer e.srLock.Unlock()

	e.svcRoomIdCounter++
	roomId := e.svcRoomIdCounter

	room := &WSServiceRoom{
		roomId:            roomId,
		useId:             userId,
		deviceId:          deviceId,
		agentConnection:   nil,
		browserConnection: nil,
		roomClosed:        false,
		browserClosed:     true,
		agentClosed:       true,
		lock:              sync.Mutex{},
		parent:            e,
	}

	e.svcRooms[roomId] = room

	return room

}

func (e *ECPWebsocket) GetRoom(roomId int64) *WSServiceRoom {
	e.srLock.RLock()
	defer e.srLock.RUnlock()

	return e.svcRooms[roomId]
}

func (e *ECPWebsocket) RemoveRoom(roomId int64) {
	e.srLock.Lock()
	defer e.srLock.Unlock()

	room := e.svcRooms[roomId]

	if room == nil {
		return
	}

	room.lock.Lock()
	defer room.lock.Unlock()

	if room.agentConnection != nil && !room.agentClosed {
		room.agentConnection.CloseNow()
		room.agentClosed = true
	}

	if room.browserConnection != nil && !room.browserClosed {
		room.browserConnection.CloseNow()
		room.browserClosed = true
	}

	delete(e.svcRooms, roomId)
}

func (ws *WSServiceRoom) GetRoomId() int64 {
	return ws.roomId
}

func (ws *WSServiceRoom) AddAgentToRoom(agentId int64, ctx *gin.Context) bool {

	conn, err := websocket.Accept(ctx.Writer, ctx.Request, nil)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return false
	}

	ws.lock.Lock()
	defer ws.lock.Unlock()

	if agentId != ws.deviceId {
		return false
	}

	oldconn := ws.agentConnection
	oldClosed := ws.agentClosed

	ws.agentConnection = conn
	ws.agentClosed = false

	if oldconn != nil && !oldClosed {
		oldconn.CloseNow()
	}

	go ws.agentEventLoop()

	return true
}

func (ws *WSServiceRoom) AddBrowserToRoom(browserId int64, ctx *gin.Context) bool {

	conn, err := websocket.Accept(ctx.Writer, ctx.Request, nil)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return false
	}

	ws.lock.Lock()
	defer ws.lock.Unlock()

	if browserId != ws.useId {
		return false
	}

	oldconn := ws.browserConnection
	oldClosed := ws.browserClosed

	ws.browserConnection = conn
	ws.browserClosed = false

	if oldconn != nil && !oldClosed {
		oldconn.CloseNow()
	}

	go ws.browserEventLoop()

	return true
}

func (ws *WSServiceRoom) browserEventLoop() {
	ctx := context.Background()
	errCounter := 0

	defer func() {
		ws.lock.Lock()
		defer ws.lock.Unlock()

		if ws.browserConnection != nil && !ws.browserClosed {
			ws.browserConnection.CloseNow()
			ws.browserClosed = true
		}

		ws.parent.closeRoomChan <- ws.roomId

	}()

	for {

		if ws.roomClosed {
			return
		}

		if errCounter > 5 {
			slog.Warn("too many errors, closing ws")
			return
		}

		wsmt, data, err := ws.browserConnection.Read(ctx)
		if err != nil {
			errCounter++
			slog.Warn("error reading from ws browser 1", "err", err)
			continue
		}

		if ws.agentConnection != nil && !ws.agentClosed {
			err = ws.agentConnection.Write(ctx, wsmt, data)
			if err != nil {
				slog.Warn("error writing to ws", "err", err)
				continue
			}
		} else {
			slog.Warn("dropping message, no agent connection")
		}

		errCounter = 0

	}
}

func (ws *WSServiceRoom) agentEventLoop() {
	ctx := context.Background()
	errCounter := 0

	defer func() {
		ws.lock.Lock()
		defer ws.lock.Unlock()

		if ws.agentConnection != nil && !ws.agentClosed {
			ws.agentConnection.CloseNow()
			ws.agentClosed = true
		}
		ws.parent.closeRoomChan <- ws.roomId

	}()

	for {

		if ws.roomClosed {
			return
		}

		if errCounter > 5 {
			slog.Warn("too many errors, closing ws")
			return
		}

		wsmt, data, err := ws.agentConnection.Read(ctx)
		if err != nil {
			errCounter++
			slog.Warn("error reading from ws agent 1", "err", err)
			continue
		}

		if ws.browserConnection != nil && !ws.browserClosed {
			err = ws.browserConnection.Write(ctx, wsmt, data)
			if err != nil {
				slog.Warn("error writing to ws", "err", err)
				continue
			}
		} else {
			slog.Warn("dropping message, no browser connection")
		}

		errCounter = 0

	}
}
