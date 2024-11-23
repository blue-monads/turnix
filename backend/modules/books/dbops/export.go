package dbops

import (
	"time"

	"github.com/blue-monads/turnix/backend/modules/books/models"
	"github.com/k0kubun/pp"
	"github.com/upper/db/v4"
)

type ExportData struct {
	Accounts         []models.Account         `json:"accounts"`
	Transactions     []models.Transaction     `json:"transactions"`
	TransactionLines []models.TransactionLine `json:"transaction_lines"`

	Catagories []models.Catagory `json:"catagories"`
	Products   []models.Product  `json:"products"`
	Taxes      []models.Tax      `json:"taxes"`
	Contacts   []models.Contact  `json:"contacts"`
}

type ImportOptions struct {
	Data     ExportData `json:"data"`
	AsNewTxn bool       `json:"as_new_txn"`
}

func (b *DbOps) Export(pid, uid int64) (*ExportData, error) {

	pp.Println("@dbOpsExport/1")

	err := b.userHasScope(pid, uid, "read")
	if err != nil {
		return nil, err
	}

	pp.Println("@dbOpsExport/2")

	data := ExportData{
		Accounts:         make([]models.Account, 0),
		Transactions:     make([]models.Transaction, 0),
		TransactionLines: make([]models.TransactionLine, 0),
		Catagories:       make([]models.Catagory, 0),
		Products:         make([]models.Product, 0),
		Taxes:            make([]models.Tax, 0),
		Contacts:         make([]models.Contact, 0),
	}

	accTable := b.accountsTable(pid)
	txnTable := b.txnTable(pid)
	txnLineTable := b.txnLineTable(pid)
	catagoryTable := b.catagoryTable(pid)
	productTable := b.productTable(pid)
	contactTable := b.contactTable(pid)
	taxTable := b.taxTable(pid)

	pp.Println("@dbOpsExport/3")

	err = accTable.Find().All(&data.Accounts)
	if err != nil {
		return nil, err
	}

	pp.Println("@dbOpsExport/4")

	err = txnTable.Find().All(&data.Transactions)
	if err != nil {
		return nil, err
	}

	pp.Println("@dbOpsExport/5")

	err = txnLineTable.Find().All(&data.TransactionLines)
	if err != nil {
		return nil, err
	}
	pp.Println("@dbOpsExport/6")

	err = catagoryTable.Find().All(&data.Catagories)
	if err != nil {
		return nil, err
	}

	err = productTable.Find().All(&data.Products)
	if err != nil {
		return nil, err
	}

	err = contactTable.Find().All(&data.Contacts)
	if err != nil {
		return nil, err

	}

	err = taxTable.Find().All(&data.Taxes)
	if err != nil {
		return nil, err
	}

	return &data, nil
}

