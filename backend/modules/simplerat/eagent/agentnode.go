package main

import (
	"context"
	"log/slog"
	"strings"
	"time"

	"github.com/coder/websocket"
	"github.com/coder/websocket/wsjson"
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

	workersChan chan *Packet
	writeLoopCh chan *Packet
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

type Packet struct {
	MessageId int64   `json:"mid"`
	Data      Message `json:"data"`
}

type Message struct {
	MType string `json:"mtype"`
	Data  any    `json:"data"`
}

func (a *AgentNode) evLoop(c *websocket.Conn) {
	closeCh := make(chan struct{})

	defer func() {
		c.CloseNow()
		a.onClose <- struct{}{}
	}()

	go func() {
		defer func() {
			closeCh <- struct{}{}
		}()

		consecutiveErrors := 0

		for {

			if consecutiveErrors > 5 {
				return
			}

			var pkt Packet
			err := wsjson.Read(context.Background(), c, &pkt)
			if err != nil {
				slog.Error("Error reading ws: ", slog.String("error", err.Error()))
				consecutiveErrors++
				return
			}

			consecutiveErrors = 0

			pp.Println("@pkt", pkt)

			a.workersChan <- &pkt

		}

	}()

	for {

		slog.Info("waiting for write loop")

		select {
		case <-closeCh:
			return
		case pkt := <-a.writeLoopCh:
			slog.Info("writing to ws", "pkt", pkt)

			err := wsjson.Write(context.Background(), c, pkt)
			if err != nil {
				slog.Error("Error writing ws: ", slog.String("error", err.Error()))
				continue
			}

		}

	}

}
