package main

import (
	"context"
	"fmt"
	"log/slog"
	"strings"

	"github.com/coder/websocket"
	"github.com/k0kubun/pp"
)

// http://localhost:7703/z/project/simplerat/5/device/

func handleServiceHello(ctx *WHContext) (any, error) {

	roomId := ctx.GetAsInt("room_id")

	pp.Println("roomId", roomId)

	baseUrl := GetWSRoomUrl(ctx.Node, roomId)

	wctx := context.Background()

	c, _, err := websocket.Dial(wctx, baseUrl, nil)
	if err != nil {
		slog.Error("Error connecting to ws: ", slog.String("error", err.Error()))
		return nil, fmt.Errorf("failed to connect to ws: %v", err)
	}

	s := &ServiceHello{
		ws: c,
	}

	// fixme make list of running services

	go s.ReadLoop()
	go s.WriteLoop()

	return "", nil

}

func handleServiceShell(ctx *WHContext) (any, error) {

	roomId := ctx.GetAsInt("room_id")

	pp.Println("roomId", roomId)

	baseUrl := GetWSRoomUrl(ctx.Node, roomId)

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

	err = s.Run()
	if err != nil {
		return nil, fmt.Errorf("failed to run shell: %v", err)
	}

	return "", nil

}

func GetWSRoomUrl(node *AgentNode, roomId int) string {
	baseUrl := fmt.Sprintf("%sroom/join/%d?token=%s", node.baseURL, roomId, node.authSessionToken)

	if strings.HasPrefix(baseUrl, "https") {
		baseUrl = strings.Replace(baseUrl, "https", "wss", 1)
	} else {
		baseUrl = strings.Replace(baseUrl, "http", "ws", 1)
	}

	return baseUrl
}
