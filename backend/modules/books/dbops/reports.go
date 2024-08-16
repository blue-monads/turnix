package dbops

import (
	"fmt"
	"time"

	"github.com/bornjre/turnix/backend/modules/books/models"
	"github.com/jmoiron/sqlx"
	"github.com/k0kubun/pp"
)

type ReportOptions struct {
	ReportType string     `json:"report_type"`
	TemplateId int64      `json:"template_id"`
	FromDate   *time.Time `json:"from_date"`
	ToDate     *time.Time `json:"to_date"`
}

// report

func (b *DbOps) GenerateReportLongLedger(pid, uid int64, opts ReportOptions) ([]models.LongLedgerRecord, error) {
	err := b.userHasScope(pid, uid, "read")
	if err != nil {
		return nil, err
	}

	pp.Println("@ops", opts)

	records := []models.LongLedgerRecord{}

	rows, err := b.db.GetSession().SQL().Query(fmt.Sprintf(`
SELECT 
	tx.title, 
	tx.id as txn_id, 
	a.name as account_name, 
	a.acc_type,  
	tl.debit_amount, 
	tl.credit_amount, 
	SUM( tl.debit_amount) OVER (PARTITION BY tl.account_id) as total_debit, 
	SUM(tl.credit_amount) OVER (PARTITION BY tl.account_id) as total_credit, 
	tl.account_id
FROM z_%d_Accounts a
	JOIN z_%d_TransactionLines tl ON tl.account_id = a.id
	JOIN z_%d_Transactions tx on tl.txn_id = tx.id
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

func (b *DbOps) GenerateReportShortLedger(pid, uid int64, opts ReportOptions) ([]models.ShortLedgerRecord, error) {
	err := b.userHasScope(pid, uid, "read")
	if err != nil {
		return nil, err
	}

	pp.Println("@ops", opts)

	records := []models.ShortLedgerRecord{}

	rows, err := b.db.GetSession().SQL().Query(fmt.Sprintf(`
SELECT
	a.id as account_id,
	a.name as account_name,
	a.acc_type,
	SUM( tl.debit_amount) as  total_debit,
	SUM(tl.credit_amount) as total_credit
FROM z_%d_TransactionLines tl
INNER JOIN z_%d_Accounts a ON tl.account_id = a.id
INNER JOIN z_%d_Transactions t ON tl.txn_id = t.id
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
