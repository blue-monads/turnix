package dbops

import (
	"context"
	"fmt"

	"github.com/bornjre/turnix/backend/modules/books/models"
	"github.com/upper/db/v4"
)

// products

func (b *DbOps) ProductAdd(pid, uid int64, data *models.Product) (int64, error) {
	err := b.userHasScope(pid, uid, "write")
	if err != nil {
		return 0, err
	}

	table := b.productTable(pid)

	r, err := table.Insert(data)
	if err != nil {
		return 0, err
	}

	return r.ID().(int64), nil
}

func (b *DbOps) ProductUpdate(pid, uid, id int64, data map[string]any) error {
	err := b.userHasScope(pid, uid, "write")
	if err != nil {
		return err
	}

	return b.productTable(pid).Find(db.Cond{"id": id}).Update(data)
}

func (b *DbOps) ProductGet(pid, uid, id int64) (*models.Product, error) {
	err := b.userHasScope(pid, uid, "read")
	if err != nil {
		return nil, err
	}

	data := &models.Product{}
	table := b.productTable(pid)

	err = table.Find(db.Cond{"id": id}).One(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (b *DbOps) ProductList(pid, uid int64) ([]models.Product, error) {
	err := b.userHasScope(pid, uid, "read")
	if err != nil {
		return nil, err
	}

	datas := make([]models.Product, 0)
	table := b.productTable(pid)

	err = table.Find().All(&datas)
	if err != nil {
		return nil, err
	}

	return datas, nil
}

func (b *DbOps) ProductDelete(pid, uid, id int64) error {
	err := b.userHasScope(pid, uid, "write")
	if err != nil {
		return err
	}

	return b.productTable(pid).Find(db.Cond{"id": id}).Delete()
}

// product stock in

type StockInData struct {
	Info     string            `json:"info"`
	Amount   float64           `json:"amount"`
	VendorID int64             `json:"vendor_id"`
	Lines    []StockInLineData `json:"lines"`
}

type StockInLineData struct {
	Info      string  `json:"info"`
	ProductID int64   `json:"product_id"`
	Qty       float64 `json:"qty"`
	Amount    float64 `json:"amount"`
}

func (b *DbOps) ProductStockInAdd(pid, uid int64, data *StockInData) (pstockInid int64, err error) {
	err = b.userHasScope(pid, uid, "write")
	if err != nil {
		return 0, err
	}

	psinTable := b.productStockInTable(pid)
	psinLinTable := b.productStockInLineTable(pid)

	r, err := psinTable.Insert(data)
	if err != nil {
		return 0, err
	}

	pstockInid = r.ID().(int64)
	doneLines := make([]int64, 0, len(data.Lines))

	defer func() {
		if err == nil {
			return
		}

		psinTable.Find(db.Cond{"id": pstockInid}).Delete()
		psinLinTable.Find(db.Cond{"id IN": doneLines}).Delete()

	}()

	for _, inline := range data.Lines {

		line := &models.ProductStockInLine{
			Info:             inline.Info,
			ProductID:        inline.ProductID,
			Qty:              inline.Qty,
			Amount:           inline.Amount,
			ProductStockInID: pstockInid,
			CreatedBy:        uid,
			UpdatedBy:        uid,
		}

		r, err := b.productStockInLineTable(pid).Insert(line)
		if err != nil {
			return 0, err
		}

		doneLines = append(doneLines, r.ID().(int64))
	}

	stmt := fmt.Sprintf(`UPDATE %s SET stock_count = stock_count + ? WHERE id = ?`, productTableName(pid))

	err = b.db.GetSession().TxContext(context.Background(), func(sess db.Session) error {

		for _, inline := range data.Lines {
			_, err := sess.SQL().Exec(stmt, inline.ProductID, inline.Qty)
			if err != nil {
				return err
			}
		}

		return nil
	}, nil)

	if err != nil {
		return 0, err
	}

	return pstockInid, nil
}

func (b *DbOps) ProductStockInGet(pid, uid, id int64) (*models.ProductStockIn, error) {
	err := b.userHasScope(pid, uid, "read")
	if err != nil {
		return nil, err
	}

	data := &models.ProductStockIn{}
	table := b.productStockInTable(pid)

	err = table.Find(db.Cond{"id": id}).One(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (b *DbOps) ProductStockInList(pid, uid int64) ([]models.ProductStockIn, error) {
	err := b.userHasScope(pid, uid, "read")
	if err != nil {
		return nil, err
	}

	datas := make([]models.ProductStockIn, 0)
	table := b.productStockInTable(pid)

	err = table.Find().All(&datas)
	if err != nil {
		return nil, err
	}

	return datas, nil
}

func (b *DbOps) ProductStockInDelete(pid, uid, id int64) error {
	err := b.userHasScope(pid, uid, "write")
	if err != nil {
		return err
	}

	return b.productStockInTable(pid).Find(db.Cond{"id": id}).Delete()
}
