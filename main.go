package main

import (
	"log"

	"github.com/bornjre/trunis/backend"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	// migration.Migrate()

	err := backend.Run()
	if err != nil {
		log.Fatalln(err)
	}

}
