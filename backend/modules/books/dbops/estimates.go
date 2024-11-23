package dbops

import (
	"github.com/bornjre/turnix/backend/modules/books/models"
	"github.com/upper/db/v4"
)

type EstimateData struct {
	Estimate models.Estimate       `json:"estimate"`
	Lines    []models.EstimateLine `json:"lines"`
}

func (b *DbOps) EstimateAdd(pid, uid int64, data *EstimateData) (int64, error) {
	err := b.userHasScope(pid, uid, "write")
	if err != nil {
		return 0, err
	}

	table := b.estimateTable(pid)

	r, err := table.Insert(data.Estimate)
	if err != nil {
		return 0, err
	}

	estimateId := r.ID().(int64)

	lineTable := b.estimateLineTable(pid)

	for _, line := range data.Lines {
		line.EstimateId = estimateId

		_, err := lineTable.Insert(line)
		if err != nil {
			return 0, err
		}

	}

	return estimateId, nil
}

func (b *DbOps) EstimateUpdate(pid, uid, id int64, data map[string]any) error {
	err := b.userHasScope(pid, uid, "write")
	if err != nil {
		return err
	}

	return b.estimateTable(pid).Find(db.Cond{"id": id}).Update(data)
}

func (b *DbOps) EstimateGet(pid, uid, id int64) (*EstimateData, error) {
	err := b.userHasScope(pid, uid, "read")
	if err != nil {
		return nil, err
	}

	var estimateData EstimateData

	table := b.estimateTable(pid)

	err = table.Find(db.Cond{"id": id}).One(&estimateData.Estimate)
	if err != nil {
		return nil, err
	}

	lineTable := b.estimateLineTable(pid)

	err = lineTable.Find(db.Cond{"estimate_id": id}).All(&estimateData.Lines)
	if err != nil {
		return nil, err
	}

	return &estimateData, nil
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
