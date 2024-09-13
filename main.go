package main

import (
	"github.com/bornjre/turnix/cmd/cli"

	_ "github.com/bornjre/turnix/backend/distro/climux/ebrowser"
	_ "github.com/bornjre/turnix/backend/distro/climux/node"
	_ "github.com/mattn/go-sqlite3"

	_ "github.com/bornjre/turnix/backend/modules/books"
	_ "github.com/bornjre/turnix/backend/modules/unloop"
)

func main() {

	cli.RunCLI()

}
