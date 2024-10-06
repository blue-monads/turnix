package ebrowser

import (
	"time"

	"github.com/bornjre/turnix/backend/distro/climux"
)

func init() {

	climux.Register(&climux.Action{
		Name: "ebrowser",
		Help: "Run embed browser with state folder",
		Func: func(ctx climux.Context) error {

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

	climux.DefaultCLI = "ebrowser"

}
