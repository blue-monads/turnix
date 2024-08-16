package dbops

/*

// sales

type SalesData struct {
	Title       string `json:"title"`
	ClientID    int64  `json:"client_id"`
	ClientName  string `json:"client_name"`
	Notes       string `json:"notes"`
	Attachments string `json:"attachments"`

	TotalItemPrice          float64 `json:"total_item_price"`
	TotalItemTaxAmount      float64 `json:"total_item_tax_amount"`
	TotalItemDiscountAmount float64 `json:"total_item_discount_amount"`

	SubTotal              float64         `json:"sub_total"`
	OverallDiscountAmount float64         `json:"overall_discount_amount"`
	OverallTaxAmount      float64         `json:"overall_tax_amount"`
	Total                 float64         `json:"total"`
	Lines                 []SalesLineData `json:"lines"`
}

type SalesLineData struct {
	Info           string  `json:"info"`
	Qty            float64 `json:"qty"`
	SaleID         int64   `json:"sale_id"`
	ProductID      int64   `json:"product_id"`
	Price          float64 `json:"price"`
	TaxAmount      float64 `json:"tax_amount"`
	DiscountAmount float64 `json:"discount_amount"`
	TotalAmount    float64 `json:"total_amount"`
}

type ProductLite struct {
	ID         int64 `json:"id" db:"id,omitempty"`
	StockCount int64 `json:"stock_count" db:"stock_count"`
	Epoch      int64 `json:"epoch" db:"epoch"`
}

func (b *BookModule) dbOpsSalesAdd(pid, uid int64, data *SalesData) (int64, error) {
	err := b.userHasScope(pid, uid, "write")
	if err != nil {
		return 0, err
	}

	t := time.Now()

	sd := Sales{
		Title:       data.Title,
		ClientID:    data.ClientID,
		ClientName:  data.ClientName,
		Notes:       data.Notes,
		Attachments: data.Attachments,

		TotalItemPrice:          data.TotalItemPrice,
		TotalItemTaxAmount:      data.TotalItemTaxAmount,
		TotalItemDiscountAmount: data.TotalItemDiscountAmount,

		SubTotal:              data.SubTotal,
		OverallDiscountAmount: data.OverallDiscountAmount,
		OverallTaxAmount:      data.OverallTaxAmount,
		Total:                 data.Total,

		CreatedBy: uid,
		UpdatedBy: uid,
		CreatedAt: &t,
		UpdatedAt: &t,
		IsDeleted: false,
	}

	salesId := int64(0)

	productTable := b.productTable(pid)

	// salelines product ids

	productIds := make([]int64, 0, len(data.Lines))

	for _, line := range data.Lines {
		productIds = append(productIds, line.ProductID)
	}

	products := make([]ProductLite, 0, len(data.Lines))

	if err != nil {
		return 0, err
	}

	if len(products) != len(productIds) {
		return 0, errors.New("product not found")
	}

	b.db.GetSession().TxContext(context.Background(), func(tx db.Session) error {

		salesTable := tx.Collection("z_%d_Sales")
		salesLineTables := tx.Collection("z_%d_SalesLines")

		r, err := salesTable.Insert(sd)
		if err != nil {
			return err
		}

		salesId = r.ID().(int64)

		for _, line := range data.Lines {

			saleLine := SalesLine{
				Info:           line.Info,
				Qty:            line.Qty,
				SaleID:         salesId,
				ProductID:      line.ProductID,
				Price:          line.Price,
				TaxAmount:      line.TaxAmount,
				DiscountAmount: line.DiscountAmount,
				TotalAmount:    line.TotalAmount,
				CreatedBy:      uid,
				UpdatedBy:      uid,
				CreatedAt:      &t,
				UpdatedAt:      &t,
				IsDeleted:      false,
			}

			_, err := salesLineTables.Insert(saleLine)
			if err != nil {
				return err
			}

		}

		return nil
	}, nil)

	table := b.salesTable(pid)

	r, err := table.Insert(data)
	if err != nil {
		return 0, err
	}

	salesId = r.ID().(int64)

	return r.ID().(int64), nil
}

func (b *BookModule) dbOpsSalesUpdate(pid, uid, id int64, data map[string]any) error {
	err := b.userHasScope(pid, uid, "write")
	if err != nil {
		return err
	}

	return b.salesTable(pid).Find(db.Cond{"id": id}).Update(data)
}

func (b *BookModule) dbOpsSalesGet(pid, uid, id int64) (*Sales, error) {
	err := b.userHasScope(pid, uid, "read")
	if err != nil {
		return nil, err
	}

	data := &Sales{}
	table := b.salesTable(pid)

	err = table.Find(db.Cond{"id": id}).One(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (b *BookModule) dbOpsSalesList(pid, uid int64) ([]Sales, error) {
	err := b.userHasScope(pid, uid, "read")
	if err != nil {
		return nil, err
	}

	datas := make([]Sales, 0)
	table := b.salesTable(pid)

	err = table.Find().All(&datas)
	if err != nil {
		return nil, err
	}

	return datas, nil
}

func (b *BookModule) salesTable(pid int64) db.Collection {
	return b.db.Table(fmt.Sprintf("z_%d_Sales", pid))
}

func (b *BookModule) salesLineTable(pid int64) db.Collection {
	return b.db.Table(fmt.Sprintf("z_%d_SalesLines", pid))
}



*/

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
