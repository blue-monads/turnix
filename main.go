package main

import (
	"log"

	"github.com/bornjre/trunis/backend"
	"github.com/k0kubun/pp"
	_ "github.com/mattn/go-sqlite3"
)

func main() {

	// migration.Migrate()

	pp.Println("@start")

	err := backend.Run()
	if err != nil {
		log.Fatalln(err)
	}

}
