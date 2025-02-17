package cplane

import (
	"time"

	"github.com/blue-monads/turnix/backend/distro/climux"
)

func init() {

	climux.Register(&climux.Action{
		Name: "cplane",
		Help: "Run app control plane",
		HandleCLI: func(ctx climux.Context) error {

			ebApp := New(ctx)

			defer ebApp.Close()

			go ebApp.runPreHttpServer()

			time.Sleep(time.Second * 1)

			if ebApp.port == 0 {
				time.Sleep(time.Second * 3)
				if ebApp.port == 0 {
					panic("pre http server not running")
				}
			}

			ebApp.Run()

			return nil
		},
	})

}
