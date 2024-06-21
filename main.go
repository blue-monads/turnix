package main

import (
	"log"

	"github.com/bornjre/turnix/backend"

	"github.com/k0kubun/pp"
	_ "github.com/mattn/go-sqlite3"

	// modules

	_ "github.com/bornjre/turnix/backend/modules/books"
	_ "github.com/bornjre/turnix/backend/modules/tracker"
	_ "github.com/bornjre/turnix/backend/modules/unloop"
)

func main() {

	// migration.Migrate()

	pp.Println("@start")

	err := backend.Run()
	if err != nil {
		log.Fatalln(err)
	}

}
