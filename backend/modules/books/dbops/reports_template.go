package dbops

import (
	"github.com/bornjre/turnix/backend/modules/books/models"
	"github.com/upper/db/v4"
)

func (b *DbOps) ReportTemplateList(pid, uid, offset int64) ([]models.ReportTemplate, error) {

	table := b.reportTemplateTable(pid)

	datas := make([]models.ReportTemplate, 0)

	err := table.Find(db.Cond{"id >": offset}).
		Select("id", "name", "viewer_editable", "created_by", "updated_by").
		Limit(100).All(&datas)
	if err != nil {
		return nil, err
	}

	return datas, nil
}

func (b *DbOps) ReportTemplateAdd(pid, uid int64, data *models.ReportTemplate) (int64, error) {

	table := b.reportTemplateTable(pid)

	data.CreatedBy = uid
	data.UpdatedBy = uid

	r, err := table.Insert(data)
	if err != nil {
		return 0, err
	}

	return r.ID().(int64), nil
}

func (b *DbOps) ReportTemplateGet(pid, uid, id int64) (*models.ReportTemplate, error) {
	table := b.reportTemplateTable(pid)
	data := &models.ReportTemplate{}
	err := table.Find(db.Cond{"id": id}).One(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (b *DbOps) ReportTemplateUpdate(pid, uid, id int64, data map[string]any) error {
	table := b.reportTemplateTable(pid)
	return table.Find(db.Cond{"id": id}).Update(data)
}

func (b *DbOps) ReportTemplateDelete(pid, uid, id int64) error {
	table := b.reportTemplateTable(pid)

	return table.Find(db.Cond{"id": id}).Delete()

}
