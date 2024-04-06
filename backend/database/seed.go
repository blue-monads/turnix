package database

func (db *DB) RunSeed() error {

	userId, err := db.AddUser(&User{
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

	_, err = db.AddUser(&User{
		Name:            "zesus",
		Utype:           "device",
		Email:           "zesus@example.com",
		Bio:             "I am a zesus.",
		Password:        "dev123",
		IsEmailVerified: true,
		OwnedBy:         userId,
	})
	if err != nil {
		return err
	}

	_, err = db.AddProject(&Project{
		Name:         "zesus",
		Info:         "This is test proj",
		Ptype:        "onloop",
		OwnerID:      userId,
		IsInitilized: true,
	})

	return err

}
