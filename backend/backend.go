package backend

import (
	"github.com/bornjre/trunis/backend/app"
	"github.com/bornjre/trunis/backend/services/database"
	"github.com/bornjre/trunis/backend/services/signer"
	"github.com/bornjre/trunis/backend/xtypes/xproject"
)

func Run() error {

	db, err := database.NewDB()
	if err != nil {
		return err
	}

	defer db.Close()

	err = db.Init()
	if err != nil {
		return err
	}

	signer := signer.New([]byte("A_long_HARD_Token"))

	as := app.New(app.Options{
		DB:     db,
		Signer: signer,
		ProjectTypes: []*xproject.TypeDefination{
			{
				Name: "On Loop",
				Slug: "onloop",
				Info: "Perform action with you on loop",
				Icon: "arrow-path",
				NewFormSchemaFields: []xproject.PTypeField{
					{
						Name:       "Init examples",
						KeyName:    "init_examples",
						Ftype:      "BOOL",
						IsRequired: false,
					},
				},
			},
			{
				Name:                "Trackers",
				Slug:                "tracker",
				Info:                "Track everything",
				Icon:                "map-pin",
				NewFormSchemaFields: []xproject.PTypeField{},
			},
			{
				Name:                "Books",
				Slug:                "books",
				Info:                "Track everything",
				Icon:                "book-open",
				NewFormSchemaFields: []xproject.PTypeField{},
			},
		},
	})

	return as.Run()

}
