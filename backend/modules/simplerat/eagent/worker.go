package main

import (
	"log/slog"
)

var (
	handlers = map[string]func(*Packet) (interface{}, error){
		"run-cmd": HandleRunCommand,
		"ping":    HandlePing,
	}
)

func (a *AgentService) worker() {

	for {

		packet := <-a.workersChan

		if packet == nil {
			return
		}

		handler, ok := handlers[packet.Data.MType]
		if !ok {
			slog.Warn("Unknown packet type", slog.AnyValue(packet))

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

		resp, err := handler(packet)
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
	}

}

func HandleRunCommand(packet *Packet) (interface{}, error) {

	return "nice", nil
}

func HandlePing(packet *Packet) (interface{}, error) {

	return "pong", nil
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
