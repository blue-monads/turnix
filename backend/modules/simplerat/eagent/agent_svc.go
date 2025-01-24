package main

import (
	"log/slog"
	"time"

	"github.com/k0kubun/pp"
)

type AgentService struct {
	configHome string
	configFile string

	baseURL          string
	registerToken    string
	authSessionToken string
	authRefreshToken string
}

// http://localhost:7703/z/project/simplerat/5/device/finish-setup?token=4mzN3iGtuYdJJKM3O78OyZPjFTorUuNWox3uuK73ikXHWEFAIoFlfkM9JwPuiZyb9n0IOPB4CQT23e401RYW9tREEVxpLVap3IkTDyMo

func (a *AgentService) Run() error {

	slog.Info("Agent service started")

	pp.Println(a)

	for {
		time.Sleep(20 * time.Second)
		slog.Info("Agent service running")
	}

}
