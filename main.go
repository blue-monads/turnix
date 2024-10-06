package main

import (
	"github.com/bornjre/turnix/cmd/cli"

	// "github.com/bornjre/turnix/backend/distro/climux"
	_ "github.com/bornjre/turnix/backend/distro/climux/ebrowser"
	_ "github.com/bornjre/turnix/backend/distro/climux/node"
	_ "github.com/mattn/go-sqlite3"

	_ "github.com/bornjre/turnix/backend/modules/books"
	_ "github.com/bornjre/turnix/backend/modules/unloop"
)

func main() {

	// climux.DefaultCLI = "ebrowser" // uncomment make run ebrowser run by default

	cli.RunCLI()

}
