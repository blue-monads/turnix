package ebrowser

import "github.com/bornjre/turnix/backend/distro/climux"

func init() {

	climux.Register(&climux.Action{
		Name: "ebrowser",
		Help: "Run embed browser with state folder",
		Func: func(ctx climux.Context) error {

			New(ctx).Run()

			return nil
		},
	})

	climux.DefaultCLI = "ebrowser"

}
