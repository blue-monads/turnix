package ratws

import (
	"context"
	"log/slog"
	"net/http"

	"github.com/blue-monads/turnix/backend/modules/simplerat/wire"
	"github.com/gin-gonic/gin"

	"github.com/coder/websocket"
)

func (e *ECPWebsocket) HandleAgentWS(agentId int64, ctx *gin.Context) {

	c, err := websocket.Accept(ctx.Writer, ctx.Request, nil)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	c.SetReadLimit(1024 * 1024 * 5)

	e.acLock.Lock()
	defer e.acLock.Unlock()

	old := e.agentsConns[agentId]
	if old != nil {
		old.CloseNow()
	}

	e.agentsConns[agentId] = c

	go e.eventLoop(agentId, c)

}

func (e *ECPWebsocket) SendAgentMessage(ctx context.Context, agentId int64, mtype string, data []byte, wait bool) []byte {
	e.acLock.RLock()
	c := e.agentsConns[agentId]
	e.acLock.RUnlock()
	if c == nil {
		slog.Warn("no ws connection for agent", "agentId", agentId)

		return nil
	}

	respChan := make(chan *wire.Packet)
	counter := int64(0)

	if wait {
		e.pbLock.Lock()
		e.pbCounter++
		counter = e.pbCounter
		e.pendingBrowserRequests[e.pbCounter] = respChan
		e.pbLock.Unlock()

		defer func() {
			e.pbLock.Lock()
			delete(e.pendingBrowserRequests, e.pbCounter)
			e.pbLock.Unlock()
		}()

	}

	header := wire.Pack(&wire.Header{
		Mid:   uint64(counter),
		Ptype: wire.PtypeRequest,
		Pid:   0,
		MType: mtype,
	})

	header = append(header, data...)

	err := c.Write(ctx, websocket.MessageBinary, header)
	if err != nil {
		slog.Warn("error writing to ws", "err", err)
		return nil
	}

	if !wait {
		return nil
	}

	for {
		slog.Info("waiting for response")
		msg := <-respChan

		return msg.Data

	}

}

func (e *ECPWebsocket) RemoveAgentConn(agentId int64) {
	e.acLock.Lock()
	defer e.acLock.Unlock()

	c := e.agentsConns[agentId]
	if c != nil {
		c.CloseNow()
		delete(e.agentsConns, agentId)
	}
}

// private

func (e *ECPWebsocket) eventLoop(agentId int64, conn *websocket.Conn) {

	ctx := context.Background()

	defer e.RemoveAgentConn(agentId)

	errorCounter := 0

	for {

		if errorCounter > 5 {
			slog.Warn("too many errors, closing ws")
			return
		}

		_, out, err := conn.Read(ctx)
		if err != nil {
			slog.Warn("error reading from ws", "err", err)
			errorCounter++
			continue
		}

		var msg wire.Packet
		header := wire.Unpack(out)

		msg.Header = header
		msg.Data = out[header.Size():]

		e.pbLock.RLock()
		respChan := e.pendingBrowserRequests[int64(header.Mid)]
		e.pbLock.RUnlock()

		if respChan != nil {
			respChan <- &msg
		}

	}

}
