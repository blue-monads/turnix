package main

import (
	"github.com/bornjre/turnix/backend/distro"
	"github.com/bornjre/turnix/backend/registry"
	"github.com/bornjre/turnix/backend/xtypes/xproject"

	_ "github.com/mattn/go-sqlite3"

	// modules

	_ "github.com/bornjre/turnix/backend/modules/books"
	_ "github.com/bornjre/turnix/backend/modules/tracker"
	_ "github.com/bornjre/turnix/backend/modules/unloop"
)

func init() {
	registry.Register(&xproject.Defination{
		Name:                "test",
		Slug:                "test",
		Info:                "this is test type project",
		NewFormSchemaFields: []xproject.PTypeField{},
		Perminssions:        []string{},
		EventTypes:          []string{"sky_dropped"},
		Builder:             nil,
	})

}

func main() {

	app, err := distro.NewApp()
	handle(err)

	mig, err := app.NeedsMigrate()
	handle(err)

	if mig {
		err = app.RunTestSeed()
		handle(err)
	}

	err = app.Run()
	handle(err)

}

func handle(err error) {
	if err != nil {
		panic(err)
	}
}
