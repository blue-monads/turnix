package backend

import (
	"github.com/bornjre/trunis/backend/app"
	"github.com/bornjre/trunis/backend/database"
	"github.com/bornjre/trunis/backend/token"
)

func Run() error {

	db, err := database.NewDB()
	if err != nil {
		return err
	}

	defer db.Close()

	signer := token.New([]byte("A_long_HARD_Token"))

	as := app.New(db, signer)

	return as.Run()

}
