package distro

import "github.com/bornjre/turnix/backend/xtypes/models"

func (d *DistroApp) RunNormalSeed() error {
	db := d.App.GetDatabase()

	userId, err := db.AddUser(&models.User{
		Name:            "dev",
		Utype:           "real",
		Email:           "dev@example.com",
		Bio:             "I am a dev.",
		Password:        "dev123",
		IsEmailVerified: true,
	})
	if err != nil {
		return err
	}

	_, err = db.AddUser(&models.User{
		Name:            "zesus",
		Utype:           "device",
		Email:           "zesus@example.com",
		Bio:             "I am a zesus.",
		Password:        "dev123",
		IsEmailVerified: true,
		OwnerUserId:     userId,
	})
	if err != nil {
		return err
	}

	_, err = db.AddProject(&models.Project{
		Name:         "Automation Loop",
		Info:         "Loop for my automation projects",
		Ptype:        "unloop",
		OwnerID:      userId,
		IsInitilized: true,
	})

	if err != nil {
		return err
	}

	_, err = db.AddProject(&models.Project{
		Name:         "My books",
		Info:         "My accounts tracking",
		Ptype:        "books",
		OwnerID:      userId,
		IsInitilized: true,
	})

	return err

}

func (d *DistroApp) RunTestSeed() error {

	db := d.App.GetDatabase()

	userId, err := db.AddUser(&models.User{
		Name:            "dev",
		Utype:           "real",
		Email:           "dev@example.com",
		Bio:             "I am a dev.",
		Password:        "dev123",
		IsEmailVerified: true,
	})

	if err != nil {
		return err
	}

	db.AddProject(&models.Project{
		Name:         "test1",
		Info:         "this is a test project",
		Ptype:        "test",
		OwnerID:      userId,
		IsInitilized: true,
		IsPublic:     true,
	})

	return nil

}
