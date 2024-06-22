package main

import (
	_ "github.com/mattn/go-sqlite3"

	// modules

	"github.com/bornjre/turnix/backend/distro"
	_ "github.com/bornjre/turnix/backend/modules/books"
	_ "github.com/bornjre/turnix/backend/modules/tracker"
	_ "github.com/bornjre/turnix/backend/modules/unloop"
)

func main() {

	app, err := distro.NewApp()
	if err != nil {
		panic(err)
	}

	err = app.RunSeed()
	if err != nil {
		panic(err)
	}

}
