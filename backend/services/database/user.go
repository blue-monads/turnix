package database

import (
	"github.com/bornjre/turnix/backend/xtypes/models"
	"github.com/upper/db/v4"
)

func (d *DB) AddUser(data *models.User) (int64, error) {
	r, err := d.userTable().Insert(data)
	if err != nil {
		return 0, err
	}

	return r.ID().(int64), nil
}

func (d *DB) UpdateUser(id int64, data map[string]any) error {
	return d.userTable().Find(db.Cond{"id": id}).Update(data)
}

func (d *DB) GetUser(id int64) (*models.User, error) {

	data := &models.User{}

	err := d.userTable().Find(db.Cond{"id": id}).One(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (d *DB) GetUserByEmail(email string) (*models.User, error) {

	data := &models.User{}

	err := d.userTable().Find(db.Cond{"email": email}).One(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (d *DB) ListUser() ([]models.User, error) {

	users := make([]models.User, 0)

	err := d.userTable().Find().All(&users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (d *DB) ListUserByOwner(owner int64) ([]models.User, error) {

	users := make([]models.User, 0)

	err := d.userTable().Find(db.Cond{"owned_by": owner}).All(&users)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (d *DB) DeleteUser(id int64) error {
	return d.userTable().Find(db.Cond{"id": id}).Delete()
}

// private

func (d *DB) userTable() db.Collection {
	return d.Table("Users")
}
