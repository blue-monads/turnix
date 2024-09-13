package dbops

import (
	"context"
	"fmt"
	"time"

	"github.com/bornjre/turnix/backend/modules/books/models"
	"github.com/k0kubun/pp"
	"github.com/upper/db/v4"
)

// fixme: add updated at everywhere

type SalesData struct {
	Sale  models.Sales       `json:"sale"`
	Lines []models.SalesLine `json:"lines"`
}

type ProductLite struct {
	ID         int64  `json:"id" db:"id,omitempty"`
	Name       string `json:"name" db:"name"`
	StockCount int64  `json:"stock_count" db:"stock_count"`
	Epoch      int64  `json:"epoch" db:"epoch"`
}

func (b *DbOps) SalesAdd(pid, uid int64, data *SalesData) (int64, error) {

	pp.Println("@data", data)

	if len(data.Lines) == 0 {
		return 0, fmt.Errorf("lines should not be empty")
	}

	err := b.userHasScope(pid, uid, "write")
	if err != nil {
		return 0, err
	}

	t := time.Now()

	sd := data.Sale
	sd.CreatedBy = uid
	sd.UpdatedBy = uid
	sd.CreatedAt = &t
	sd.UpdatedAt = &t
	sd.IsDeleted = false

	salesId := int64(0)

	productIds := make([]int64, 0, len(data.Lines))
	requiredProductCount := make(map[int64]int64)

	for _, line := range data.Lines {
		productIds = append(productIds, line.ProductID)
		oldQty := requiredProductCount[line.ProductID]
		requiredProductCount[line.ProductID] = oldQty + line.Qty
	}

	products := make([]ProductLite, 0, len(data.Lines))

	err = b.db.GetSession().TxContext(context.Background(), func(tx db.Session) error {

		productTable := tx.Collection(productTableName(pid))
		salesTable := tx.Collection(salesTableName(pid))
		salesLineTables := tx.Collection(salesLineTableName(pid))

		err = productTable.Find(db.Cond{
			"id IN":      productIds,
			"is_deleted": false,
		}).Select(
			"id",
			"stock_count",
			"epoch",
			"name",
		).All(&products)
		if err != nil {
			return err
		}

		for _, product := range products {
			requiredCount := requiredProductCount[product.ID]

			if requiredCount == 0 {
				return fmt.Errorf("product quantity should not be zero, %s", product.Name)
			}

			if requiredCount < product.StockCount {
				return fmt.Errorf("product %s stock not enough, required: %d, available: %d", product.Name, requiredCount, product.StockCount)
			}
		}

		r, err := salesTable.Insert(sd)
		if err != nil {
			return err
		}

		salesId = r.ID().(int64)

		for _, line := range data.Lines {

			line.SaleID = salesId
			line.CreatedBy = uid
			line.UpdatedBy = uid
			line.CreatedAt = &t
			line.UpdatedAt = &t
			line.IsDeleted = false

			_, err := salesLineTables.Insert(line)
			if err != nil {
				return err
			}

		}

		for _, product := range products {

			reqamt := requiredProductCount[product.ID]
			pamt := product.StockCount - reqamt

			err = productTable.Find(db.Cond{
				"id": product.ID,
			}).Update(db.Cond{
				"stock_count": pamt,
			})
			if err != nil {
				return err
			}
		}

		return nil
	}, nil)

	if err != nil {
		return 0, err
	}

	return salesId, nil
}

type SalesGetData struct {
	Sale  models.Sales       `json:"sale"`
	Lines []models.SalesLine `json:"lines"`
}

func (b *DbOps) SalesGet(pid, uid, id int64) (*SalesGetData, error) {
	err := b.userHasScope(pid, uid, "read")
	if err != nil {
		return nil, err
	}

	data := &SalesGetData{
		Sale:  models.Sales{},
		Lines: make([]models.SalesLine, 0),
	}

	salestable := b.salesTable(pid)

	err = salestable.Find(db.Cond{"id": id}).One(&data.Sale)
	if err != nil {
		return nil, err
	}

	saleLinetable := b.salesLineTable(pid)

	err = saleLinetable.Find(db.Cond{"sale_id": id}).All(&data.Lines)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (b *DbOps) SalesList(pid, uid int64) ([]models.Sales, error) {
	err := b.userHasScope(pid, uid, "read")
	if err != nil {
		return nil, err
	}

	datas := make([]models.Sales, 0)
	table := b.salesTable(pid)

	err = table.Find(db.Cond{"is_deleted": false}).All(&datas)
	if err != nil {
		return nil, err
	}

	return datas, nil
}

func (b *DbOps) SalesDelete(pid, uid, id int64) error {
	err := b.userHasScope(pid, uid, "write")
	if err != nil {
		return err
	}

	return b.salesTable(pid).Find(db.Cond{"id": id}).Update(db.Cond{
		"is_deleted": true,
	})

}

func (b *DbOps) salesTable(pid int64) db.Collection {
	return b.db.Table(salesTableName(pid))
}

func (b *DbOps) salesLineTable(pid int64) db.Collection {
	return b.db.Table(salesLineTableName(pid))
}

func salesTableName(pid int64) string {
	return fmt.Sprintf("z_%d_Sales", pid)
}

func salesLineTableName(pid int64) string {
	return fmt.Sprintf("z_%d_SalesLines", pid)
}

/*

batch inert example

	err = d.txOr(txid, func(sess db.Session) error {
		inserter := sess.SQL().InsertInto(
			d.tns.Table(req.TenantId, req.Group, req.Table),
		).Returning("__id").Batch(len(req.Data))

		for _, data := range req.Data {
			inserter.Values(data)
		}

		inserter.NextResult(&keyMap)

		inserter.Done()
		return inserter.Err()
	})


*/
