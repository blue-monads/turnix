package main

import (
	_ "github.com/mattn/go-sqlite3"

	// modules

	"github.com/blue-monads/turnix/backend/distro/devmode"

	_ "github.com/blue-monads/turnix/backend/modules/books"

	// _ "github.com/blue-monads/turnix/backend/modules/tracker"
	_ "github.com/blue-monads/turnix/backend/modules/simplerat"
	_ "github.com/blue-monads/turnix/backend/modules/unloop"
)

func main() {

	devmode.Run()

}
