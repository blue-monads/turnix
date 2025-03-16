package main

import (
	"log"

	"github.com/blue-monads/turnix/backend/extras/buddy"
	"github.com/k0kubun/pp"
	"github.com/upper/db/v4/adapter/sqlite"
)

var settings = sqlite.ConnectionURL{
	Database: "./play.db",
}

func main() {
	pp.Println("@play")

	sess, err := sqlite.Open(settings)
	if err != nil {
		log.Fatalf("db.Open(): %q\n", err)
	}

	cpature := buddy.New(sess)
	err = cpature.CaptureHooks()
	if err != nil {
		pp.Println(err)
		return
	}

}
