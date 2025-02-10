package main

import (
	"encoding/json"
	"fmt"
	"log/slog"

	"github.com/blue-monads/turnix/backend/modules/simplerat/wire"
)

type Response struct {
	Mid        uint64
	Pid        uint16
	Mtype      uint8
	DataBinary []byte
}

func (a *AgentNode) worker() {

	lhandlers := make(map[string]Handler)
	hLock.Lock()
	for k, v := range handlers {
		lhandlers[k] = v
	}
	hLock.Unlock()

	writeErr := func(msgId uint64, errStr string) {
		slog.Warn("Error handling packet", slog.String("error", errStr))
		a.writeLoopCh <- &Response{
			Mid:        uint64(msgId),
			Pid:        0,
			Mtype:      wire.PtypeResponseErr,
			DataBinary: []byte(errStr),
		}
	}

	for {

		packet := <-a.workersChan

		if packet == nil {
			return
		}

		if packet.Header.MType == "help" {

			commands := make([]string, 0, len(lhandlers))

			for k := range lhandlers {
				commands = append(commands, k)
			}

			out, err := json.Marshal(commands)
			if err != nil {
				writeErr((packet.Header.Mid), fmt.Sprintf("error marshalling help: %s", err))
				continue
			}

			a.writeLoopCh <- &Response{
				Mid:        packet.Header.Mid,
				Pid:        packet.Header.Pid,
				Mtype:      wire.PtypeResponseOk,
				DataBinary: out,
			}

			return
		}

		handler, ok := handlers[packet.Header.MType]
		if !ok {
			slog.Warn("Unknown packet type", slog.Any("packet", packet))
			writeErr((packet.Header.Mid), fmt.Sprintf("unknown packet type: %s", packet.Header.MType))
			continue
		}

		handler(&WHContext{
			Packet: packet,
			Node:   a,
		})

	}

}
