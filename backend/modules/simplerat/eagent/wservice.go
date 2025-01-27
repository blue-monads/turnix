package main

import (
	"context"
	"fmt"
	"log/slog"
	"strings"
	"time"

	"github.com/coder/websocket"
	"github.com/k0kubun/pp"
)

// http://localhost:7703/z/project/simplerat/5/device/

func handleServiceJoinRoom(ctx *WHContext) (any, error) {

	roomId := ctx.GetAsInt("room_id")

	pp.Println("roomId", roomId)

	baseUrl := fmt.Sprintf("%sroom/join/%d?token=%s", ctx.Node.baseURL, roomId, ctx.Node.authSessionToken)

	if strings.HasPrefix(baseUrl, "https") {
		baseUrl = strings.Replace(baseUrl, "https", "wss", 1)
	} else {
		baseUrl = strings.Replace(baseUrl, "http", "ws", 1)
	}

	pp.Println("baseUrl", baseUrl)

	wctx := context.Background()

	c, _, err := websocket.Dial(wctx, baseUrl, nil)
	if err != nil {
		slog.Error("Error connecting to ws: ", slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to connect to ws: %v", err)
	}

	s := &ServiceShell{
		node: ctx.Node,
		ws:   c,
	}

	// fixme make list of running services

	go s.ReadLoop()
	go s.WriteLoop()

	return "", nil

}

type ServiceShell struct {
	node *AgentNode
	ws   *websocket.Conn
}

func (s *ServiceShell) ReadLoop() {
	ctx := context.Background()

	for {

		_, out, err := s.ws.Read(ctx)
		if err != nil {
			slog.Error("Error reading message: ", slog.String("error", err.Error()))
			continue
		}

		msg := string(out)

		pp.Println("msg", msg)

	}

}

func (s *ServiceShell) WriteLoop() {
	ctx := context.Background()

	counter := 0

	for {
		time.Sleep(5 * time.Second)

		msg := []byte(fmt.Sprintf("ping %d", counter))

		counter++

		err := s.ws.Write(ctx, websocket.MessageText, msg)
		if err != nil {
			slog.Error("Error writing message: ", slog.String("error", err.Error()))
			continue
		}

	}
}