func (b *DbOps) Import(pid, uid int64, opts ImportOptions) (err error) {

	pp.Println("@dbOpsImport/1")

	err = b.userHasScope(pid, uid, "write")
	if err != nil {
		return err
	}

	pp.Println("@dbOpsImport/2")

	accTable := b.accountsTable(pid)
	txnTable := b.txnTable(pid)
	txnLineTable := b.txnLineTable(pid)
	catagoryTable := b.catagoryTable(pid)
	productTable := b.productTable(pid)
	contactTable := b.contactTable(pid)
	taxTable := b.taxTable(pid)

	pp.Println("@dbOpsImport/3")

	for _, acc := range opts.Data.Accounts {

		pp.Println("@dbOpsImport/4", acc.Name)

		oldAcc, _ := b.GetAccount(pid, uid, acc.ID)
		pp.Println("@dbOpsImport/5")

		if oldAcc != nil {
			pp.Println("@dbOpsImport/6/accSkipp")

			continue
		}

		pp.Println("@dbOpsImport/7")

		r, err := accTable.Insert(acc)
		if err != nil {
			pp.Println("@dbOpsImport/8")

			return err
		}

		pp.Println("@dbOpsImport/9")

		acc.ID = r.ID().(int64)
	}

	inverseMap := make(map[int64]int64)
	txnLines := make([]int64, 0)

	defer func() {
		pp.Println("@defer")
		if err == nil {
			return
		}

		pp.Println("@defer")

		// cleanup just insered data
		keys := make([]int64, 0, len(inverseMap))
		for k := range inverseMap {
			keys = append(keys, k)
		}

		accTable.Find(db.Cond{"id in ?": keys}).Delete()
		txnLineTable.Find(db.Cond{"id in ?": txnLines}).Delete()

		pp.Println("@defer/end")

	}()

	pp.Println("@dbOpsImport/10")

	for _, txn := range opts.Data.Transactions {
		pp.Println("@dbOpsImport/11", txn.ID, txn.Title)

		t := time.Now()
		txn.CreatedAt = &t
		txn.UpdatedAt = &t
		txn.CreatedBy = uid
		txn.UpdatedBy = uid

		odlTxnId := txn.ID
		if opts.AsNewTxn {
			txn.ID = 0
		}

		r, err := txnTable.Insert(txn)
		if err != nil {
			continue
		}

		newTxnId := r.ID().(int64)
		inverseMap[odlTxnId] = newTxnId

	}

	pp.Println("@dbOpsImport/12")

	for _, line := range opts.Data.TransactionLines {
		txnId, ok := inverseMap[line.TxnID]
		if !ok {
			pp.Println("@dbOpsImport/13/skipping", txnId, line.ID)
			continue
		}

		pp.Println("@dbOpsImport/13", line.ID)
		t := time.Now()
		line.TxnID = txnId
		line.CreatedAt = &t
		line.UpdatedAt = &t
		line.CreatedBy = uid
		line.UpdatedBy = uid

		r, err := txnLineTable.Insert(line)
		if err != nil {
			return err
		}

		pp.Println("@dbOpsImport/14")

		txnLines = append(txnLines, r.ID().(int64))
	}

	pp.Println("@dbOpsImport/15")

	// restore catagories

	catagoryIdIndex := make(map[int64]int64)
	for _, cat := range opts.Data.Catagories {
		catagoryIdIndex[cat.ID] = cat.ID
	}

	for _, cat := range opts.Data.Catagories {

		exists, err := catagoryTable.Find(db.Cond{"id": cat.ID}).Exists()
		if err != nil {
			return err
		}

		if exists {
			continue
		}

		t := time.Now()

		cat.CreatedAt = &t
		cat.UpdatedAt = &t
		cat.CreatedBy = uid
		cat.UpdatedBy = uid
		r, err := catagoryTable.Insert(cat)
		if err != nil {
			return err
		}

		catagoryIdIndex[cat.ID] = r.ID().(int64)
	}

	for _, prod := range opts.Data.Products {

		exists, err := productTable.Find(db.Cond{"id": prod.ID}).Exists()
		if err != nil {
			return err
		}

		if exists {
			continue
		}

		t := time.Now()

		prod.CreatedAt = &t
		prod.UpdatedAt = &t
		prod.CreatedBy = uid
		prod.UpdatedBy = uid
		prod.CatagoryID = catagoryIdIndex[prod.CatagoryID]
		_, err = productTable.Insert(prod)
		if err != nil {
			return err
		}

	}

	for _, tax := range opts.Data.Taxes {
		t := time.Now()

		tax.CreatedAt = &t
		tax.UpdatedAt = &t
		tax.CreatedBy = uid
		tax.UpdatedBy = uid
		_, err = taxTable.Insert(tax)
		if err != nil {
			return err
		}
	}

	for _, contact := range opts.Data.Contacts {
		t := time.Now()

		contact.CreatedAt = &t
		contact.UpdatedAt = &t
		contact.CreatedBy = uid
		contact.UpdatedBy = uid
		_, err = contactTable.Insert(contact)
		if err != nil {
			return err
		}
	}

	return nil
}
