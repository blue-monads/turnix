package books

import (
	"database/sql"
	"fmt"
	"strings"
	"time"

	"github.com/bornjre/turnix/backend/xtypes/services/xdatabase"
	"github.com/jmoiron/sqlx"
	"github.com/k0kubun/pp"
	"github.com/upper/db/v4"
)

// account

func (b *BookModule) dbOpAddAccount(pid, uid int64, data *Account) (int64, error) {

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

func (b *BookModule) dbOpUpdateAccount(pid, uid, id int64, data map[string]any) error {
	err := b.userHasScope(pid, uid, "write")
	if err != nil {
		return err

	}

	return b.accountsTable(pid).Find(db.Cond{"id": id}).Update(data)
}

func (b *BookModule) dbOpGetAccount(pid, uid, aid int64) (*Account, error) {

	err := b.userHasScope(pid, uid, "read")
	if err != nil {
		return nil, err
	}

	data := &Account{}
	table := b.accountsTable(pid)

	err = table.Find(db.Cond{"id": aid}).One(data)
	if err != nil {
		return nil, err
	}

	return data, nil

}

func (b *BookModule) dbOpListAccount(pid, uid int64) ([]Account, error) {

	err := b.userHasScope(pid, uid, "read")
	if err != nil {
		return nil, err
	}

	datas := make([]Account, 0)
	table := b.accountsTable(pid)

	err = table.Find().All(&datas)
	if err != nil {
		return nil, err
	}

	return datas, nil

}

func (b *BookModule) dbOpDeleteAccount(pid, uid, aid int64) error {

	err := b.userHasScope(pid, uid, "write")
	if err != nil {
		return err
	}

	table := b.accountsTable(pid)
	err = table.Find(db.Cond{"id": aid}).Delete()
	if err != nil {
		return err
	}

	// fixme cleanup txns

	return nil

}

// transactions

func (b *BookModule) dbOpAddTxn(pid, uid int64, data *Transaction) (int64, error) {

	err := b.userHasScope(pid, uid, "write")
	if err != nil {
		return 0, err
	}

	table := b.txnTable(pid)

	r, err := table.Insert(data)
	if err != nil {
		return 0, err
	}

	return r.ID().(int64), nil
}

func (b *BookModule) dbOpUpdateTxn(pid, uid, id int64, data map[string]any) error {
	err := b.userHasScope(pid, uid, "write")
	if err != nil {
		return err

	}

	pp.Println("@dbOpUpdateTxn", data)

	return b.txnTable(pid).Find(db.Cond{"id": id}).Update(data)
}

func (b *BookModule) dbOpGetTxn(pid, uid, tid int64) (*Transaction, error) {

	pp.Println("@dbOpGetTxn", pid, uid, tid)

	err := b.userHasScope(pid, uid, "read")
	if err != nil {
		pp.Println("@userHasScope", pid, uid, tid)

		return nil, err
	}

	data := &Transaction{}
	table := b.txnTable(pid)

	err = table.Find(db.Cond{"id": tid}).One(data)
	if err != nil {
		return nil, err
	}

	return data, nil

}

func (b *BookModule) dbOpGetTxnWithLine(pid, uid, tid int64) (*TransactionWithLine, error) {

	err := b.userHasScope(pid, uid, "read")
	if err != nil {
		return nil, err
	}

	txnTable := b.txnTable(pid)
	lineTable := b.txnLineTable(pid)

	data := &TransactionWithLine{}

	lines := []TransactionLine{}

	err = txnTable.Find(db.Cond{"id": tid}).One(&data.Txn)
	if err != nil {
		return nil, err
	}

	err = lineTable.Find(db.Cond{"txn_id": tid}).All(&lines)
	if err != nil {
		return nil, err
	}

	data.FirstLine = &lines[0]
	data.SecondLine = &lines[1]

	return data, nil

}

func (b *BookModule) dbOpListTxn(pid, uid int64) ([]Transaction, error) {

	err := b.userHasScope(pid, uid, "read")
	if err != nil {
		return nil, err
	}

	datas := make([]Transaction, 0)
	table := b.txnTable(pid)

	err = table.Find().All(&datas)
	if err != nil {
		return nil, err
	}

	return datas, nil

}

type TxnRecords struct {
	Transactions []Transaction     `json:"transactions"`
	Lines        []TransactionLine `json:"lines"`
}

func (b *BookModule) dbOpListTxnWithLines(pid, uid, offset int64) (*TxnRecords, error) {

	err := b.userHasScope(pid, uid, "read")
	if err != nil {
		return nil, err
	}

	record := &TxnRecords{
		Transactions: make([]Transaction, 0),
		Lines:        make([]TransactionLine, 0),
	}

	txtable := b.txnTable(pid)

	err = txtable.Find(db.Cond{"id >": offset}).Limit(100).All(&record.Transactions)
	if err != nil {
		return nil, err
	}

	lineTable := b.txnLineTable(pid)

	err = lineTable.Find(db.Cond{
		"txn_id >": offset,
		"txn_id <": offset + 101,
	}).All(&record.Lines)
	if err != nil {
		return nil, err
	}

	return record, nil

}

func (b *BookModule) dbOpListAccountTxnWithLines(pid, uid, accId, offset int64) (*TxnRecords, error) {

	err := b.userHasScope(pid, uid, "read")
	if err != nil {
		return nil, err
	}

	record := &TxnRecords{
		Transactions: make([]Transaction, 0),
		Lines:        make([]TransactionLine, 0),
	}

	driver := b.db.GetSession().Driver().(*sql.DB)

	stmt, err := driver.Prepare(fmt.Sprintf(`	
	SELECT
		t1.id AS first_id, t1.account_id AS first_account_id,
		t1.debit_amount AS first_debit_amount, t1.credit_amount AS first_credit_amount,
		t1.created_by AS first_created_by, t1.updated_by AS first_updated_by,
		t1.created_at AS first_created_at, t1.updated_at AS first_updated_at,
		t2.id AS second_id, t2.account_id AS second_account_id,
		t2.debit_amount AS second_debit_amount, t2.credit_amount AS second_credit_amount,
		t2.created_by AS second_created_by, t2.updated_by AS second_updated_by,
		t2.created_at AS second_created_at, t2.updated_at AS second_updated_at,
		t.id AS id, t.title, t.notes, t.linked_sales_id, t.linked_invoice_id,
		t.reference_id, t.attachments, t.created_by AS txn_created_by,
		t.updated_by AS txn_updated_by, t.created_at AS txn_created_at,
		t.updated_at AS txn_updated_at, t.is_deleted
	FROM
		TransactionLines_%d_ t1, TransactionLines_%d_ t2
	INNER JOIN
		Transactions_%d_ t ON t.id = t1.txn_id
	WHERE
		t1.account_id = ? AND t1.txn_id = t2.txn_id AND t2.account_id <> ? and t.is_deleted = FALSE and t.id > ? ORDER BY t.id LIMIT 100;
		`, pid, pid, pid))

	if err != nil {
		return nil, err
	}

	rows, err := stmt.Query(accId, accId, offset)
	if err != nil {
		return nil, err
	}

	stmt.Close()

	results := []TransactionResult{}
	err = sqlx.StructScan(rows, &results)
	if err != nil {
		return nil, err
	}

	for _, result := range results {
		record.Transactions = append(record.Transactions, Transaction{
			ID:              result.Id,
			Title:           result.Title,
			Notes:           result.Notes,
			LinkedSalesID:   int64(result.LinkedSalesID),
			LinkedInvoiceID: int64(result.LinkedSalesID),
			ReferenceID:     result.ReferenceID,
			Attachments:     result.Attachments,
			CreatedBy:       int64(result.TxnCreatedBy),
			UpdatedBy:       int64(result.TxnUpdatedBy),
			CreatedAt:       result.TxnCreatedAt,
			UpdatedAt:       result.TxnUpdatedAt,
			IsDeleted:       false,
		})

		record.Lines = append(record.Lines, TransactionLine{
			ID:           int64(result.FirstID),
			AccountID:    int64(result.FirstAccountID),
			TxnID:        int64(result.Id),
			DebitAmount:  result.FirstDebitAmount,
			CreditAmount: result.FirstCreditAmount,
			CreatedBy:    int64(result.FirstCreatedBy),
			UpdatedBy:    int64(result.FirstUpdatedBy),
			CreatedAt:    result.FirstCreatedAt,
			UpdatedAt:    result.SecondUpdatedAt,
		})

		record.Lines = append(record.Lines, TransactionLine{
			ID:           int64(result.SecondID),
			AccountID:    int64(result.SecondAccountID),
			TxnID:        int64(result.Id),
			DebitAmount:  result.SecondDebitAmount,
			CreditAmount: result.SecondCreditAmount,
			CreatedBy:    int64(result.SecondCreatedBy),
			UpdatedBy:    int64(result.SecondUpdatedBy),
			CreatedAt:    result.SecondCreatedAt,
			UpdatedAt:    result.SecondUpdatedAt,
		})

	}

	return record, nil

}

func (b *BookModule) dbOpDeleteTxn(pid, uid, aid int64) error {

	err := b.userHasScope(pid, uid, "write")
	if err != nil {
		return err
	}

	table := b.txnTable(pid)
	err = table.Find(db.Cond{"id": aid}).Delete()
	if err != nil {
		return err
	}

	return nil

}

// txn line

func (b *BookModule) dbOpAddTxnLine(pid, uid int64, data *TransactionLine) (int64, error) {
	err := b.userHasScope(pid, uid, "write")
	if err != nil {
		return 0, err
	}

	table := b.txnLineTable(pid)

	r, err := table.Insert(data)
	if err != nil {
		return 0, err
	}

	return r.ID().(int64), nil
}

func (b *BookModule) dbOpUpdateTxnLine(pid, uid, txnId, id int64, data map[string]any) error {
	err := b.userHasScope(pid, uid, "write")
	if err != nil {
		return err
	}

	pp.Println("@dbOpUpdateTxnLine", pid, uid, txnId, id, data)

	return b.txnLineTable(pid).Find(db.Cond{"id": id, "txn_id": txnId}).Update(data)
}

// report

func (b *BookModule) dbOptsGenerateReportLongLedger(pid, uid int64, opts ReportOptions) ([]LongLedgerRecord, error) {
	err := b.userHasScope(pid, uid, "read")
	if err != nil {
		return nil, err
	}

	records := []LongLedgerRecord{}

	rows, err := b.db.GetSession().SQL().Query(fmt.Sprintf(`
SELECT tx.title, tx.id as txn_id, a.name as account_name, a.acc_type,  tl.debit_amount, tl.credit_amount, SUM( tl.debit_amount) OVER (PARTITION BY tl.account_id) as total_debit, SUM(tl.credit_amount) OVER (PARTITION BY tl.account_id) as total_credit, tl.account_id
FROM Accounts_%d_ a
	JOIN TransactionLines_%d_ tl ON tl.account_id = a.id
	JOIN Transactions_%d_ tx on tl.txn_id = tx.id
WHERE
	a.is_deleted = FALSE AND
	tx.is_deleted = FALSE		
ORDER BY tl.account_id;

`, pid, pid, pid))

	if err != nil {
		return nil, err
	}

	err = sqlx.StructScan(rows, &records)
	if err != nil {
		return nil, err
	}

	return records, nil
}

func (b *BookModule) dbOptsGenerateReportShortLedger(pid, uid int64, opts ReportOptions) ([]ShortLedgerRecord, error) {
	err := b.userHasScope(pid, uid, "read")
	if err != nil {
		return nil, err
	}

	records := []ShortLedgerRecord{}

	rows, err := b.db.GetSession().SQL().Query(fmt.Sprintf(`
SELECT
	a.id as account_id,
	a.name as account_name,
	a.acc_type,
	SUM( tl.debit_amount) as  total_debit,
	SUM(tl.credit_amount) as total_credit
FROM TransactionLines_%d_ tl
INNER JOIN Accounts_%d_ a ON tl.account_id = a.id
INNER JOIN Transactions_%d_ t ON tl.txn_id = t.id
WHERE 
	t.is_deleted = FALSE AND
	a.is_deleted = FALSE 
GROUP BY tl.account_id;

`, pid, pid, pid))

	if err != nil {
		return nil, err
	}

	err = sqlx.StructScan(rows, &records)
	if err != nil {
		return nil, err
	}

	return records, nil
}

// utils

func (b *BookModule) dbOpsExport(pid, uid int64) (*ExportData, error) {

	pp.Println("@dbOpsExport/1")

	err := b.userHasScope(pid, uid, "read")
	if err != nil {
		return nil, err
	}

	pp.Println("@dbOpsExport/2")

	data := ExportData{
		Accounts:         make([]Account, 0),
		Transactions:     make([]Transaction, 0),
		TransactionLines: make([]TransactionLine, 0),
	}

	accTable := b.accountsTable(pid)
	txnTable := b.txnTable(pid)
	txnLineTable := b.txnLineTable(pid)

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

	return &data, nil
}

func (b *BookModule) dbOpsImport(pid, uid int64, opts ImportOptions) (err error) {

	pp.Println("@dbOpsImport/1")

	err = b.userHasScope(pid, uid, "write")
	if err != nil {
		return err
	}

	pp.Println("@dbOpsImport/2")

	accTable := b.accountsTable(pid)
	txnTable := b.txnTable(pid)
	txnLineTable := b.txnLineTable(pid)

	pp.Println("@dbOpsImport/3")

	for _, acc := range opts.Data.Accounts {

		pp.Println("@dbOpsImport/4", acc.Name)

		oldAcc, _ := b.dbOpGetAccount(pid, uid, acc.ID)
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

	return nil
}

func (b *BookModule) dbOpsCatagoryAdd(pid, uid int64, data *Catagory) (int64, error) {
	err := b.userHasScope(pid, uid, "write")
	if err != nil {
		return 0, err
	}

	table := b.catagoryTable(pid)

	r, err := table.Insert(data)
	if err != nil {
		return 0, err
	}

	return r.ID().(int64), nil
}

func (b *BookModule) dbOpsCatagoryUpdate(pid, uid, id int64, data map[string]any) error {
	err := b.userHasScope(pid, uid, "write")
	if err != nil {
		return err
	}

	return b.catagoryTable(pid).Find(db.Cond{"id": id}).Update(data)
}

func (b *BookModule) dbOpsCatagoryGet(pid, uid, id int64) (*Catagory, error) {
	err := b.userHasScope(pid, uid, "read")
	if err != nil {
		return nil, err
	}

	data := &Catagory{}
	table := b.catagoryTable(pid)

	err = table.Find(db.Cond{"id": id}).One(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (b *BookModule) dbOpsCatagoryList(pid, uid int64) ([]Catagory, error) {
	err := b.userHasScope(pid, uid, "read")
	if err != nil {
		return nil, err
	}

	datas := make([]Catagory, 0)
	table := b.catagoryTable(pid)

	err = table.Find().All(&datas)
	if err != nil {
		return nil, err
	}

	return datas, nil
}

func (b *BookModule) dbOpsCatagoryDelete(pid, uid, id int64) error {
	err := b.userHasScope(pid, uid, "write")
	if err != nil {
		return err
	}

	return b.catagoryTable(pid).Find(db.Cond{"id": id}).Delete()
}

// products

func (b *BookModule) dbOpsProductAdd(pid, uid int64, data *Product) (int64, error) {
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

func (b *BookModule) dbOpsProductUpdate(pid, uid, id int64, data map[string]any) error {
	err := b.userHasScope(pid, uid, "write")
	if err != nil {
		return err
	}

	return b.productTable(pid).Find(db.Cond{"id": id}).Update(data)
}

func (b *BookModule) dbOpsProductGet(pid, uid, id int64) (*Product, error) {
	err := b.userHasScope(pid, uid, "read")
	if err != nil {
		return nil, err
	}

	data := &Product{}
	table := b.productTable(pid)

	err = table.Find(db.Cond{"id": id}).One(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (b *BookModule) dbOpsProductList(pid, uid int64) ([]Product, error) {
	err := b.userHasScope(pid, uid, "read")
	if err != nil {
		return nil, err
	}

	datas := make([]Product, 0)
	table := b.productTable(pid)

	err = table.Find().All(&datas)
	if err != nil {
		return nil, err
	}

	return datas, nil
}

func (b *BookModule) dbOpsProductDelete(pid, uid, id int64) error {
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

func (b *BookModule) dbOpsProductStockInAdd(pid, uid int64, data *StockInData) (pstockInid int64, err error) {
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

		line := &ProductStockInLine{
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

	return pstockInid, nil
}

// tax

func (b *BookModule) dbOpsTaxAdd(pid, uid int64, data *Tax) (int64, error) {
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

func (b *BookModule) dbOpsTaxUpdate(pid, uid, id int64, data map[string]any) error {
	err := b.userHasScope(pid, uid, "write")
	if err != nil {
		return err
	}

	return b.taxTable(pid).Find(db.Cond{"id": id}).Update(data)
}

func (b *BookModule) dbOpsTaxGet(pid, uid, id int64) (*Tax, error) {
	err := b.userHasScope(pid, uid, "read")
	if err != nil {
		return nil, err
	}

	data := &Tax{}
	table := b.taxTable(pid)

	err = table.Find(db.Cond{"id": id}).One(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (b *BookModule) dbOpsTaxList(pid, uid int64) ([]Tax, error) {
	err := b.userHasScope(pid, uid, "read")
	if err != nil {
		return nil, err
	}

	datas := make([]Tax, 0)
	table := b.taxTable(pid)

	err = table.Find().All(&datas)
	if err != nil {
		return nil, err
	}

	return datas, nil
}

func (b *BookModule) dbOpsTaxDelete(pid, uid, id int64) error {
	err := b.userHasScope(pid, uid, "write")
	if err != nil {
		return err
	}

	return b.taxTable(pid).Find(db.Cond{"id": id}).Delete()
}

// utils

func (b *BookModule) txnTable(pid int64) db.Collection {
	return b.db.Table(fmt.Sprintf("Transactions_%d_", pid))
}

func (b *BookModule) accountsTable(pid int64) db.Collection {
	return b.db.Table(fmt.Sprintf("Accounts_%d_", pid))
}

func (b *BookModule) txnLineTable(pid int64) db.Collection {
	return b.db.Table(fmt.Sprintf("TransactionLines_%d_", pid))
}

func (b *BookModule) catagoryTable(pid int64) db.Collection {
	return b.db.Table(fmt.Sprintf("Catagories_%d_", pid))
}

func (b *BookModule) productTable(pid int64) db.Collection {
	return b.db.Table(fmt.Sprintf("Products_%d_", pid))
}

func (b *BookModule) productStockInTable(pid int64) db.Collection {
	return b.db.Table(fmt.Sprintf("ProductStockIn_%d_", pid))
}

func (b *BookModule) productStockInLineTable(pid int64) db.Collection {
	return b.db.Table(fmt.Sprintf("ProductStockInLines_%d_", pid))
}

/*
func (b *BookModule) salesTable(pid int64) db.Collection {
	return b.db.Table(fmt.Sprintf("Sales_%d_", pid))
}

func (b *BookModule) salesLineTable(pid int64) db.Collection {
	return b.db.Table(fmt.Sprintf("SalesLines_%d_", pid))
}

func (b *BookModule) invoiceTable(pid int64) db.Collection {
	return b.db.Table(fmt.Sprintf("Invoices_%d_", pid))
}

func (b *BookModule) invoiceLineTable(pid int64) db.Collection {
	return b.db.Table(fmt.Sprintf("InvoiceLines_%d_", pid))
}

func (b *BookModule) estimateTable(pid int64) db.Collection {
	return b.db.Table(fmt.Sprintf("Estimates_%d_", pid))
}

func (b *BookModule) estimateLineTable(pid int64) db.Collection {
	return b.db.Table(fmt.Sprintf("EstimateLines_%d_", pid))
}

*/

func (b *BookModule) taxTable(pid int64) db.Collection {
	return b.db.Table(fmt.Sprintf("Tax_%d_", pid))
}

func (b *BookModule) userHasScope(pid, uid int64, scope string) error {
	dbscope, err := b.db.GetProjectUserScope(uid, pid)
	if err != nil {
		return err
	}

	if dbscope != xdatabase.ScopeOwner && strings.Contains(dbscope, scope) {
		return xdatabase.ErrUserNoScope
	}

	return nil
}
