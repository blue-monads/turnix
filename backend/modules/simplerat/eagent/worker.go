package main

import (
	"log/slog"
)

func (a *AgentService) worker() {

	lhandlers := make(map[string]Handler)
	hLock.Lock()
	for k, v := range handlers {
		lhandlers[k] = v
	}
	hLock.Unlock()

	for {

		packet := <-a.workersChan

		if packet == nil {
			return
		}

		handler, ok := handlers[packet.Data.MType]
		if !ok {
			slog.Warn("Unknown packet type", slog.Any("packet", packet))

			a.writeLoopCh <- &Packet{
				Data: Message{
					MType: "error",
					Data: map[string]any{
						"error": "Unknown packet type",
					},
				},
				MessageId: packet.MessageId,
			}

			continue
		}

		resp, err := handler(&WHContext{
			Packet:  packet,
			Service: a,
		})
		if err != nil {
			slog.Warn("Error handling packet", slog.String("error", err.Error()))
			a.writeLoopCh <- &Packet{
				MessageId: packet.MessageId,
				Data: Message{
					MType: "response",
					Data:  resp,
				},
			}

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

func HandleRunCommand(ctx *WHContext) (interface{}, error) {

	return map[string]any{
		"message": "command executed",
	}, nil
}

func HandlePing(ctx *WHContext) (interface{}, error) {

	return map[string]any{
		"message": "pong",
	}, nil
}

/*

	FSListDir(path string) ([]string, error)
	FSReadFile(path string) ([]byte, error)
	FSWriteFile(path string, data []byte) error
	FSRemove(path string) error
	FSMkdir(path string) error
	FSRename(old string, new string) error

	SystemExec(command string) ([]byte, error)
	SystemStartShell(cmd string, submitFunc func([]byte) error) (chan []byte, error)

*/
