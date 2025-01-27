package main

import (
	"context"
	"fmt"
	"log/slog"
	"time"

	"github.com/coder/websocket"
	"github.com/k0kubun/pp"
)

type ServiceHello struct {
	ws *websocket.Conn
}

func (s *ServiceHello) ReadLoop() {
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

func (s *ServiceHello) WriteLoop() {
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
