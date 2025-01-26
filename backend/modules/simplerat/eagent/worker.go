package main

import (
	"fmt"
	"log/slog"
)

func (a *AgentService) worker() {

	lhandlers := make(map[string]Handler)
	hLock.Lock()
	for k, v := range handlers {
		lhandlers[k] = v
	}
	hLock.Unlock()

	writeErr := func(msgId int64, errStr string) {
		slog.Warn("Error handling packet", slog.String("error", errStr))
		a.writeLoopCh <- &Packet{
			MessageId: msgId,
			Data: Message{
				MType: "response",
				Data:  errStr,
			},
		}
	}

	for {

		packet := <-a.workersChan

		if packet == nil {
			return
		}

		if packet.Data.MType == "help" {

			commands := make([]string, 0, len(lhandlers))

			for k := range lhandlers {
				commands = append(commands, k)
			}

			a.writeLoopCh <- &Packet{
				MessageId: packet.MessageId,
				Data: Message{
					MType: "response",
					Data:  commands,
				},
			}

			return
		}

		handler, ok := handlers[packet.Data.MType]
		if !ok {
			slog.Warn("Unknown packet type", slog.Any("packet", packet))
			writeErr(packet.MessageId, fmt.Sprintf("unknown packet type: %s", packet.Data.MType))
			continue
		}

		resp, err := handler(&WHContext{
			Packet:  packet,
			Service: a,
		})
		if err != nil {
			slog.Warn("Error handling packet", slog.String("error", err.Error()))
			continue
		}

		slog.Info("Sending response", slog.Any("response", resp))

		a.writeLoopCh <- &Packet{
			MessageId: packet.MessageId,
			Data: Message{
				MType: "response",
				Data:  resp,
			},
		}
	}

}
