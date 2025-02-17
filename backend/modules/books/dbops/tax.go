package dbops

import (
	"github.com/blue-monads/turnix/backend/modules/books/models"
	"github.com/upper/db/v4"
)

// tax

func (b *DbOps) TaxAdd(pid, uid int64, data *models.Tax) (int64, error) {
	err := b.userHasScope(pid, uid, "write")
	if err != nil {
		return 0, err
	}

	table := b.taxTable(pid)

	r, err := table.Insert(data)
	if err != nil {
		return 0, err
	}

	return r.ID().(int64), nil
}

func (b *DbOps) TaxUpdate(pid, uid, id int64, data map[string]any) error {
	err := b.userHasScope(pid, uid, "write")
	if err != nil {
		return err
	}

	return b.taxTable(pid).Find(db.Cond{"id": id}).Update(data)
}

func (b *DbOps) TaxGet(pid, uid, id int64) (*models.Tax, error) {
	err := b.userHasScope(pid, uid, "read")
	if err != nil {
		return nil, err
	}

	data := &models.Tax{}
	table := b.taxTable(pid)

	err = table.Find(db.Cond{"id": id}).One(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (b *DbOps) TaxList(pid, uid int64) ([]models.Tax, error) {
	err := b.userHasScope(pid, uid, "read")
	if err != nil {
		return nil, err
	}

	datas := make([]models.Tax, 0)
	table := b.taxTable(pid)

	err = table.Find().All(&datas)
	if err != nil {
		return nil, err
	}

	return datas, nil
}

func (b *DbOps) TaxDelete(pid, uid, id int64) error {
	err := b.userHasScope(pid, uid, "write")
	if err != nil {
		return err
	}

	return b.taxTable(pid).Find(db.Cond{"id": id}).Delete()
}

// ProductTax

func (b *DbOps) ProductTaxAdd(pid, uid int64, data *models.ProductTax) (int64, error) {
	err := b.userHasScope(pid, uid, "write")
	if err != nil {
		return 0, err
	}

	table := b.productTaxTable(pid)

	r, err := table.Insert(data)
	if err != nil {
		return 0, err
	}

	return r.ID().(int64), nil
}

func (b *DbOps) ProductTaxUpdate(pid, uid, id int64, data map[string]any) error {
	err := b.userHasScope(pid, uid, "write")
	if err != nil {
		return err
	}

	return b.productTaxTable(pid).Find(db.Cond{"id": id}).Update(data)
}

func (b *DbOps) ProductTaxGet(pid, uid, id int64) (*models.ProductTax, error) {
	err := b.userHasScope(pid, uid, "read")
	if err != nil {
		return nil, err
	}

	data := &models.ProductTax{}
	table := b.productTaxTable(pid)

	err = table.Find(db.Cond{"id": id}).One(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (b *DbOps) ProductTaxList(pid, uid int64) ([]models.ProductTax, error) {
	err := b.userHasScope(pid, uid, "read")
	if err != nil {
		return nil, err
	}

	datas := make([]models.ProductTax, 0)
	table := b.productTaxTable(pid)

	err = table.Find().All(&datas)
	if err != nil {
		return nil, err
	}

	return datas, nil
}

func (b *DbOps) ProductTaxDelete(pid, uid, id int64) error {
	err := b.userHasScope(pid, uid, "write")
	if err != nil {
		return err
	}

	return b.productTaxTable(pid).Find(db.Cond{"id": id}).Delete()
}
