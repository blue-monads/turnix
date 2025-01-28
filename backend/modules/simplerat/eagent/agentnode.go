package main

import (
	"context"
	"log/slog"
	"strings"
	"time"

	"github.com/blue-monads/turnix/backend/modules/simplerat/wire"
	"github.com/coder/websocket"

	"github.com/k0kubun/pp"
)

// read limited at 32769 bytes
type AgentNode struct {
	configHome string
	configFile string

	noOfWorkers int

	baseURL          string
	registerToken    string
	authSessionToken string
	authRefreshToken string

	onClose chan struct{}

	workersChan chan *wire.Packet
	writeLoopCh chan *Response
}

func (a *AgentNode) Run() error {

	slog.Info("Agent service started")

	pp.Println(a)

	for i := 0; i < a.noOfWorkers; i++ {
		go a.worker()
	}

	for {

		ok := a.startWS()

		if !ok {
			slog.Error("Error starting ws")
			time.Sleep(5 * time.Second)
			continue
		}

		<-a.onClose

		slog.Info("Agent service closed")

		return nil

	}

}

func (a *AgentNode) startWS() bool {

	ctx, cancel := context.WithTimeout(context.Background(), time.Minute)
	defer cancel()

	baseUrl := a.baseURL + "device-ws?token=" + a.authSessionToken

	if strings.HasPrefix(baseUrl, "http://") {
		baseUrl = strings.Replace(baseUrl, "http://", "ws://", 1)
	} else if strings.HasPrefix(baseUrl, "https://") {
		baseUrl = strings.Replace(baseUrl, "https://", "wss://", 1)
	}

	counter := 0

	for {

		if counter > 5 {
			return false
		}

		if counter > 0 {
			time.Sleep(5 * time.Second)
		}

		counter++

		c, _, err := websocket.Dial(ctx, baseUrl, nil)
		if err != nil {
			slog.Error("Error connecting to ws: ", slog.String("error", err.Error()))
			continue
		}

		c.SetReadLimit(1024 * 1024 * 5)

		go a.evLoop(c)

		return true

	}

}

func (a *AgentNode) evLoop(c *websocket.Conn) {
	closeCh := make(chan struct{})

	defer func() {
		c.CloseNow()
		a.onClose <- struct{}{}
	}()

	ctx := context.Background()

	go func() {
		defer func() {
			closeCh <- struct{}{}
		}()

		consecutiveErrors := 0

		for {

			if consecutiveErrors > 5 {
				return
			}

			_, out, err := c.Read(ctx)
			if err != nil {
				slog.Error("Error reading ws: ", slog.String("error", err.Error()))
				consecutiveErrors++
				continue
			}

			header := wire.Unpack(out)

			pkt := &wire.Packet{
				Header: header,
				Data:   out[header.Size():],
			}

			consecutiveErrors = 0

			pp.Println("@pkt", pkt.Header, len(pkt.Data))

			a.workersChan <- pkt

		}

	}()

	for {

		slog.Info("waiting for write loop")

		select {
		case <-closeCh:
			return
		case pkt := <-a.writeLoopCh:
			slog.Info("writing to ws", slog.Int("mid", int(pkt.Mid)), slog.Int("mtype", int(pkt.Mtype)))

			head := wire.Pack(&wire.Header{
				Ptype: pkt.Mtype,
				Mid:   pkt.Mid,
				Pid:   pkt.Pid,
				MType: "",
			})

			head = append(head, pkt.DataBinary...)

			err := c.Write(ctx, websocket.MessageBinary, head)
			if err != nil {
				slog.Error("Error writing to ws: ", slog.String("error", err.Error()))
				continue
			}

		}

	}

}
