package dbops

import (
	"github.com/bornjre/turnix/backend/modules/books/models"
	"github.com/upper/db/v4"
)

func (b *DbOps) EstimateAdd(pid, uid int64, data *models.Estimate) (int64, error) {
	err := b.userHasScope(pid, uid, "write")
	if err != nil {
		return 0, err
	}

	table := b.estimateTable(pid)

	r, err := table.Insert(data)
	if err != nil {
		return 0, err
	}

	return r.ID().(int64), nil
}

func (b *DbOps) EstimateUpdate(pid, uid, id int64, data map[string]any) error {
	err := b.userHasScope(pid, uid, "write")
	if err != nil {
		return err
	}

	return b.estimateTable(pid).Find(db.Cond{"id": id}).Update(data)
}

func (b *DbOps) EstimateGet(pid, uid, id int64) (*models.Estimate, error) {
	err := b.userHasScope(pid, uid, "read")
	if err != nil {
		return nil, err
	}

	data := &models.Estimate{}
	table := b.estimateTable(pid)

	err = table.Find(db.Cond{"id": id}).One(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (b *DbOps) EstimateList(pid, uid, offset int64) ([]models.Estimate, error) {
	err := b.userHasScope(pid, uid, "read")
	if err != nil {
		return nil, err
	}

	datas := make([]models.Estimate, 0)
	table := b.estimateTable(pid)

	err = table.Find(db.Cond{"id >": offset}).Limit(100).All(&datas)
	if err != nil {
		return nil, err
	}

	return datas, nil
}

func (b *DbOps) EstimateDelete(pid, uid, id int64) error {
	err := b.userHasScope(pid, uid, "write")
	if err != nil {
		return err
	}

	return b.estimateTable(pid).Find(db.Cond{"id": id}).Delete()
}
