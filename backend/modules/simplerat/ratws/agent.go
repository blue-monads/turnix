package ratws

import (
	"context"
	"encoding/json"
	"log/slog"
	"net/http"

	"github.com/gin-gonic/gin"

	"github.com/coder/websocket"
)

func (e *ECPWebsocket) HandleAgentWS(agentId int64, ctx *gin.Context) {

	c, err := websocket.Accept(ctx.Writer, ctx.Request, nil)
	if err != nil {
		ctx.AbortWithError(http.StatusInternalServerError, err)
		return
	}

	e.acLock.Lock()
	defer e.acLock.Unlock()

	old := e.agentsConns[agentId]
	if old != nil {
		old.CloseNow()
	}

	e.agentsConns[agentId] = c

	go e.eventLoop(agentId, c)

}

type Message struct {
	MessageId int64           `json:"mid"`
	Data      json.RawMessage `json:"data"`
	AgentId   int64           `json:",omitempty"`
}

func (e *ECPWebsocket) SendAgentMessage(ctx context.Context, agentId int64, data []byte) []byte {
	e.acLock.RLock()
	c := e.agentsConns[agentId]
	e.acLock.RUnlock()
	if c == nil {
		slog.Warn("no ws connection for agent", "agentId", agentId)

		return nil
	}

	respChan := make(chan *Message)
	counter := int64(0)

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

	msg := &Message{
		MessageId: counter,
		Data:      data,
		AgentId:   agentId,
	}

	jdata, err := json.Marshal(msg)
	if err != nil {
		slog.Warn("error marshalling message", "err", err)
		return nil
	}

	err = c.Write(ctx, websocket.MessageText, jdata)
	if err != nil {
		slog.Warn("error writing to ws", "err", err)
		return nil
	}

	for {
		slog.Info("waiting for response")
		msg := <-respChan

		if msg.AgentId == agentId {
			return msg.Data
		}

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

		select {
		case <-ctx.Done():
			return
		default:
			_, data, err := conn.Read(ctx)
			if err != nil {
				errorCounter++
				slog.Warn("error reading from ws 1", "err", err)
				continue
			}

			errorCounter = 0

			var msg Message
			err = json.Unmarshal(data, &msg)
			if err != nil {
				slog.Warn("error reading from ws 2", "err", err)
				continue
			}

			msg.AgentId = agentId

			e.pbLock.RLock()
			respChan := e.pendingBrowserRequests[msg.MessageId]
			e.pbLock.RUnlock()

			if respChan != nil {
				respChan <- &msg
			}

		}

	}

}
