package dbops

import (
	"github.com/blue-monads/turnix/backend/modules/books/models"
	"github.com/upper/db/v4"
)

func (b *DbOps) CatagoryAdd(pid, uid int64, data *models.Catagory) (int64, error) {
	err := b.userHasScope(pid, uid, "write")
	if err != nil {
		return 0, err
	}

	table := b.catagoryTable(pid)

	r, err := table.Insert(data)
	if err != nil {
		return 0, err
	}

	return r.ID().(int64), nil
}

func (b *DbOps) CatagoryUpdate(pid, uid, id int64, data map[string]any) error {
	err := b.userHasScope(pid, uid, "write")
	if err != nil {
		return err
	}

	return b.catagoryTable(pid).Find(db.Cond{"id": id}).Update(data)
}

func (b *DbOps) CatagoryGet(pid, uid, id int64) (*models.Catagory, error) {
	err := b.userHasScope(pid, uid, "read")
	if err != nil {
		return nil, err
	}

	data := &models.Catagory{}
	table := b.catagoryTable(pid)

	err = table.Find(db.Cond{"id": id}).One(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (b *DbOps) CatagoryList(pid, uid int64) ([]models.Catagory, error) {
	err := b.userHasScope(pid, uid, "read")
	if err != nil {
		return nil, err
	}

	datas := make([]models.Catagory, 0)
	table := b.catagoryTable(pid)

	err = table.Find().All(&datas)
	if err != nil {
		return nil, err
	}

	return datas, nil
}

func (b *DbOps) CatagoryDelete(pid, uid, id int64) error {
	err := b.userHasScope(pid, uid, "write")
	if err != nil {
		return err
	}

	return b.catagoryTable(pid).Find(db.Cond{"id": id}).Update(db.Cond{
		"is_deleted": true,
	})
}
