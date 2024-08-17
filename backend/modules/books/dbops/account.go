package dbops

import (
	"github.com/bornjre/turnix/backend/modules/books/models"
	"github.com/upper/db/v4"
)

func (b *DbOps) AddAccount(pid, uid int64, data *models.Account) (int64, error) {

	err := b.userHasScope(pid, uid, "write")
	if err != nil {
		return 0, err
	}

	table := b.accountsTable(pid)

	r, err := table.Insert(data)
	if err != nil {
		return 0, err
	}

	return r.ID().(int64), nil
}

func (b *DbOps) UpdateAccount(pid, uid, id int64, data map[string]any) error {
	err := b.userHasScope(pid, uid, "write")
	if err != nil {
		return err

	}

	return b.accountsTable(pid).Find(db.Cond{"id": id}).Update(data)
}

func (b *DbOps) GetAccount(pid, uid, aid int64) (*models.Account, error) {

	err := b.userHasScope(pid, uid, "read")
	if err != nil {
		return nil, err
	}

	data := &models.Account{}
	table := b.accountsTable(pid)

	err = table.Find(db.Cond{"id": aid}).One(data)
	if err != nil {
		return nil, err
	}

	return data, nil

}

func (b *DbOps) ListAccount(pid, uid int64) ([]models.Account, error) {

	err := b.userHasScope(pid, uid, "read")
	if err != nil {
		return nil, err
	}

	datas := make([]models.Account, 0)
	table := b.accountsTable(pid)

	err = table.Find(db.Cond{"is_deleted": false}).All(&datas)
	if err != nil {
		return nil, err
	}

	return datas, nil

}

func (b *DbOps) DeleteAccount(pid, uid, aid int64) error {

	err := b.userHasScope(pid, uid, "write")
	if err != nil {
		return err
	}

	table := b.accountsTable(pid)
	err = table.Find(db.Cond{"id": aid}).Update(db.Cond{
		"is_deleted": true,
	})

	return err
}
