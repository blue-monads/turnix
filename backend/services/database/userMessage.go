package database

import (
	"github.com/blue-monads/turnix/backend/xtypes/models"
	"github.com/upper/db/v4"
)

func (d *DB) AddUserMessage(msg *models.UserMessage) (int64, error) {
	rid, err := d.userMessagesTable().Insert(msg)
	if err != nil {
		return 0, err
	}

	id := rid.ID().(int64)

	return id, nil
}

func (d *DB) ListUserMessages(uid, count, cursor int64) ([]models.UserMessage, error) {
	messages := make([]models.UserMessage, 0)
	cond := db.Cond{
		"to_user": uid,
	}

	if cursor != 0 {
		cond["id >"] = cursor
	}

	result := d.userMessagesTable().Find(cond)

	var err error
	if count != 0 {
		err = result.Limit(int(count)).OrderBy("id").All(&messages)
	} else {
		err = result.OrderBy("id").All(&messages)
	}

	if err != nil {
		return nil, err
	}

	return messages, nil
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
