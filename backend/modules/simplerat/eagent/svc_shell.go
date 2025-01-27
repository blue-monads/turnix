package main

import (
	"context"
	"log/slog"
	"os"
	"os/exec"
	"sync"

	"github.com/blue-monads/turnix/backend/utils/kosher"
	"github.com/coder/websocket"
	"github.com/creack/pty"
	"github.com/k0kubun/pp"
)

type ServiceShell struct {
	node *AgentNode
	ws   *websocket.Conn
	cmd  *exec.Cmd
	file *os.File

	isClosed  bool
	closeLock sync.Mutex
}

func (s *ServiceShell) Run() error {

	c := exec.Command("bash")
	tty, err := pty.Start(c)
	if err != nil {
		slog.Error("Error starting shell: ", slog.String("error", err.Error()))
		return err
	}

	s.cmd = c
	s.file = tty

	go s.ReadLoop()
	go s.WriteLoop()

	return nil
}

func (s *ServiceShell) close() {
	s.closeLock.Lock()
	defer s.closeLock.Unlock()

	if s.isClosed {
		return
	}

	s.isClosed = true

	s.file.Close()
	s.ws.CloseNow()
	s.cmd.Process.Kill()
}

func (s *ServiceShell) ReadLoop() {
	ctx := context.Background()

	defer func() {
		s.close()
	}()

	errCounter := 0

	for {

		if errCounter > 5 {
			return
		}

		_, out, err := s.ws.Read(ctx)
		if err != nil {
			slog.Error("Error reading message: ", slog.String("error", err.Error()))
			errCounter++
			continue
		}

		errCounter = 0

		pp.Println("msg", kosher.Str(out))

		s.file.Write(out)
	}

}

func (s *ServiceShell) WriteLoop() {
	ctx := context.Background()

	buf := make([]byte, 1024)

	defer func() {
		s.close()
	}()

	errCounter := 0

	for {

		if errCounter > 7 {
			return
		}

		n, err := s.file.Read(buf)
		if err != nil {
			slog.Error("Error reading from shell: ", slog.String("error", err.Error()))
			errCounter++
			continue
		}

		msg := buf[:n]

		pp.Println("msg", kosher.Str(msg))

		err = s.ws.Write(ctx, websocket.MessageText, msg)
		if err != nil {
			slog.Error("Error writing message: ", slog.String("error", err.Error()))
			errCounter++
			continue
		}

		errCounter = 0

	}
}
