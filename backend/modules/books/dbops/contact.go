package dbops

import (
	"github.com/blue-monads/turnix/backend/modules/books/models"
	"github.com/upper/db/v4"
)

// contacts

func (b *DbOps) ContactAdd(pid, uid int64, data *models.Contact) (int64, error) {
	err := b.userHasScope(pid, uid, "write")
	if err != nil {
		return 0, err
	}

	table := b.contactTable(pid)

	r, err := table.Insert(data)
	if err != nil {
		return 0, err
	}

	return r.ID().(int64), nil
}

func (b *DbOps) ContactUpdate(pid, uid, id int64, data map[string]any) error {
	err := b.userHasScope(pid, uid, "write")
	if err != nil {
		return err
	}

	return b.contactTable(pid).Find(db.Cond{"id": id}).Update(data)
}

func (b *DbOps) ContactGet(pid, uid, id int64) (*models.Contact, error) {
	err := b.userHasScope(pid, uid, "read")
	if err != nil {
		return nil, err
	}

	data := &models.Contact{}
	table := b.contactTable(pid)

	err = table.Find(db.Cond{"id": id}).One(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (b *DbOps) ContactList(pid, uid, offset int64) ([]models.Contact, error) {
	err := b.userHasScope(pid, uid, "read")
	if err != nil {
		return nil, err
	}

	datas := make([]models.Contact, 0)
	table := b.contactTable(pid)

	err = table.Find(db.Cond{"is_deleted": false, "id >": offset}).Limit(100).All(&datas)
	if err != nil {
		return nil, err
	}

	return datas, nil
}

func (b *DbOps) ContactDelete(pid, uid, id int64) error {
	err := b.userHasScope(pid, uid, "write")
	if err != nil {
		return err
	}

	return b.contactTable(pid).Find(db.Cond{"id": id}).Update(db.Cond{
		"is_deleted": true,
	})
}
