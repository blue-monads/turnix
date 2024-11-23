package distro

import (
	"github.com/blue-monads/turnix/backend/services/database"
	"github.com/blue-monads/turnix/backend/xtypes/models"
)

func (d *DistroApp) RunNormalSeed() error {
	dz := d.App.GetDatabase()

	db := dz.(*database.DB)

	//	innerApp := d.App.(*app.App)

	userId, err := db.AddUser(&models.User{
		Name:       "dev",
		Utype:      "real",
		Email:      "dev@example.com",
		Bio:        "I am a dev.",
		Password:   "dev123",
		IsVerified: true,
	})
	if err != nil {
		return err
	}

	_, err = db.AddUser(&models.User{
		Name:        "zesus",
		Utype:       "device",
		Email:       "zesus@example.com",
		Bio:         "I am a zesus.",
		Password:    "dev123",
		IsVerified:  true,
		OwnerUserId: userId,
	})
	if err != nil {
		return err
	}

	// _, err = innerApp.Addproject(&models.Project{
	// 	Name:         "Automation Loop",
	// 	Info:         "Loop for my automation projects",
	// 	Ptype:        "unloop",
	// 	OwnerID:      userId,
	// 	IsInitilized: true,
	// })

	// if err != nil {
	// 	return err
	// }

	// _, err = innerApp.Addproject(&models.Project{
	// 	Name:         "My books",
	// 	Info:         "My accounts tracking",
	// 	Ptype:        "books",
	// 	OwnerID:      userId,
	// 	IsInitilized: true,
	// })

	return err

}
