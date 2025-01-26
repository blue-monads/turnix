package main

import (
	"github.com/k0kubun/pp"
)

func handleServiceJoinRoom(ctx *WHContext) (any, error) {

	roomId := ctx.GetAsInt("room_id")

	pp.Println("roomId", roomId)

	return "pong", nil

}
