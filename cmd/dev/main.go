package main

import (
	_ "github.com/mattn/go-sqlite3"

	// modules

	"github.com/bornjre/turnix/backend/distro/devmode"

	_ "github.com/bornjre/turnix/backend/modules/books"

	// _ "github.com/bornjre/turnix/backend/modules/tracker"
	_ "github.com/bornjre/turnix/backend/modules/unloop"
)

func main() {

	devmode.Run()

}
