package dbops

import (
	"github.com/blue-monads/turnix/backend/modules/books/models"
	"github.com/upper/db/v4"
)

func (b *DbOps) SavedTemplateList(pid, uid, offset int64) ([]models.SavedReport, error) {
	err := b.userHasScope(pid, uid, "read")
	if err != nil {
		return nil, err
	}

	datas := make([]models.SavedReport, 0)
	table := b.savedReportTable(pid)

	err = table.Find(db.Cond{"id >": offset}).Limit(100).All(&datas)
	if err != nil {
		return nil, err
	}

	return datas, nil
}

func (b *DbOps) SavedTemplateGet(pid, uid, id int64) (*models.SavedReport, error) {
	err := b.userHasScope(pid, uid, "read")
	if err != nil {
		return nil, err
	}

	data := &models.SavedReport{}
	table := b.savedReportTable(pid)

	err = table.Find(db.Cond{"id": id}).One(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (b *DbOps) SavedTemplateAdd(pid, uid int64, data *models.SavedReport) (int64, error) {
	err := b.userHasScope(pid, uid, "write")
	if err != nil {
		return 0, err
	}

	data.CreatedBy = uid
	data.UpdatedBy = uid

	table := b.savedReportTable(pid)

	r, err := table.Insert(data)
	if err != nil {
		return 0, err
	}

	return r.ID().(int64), nil
}

func (b *DbOps) SavedTemplateUpdate(pid, uid, id int64, data map[string]any) error {
	err := b.userHasScope(pid, uid, "write")
	if err != nil {
		return err
	}

	return b.savedReportTable(pid).Find(db.Cond{"id": id}).Update(data)
}

func (b *DbOps) SavedTemplateDelete(pid, uid, id int64) error {
	err := b.userHasScope(pid, uid, "write")
	if err != nil {
		return err
	}

	return b.savedReportTable(pid).Find(db.Cond{"id": id}).Delete()

}
