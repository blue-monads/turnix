package dbops

import (
	"database/sql"
	"fmt"

	"github.com/bornjre/turnix/backend/modules/books/models"
	"github.com/jmoiron/sqlx"
	"github.com/k0kubun/pp"
	"github.com/upper/db/v4"
)

// transactions

func (b *DbOps) AddTxn(pid, uid int64, data *models.Transaction) (int64, error) {

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

func (b *DbOps) UpdateTxn(pid, uid, id int64, data map[string]any) error {
	err := b.userHasScope(pid, uid, "write")
	if err != nil {
		return err

	}

	pp.Println("@dbOpUpdateTxn", data)

	return b.txnTable(pid).Find(db.Cond{"id": id}).Update(data)
}

func (b *DbOps) GetTxn(pid, uid, tid int64) (*models.Transaction, error) {

	pp.Println("@dbOpGetTxn", pid, uid, tid)

	err := b.userHasScope(pid, uid, "read")
	if err != nil {
		pp.Println("@userHasScope", pid, uid, tid)

		return nil, err
	}

	data := &models.Transaction{}
	table := b.txnTable(pid)

	err = table.Find(db.Cond{"id": tid}).One(data)
	if err != nil {
		return nil, err
	}

	return data, nil

}

func (b *DbOps) GetTxnWithLine(pid, uid, tid int64) (*models.TransactionWithLine, error) {

	err := b.userHasScope(pid, uid, "read")
	if err != nil {
		return nil, err
	}

	txnTable := b.txnTable(pid)
	lineTable := b.txnLineTable(pid)

	data := &models.TransactionWithLine{}

	lines := []models.TransactionLine{}

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

func (b *DbOps) ListTxn(pid, uid int64) ([]models.Transaction, error) {

	err := b.userHasScope(pid, uid, "read")
	if err != nil {
		return nil, err
	}

	datas := make([]models.Transaction, 0)
	table := b.txnTable(pid)

	err = table.Find().All(&datas)
	if err != nil {
		return nil, err
	}

	return datas, nil

}

type TxnRecords struct {
	Transactions []models.Transaction     `json:"transactions"`
	Lines        []models.TransactionLine `json:"lines"`
}

func (b *DbOps) ListTxnWithLines(pid, uid, offset int64) (*TxnRecords, error) {

	err := b.userHasScope(pid, uid, "read")
	if err != nil {
		return nil, err
	}

	record := &TxnRecords{
		Transactions: make([]models.Transaction, 0),
		Lines:        make([]models.TransactionLine, 0),
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

func (b *DbOps) ListAccountTxnWithLines(pid, uid, accId, offset int64) (*TxnRecords, error) {

	err := b.userHasScope(pid, uid, "read")
	if err != nil {
		return nil, err
	}

	record := &TxnRecords{
		Transactions: make([]models.Transaction, 0),
		Lines:        make([]models.TransactionLine, 0),
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
		z_%d_TransactionLines t1, z_%d_TransactionLines t2
	INNER JOIN
		z_%d_Transactions t ON t.id = t1.txn_id
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

	results := []models.TransactionResult{}
	err = sqlx.StructScan(rows, &results)
	if err != nil {
		return nil, err
	}

	for _, result := range results {
		record.Transactions = append(record.Transactions, models.Transaction{
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

		record.Lines = append(record.Lines, models.TransactionLine{
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

		record.Lines = append(record.Lines, models.TransactionLine{
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

func (b *DbOps) DeleteTxn(pid, uid, aid int64) error {

	err := b.userHasScope(pid, uid, "write")
	if err != nil {
		return err
	}

	table := b.txnTable(pid)
	err = table.Find(db.Cond{"id": aid}).Update(db.Cond{
		"is_deleted": true,
	})
	if err != nil {
		return err
	}

	return nil

}

// txn line

func (b *DbOps) AddTxnLine(pid, uid int64, data *models.TransactionLine) (int64, error) {
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

func (b *DbOps) UpdateTxnLine(pid, uid, txnId, id int64, data map[string]any) error {
	err := b.userHasScope(pid, uid, "write")
	if err != nil {
		return err
	}

	pp.Println("@dbOpUpdateTxnLine", pid, uid, txnId, id, data)

	return b.txnLineTable(pid).Find(db.Cond{"id": id, "txn_id": txnId}).Update(data)
}
