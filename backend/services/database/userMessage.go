package database

import (
	"github.com/bornjre/turnix/backend/xtypes/models"
	"github.com/upper/db/v4"
)

func (d *DB) AddUserMessage(msg *models.User) (int64, error) {
	rid, err := d.userMessagesTable().Insert(msg)
	if err != nil {
		return 0, err
	}

	id := rid.ID().(int64)

	return id, nil
}

func (d *DB) UserMessageSetRead(user string, id int64) error {
	return d.userMessagesTable().Find(
		db.Cond{
			"toUser": user,
			"id":     id},
	).Update(db.Cond{
		"isRead": true,
	})
}

func (d *DB) RemoveUserMessage(userId string, id int64) error {
	return d.userMessagesTable().Find(db.Cond{"toUser": userId, "id": id}).Delete()
}

func (d *DB) ListUserMessages(uid, count, cursor int64) ([]models.UserMessage, error) {
	messages := make([]models.UserMessage, 0)
	cond := db.Cond{
		"toUser": uid,
	}

	if cursor != 0 {
		cond["id >"] = cursor
	}

	result := d.userMessagesTable().Find(cond)

	var err error
	if count != 0 {
		err = result.Limit(int(count)).All(&messages)
	} else {
		err = result.All(&messages)
	}

	if err != nil {
		return nil, err
	}

	return messages, nil
}

func (d *DB) ReadUserMessages(userId string, id []int64) error {
	return d.userMessagesTable().Find(db.Cond{

		"toUser": userId,
		"id IN":  id,
	}).Update(db.Cond{
		"read": true,
	})
}

func (d *DB) DeleteUserMessages(userId string, id []int64) error {
	return d.userMessagesTable().Find(db.Cond{
		"toUser": userId,
		"id IN":  id,
	}).Delete()
}

func (d *DB) userMessagesTable() db.Collection {
	return d.Table("UserMessages")
}
