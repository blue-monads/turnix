package main

import (
	"github.com/blue-monads/turnix/cmd/cli"

	// "github.com/blue-monads/turnix/backend/distro/climux"

	_ "github.com/blue-monads/turnix/backend/distro/climux/cplane"
	_ "github.com/blue-monads/turnix/backend/distro/climux/node"
	_ "github.com/mattn/go-sqlite3"

	_ "github.com/blue-monads/turnix/backend/modules/books"
	_ "github.com/blue-monads/turnix/backend/modules/unloop"
)

func main() {

	cli.RunCLI()

}
