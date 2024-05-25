package books

import "time"

type Account struct {
	ID          int64      `json:"id" db:"id,omitempty"`
	Name        string     `json:"name" db:"name"`
	Info        string     `json:"info" db:"info"`
	AccType     string     `json:"acc_type" db:"acc_type"`
	TotalDebit  int64      `json:"total_debit" db:"total_debit"`
	TotalCredit int64      `json:"total_credit" db:"total_credit"`
	ContactID   int64      `json:"contact_id" db:"contact_id"`
	CreatedAt   *time.Time `json:"created_at" db:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at" db:"updated_at,omitempty"`
	IsDeleted   bool       `json:"is_deleted" db:"is_deleted,omitempty"`
}

type Transaction struct {
	ID              int        `json:"id" db:"id,omitempty"`
	Title           string     `json:"title" db:"title"`
	Notes           string     `json:"notes" db:"notes"`
	LinkedSalesID   int64      `json:"linked_sales_id" db:"linked_sales_id,omitempty"`
	LinkedInvoiceID int64      `json:"linked_invoice_id" db:"linked_invoice_id,omitempty"`
	ReferenceID     string     `json:"reference_id" db:"reference_id,omitempty"`
	Attachments     string     `json:"attachments" db:"attachments,omitempty"`
	CreatedBy       int64      `json:"created_by" db:"created_by"`
	UpdatedBy       int64      `json:"updated_by" db:"updated_by"`
	CreatedAt       *time.Time `json:"created_at" db:"created_at,omitempty"`
	UpdatedAt       *time.Time `json:"updated_at" db:"updated_at,omitempty"`
	IsDeleted       bool       `json:"is_deleted" db:"is_deleted,omitempty"`
}

type TransactionLine struct {
	ID           int64      `json:"id" db:"id,omitempty"`
	AccountID    int64      `json:"account_id" db:"account_id"`
	TxnID        int64      `json:"txn_id" db:"txn_id"`
	DebitAmount  float64    `json:"debit_amount" db:"debit_amount"`
	CreditAmount float64    `json:"credit_amount" db:"credit_amount"`
	CreatedBy    int64      `json:"created_by" db:"created_by"`
	UpdatedBy    int64      `json:"updated_by" db:"updated_by"`
	CreatedAt    *time.Time `json:"created_at" db:"created_at,omitempty"`
	UpdatedAt    *time.Time `json:"updated_at" db:"updated_at,omitempty"`
}

type ReportTemplate struct {
	ID            int64  `json:"id" db:"id,omitempty"`
	Name          string `json:"name" db:"name"`
	InputOptions  string `json:"input_options" db:"input_options"`
	OutputOptions string `json:"output_options" db:"output_options"`
	QueryTemplate string `json:"query_template" db:"query_template"`
	FilterScript  string `json:"filter_script" db:"filter_script"`
}

type Report struct {
	ID         int64      `json:"id" db:"id,omitempty"`
	TemplateID int64      `json:"template_id" db:"template_id"`
	Result     string     `json:"result" db:"result"`
	CreatedBy  int64      `json:"created_by" db:"created_by"`
	UpdatedBy  int64      `json:"updated_by" db:"updated_by"`
	CreatedAt  *time.Time `json:"created_at" db:"created_at,omitempty"`
	UpdatedAt  *time.Time `json:"updated_at" db:"updated_at,omitempty"`
	IsDeleted  bool       `json:"is_deleted" db:"is_deleted,omitempty"`
}

type TxnRecord struct {
	Title           string            `json:"title,omitempty"`
	Notes           string            `json:"notes,omitempty"`
	DebitAccountID  int64             `json:"debit_account_id,omitempty"`
	CreditAccountID int64             `json:"credit_account_id,omitempty"`
	DebitAmount     float64           `json:"debit_amount,omitempty"`
	CreditAmount    float64           `json:"credit_amount,omitempty"`
	LinkedSalesID   int64             `json:"linked_sales_id,omitempty"`
	LinkedInvoiceID int64             `json:"linked_invoice_id,omitempty"`
	ReferenceID     string            `json:"reference_id,omitempty"`
	Attachments     map[string]string `json:"attachments,omitempty"`
}

type TxnResult struct {
	TxnId        int64 `json:"txn_id,omitempty"`
	DebitLineId  int64 `json:"debit_line_id,omitempty"`
	CreditLineId int64 `json:"credit_line_id,omitempty"`
}

/*

SELECT *
FROM TransactionLines t1, TransactionLines t2
INNER JOIN Transactions t on t.id = t1.txn_id
WHERE t1.account_id = 2 and t1.txn_id = t2.txn_id and t2.account_id <> 2;



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
    TransactionLines t1, TransactionLines t2
INNER JOIN
    Transactions t ON t.id = t1.txn_id
WHERE
    t1.account_id = 2 AND t1.txn_id = t2.txn_id AND t2.account_id <> 2 and t.is_deleted = FALSE;


*/

type TransactionResult struct {
	Id              int        `db:"id" json:"id"`
	Title           string     `db:"title" json:"title"`
	Notes           string     `db:"notes" json:"notes"`
	LinkedSalesID   int        `db:"linked_sales_id" json:"linked_sales_id"`
	LinkedInvoiceID int        `db:"linked_invoice_id" json:"linked_invoice_id"`
	ReferenceID     int        `db:"reference_id" json:"reference_id"`
	Attachments     string     `db:"attachments" json:"attachments"`
	TxnCreatedBy    int        `db:"txn_created_by" json:"txn_created_by"`
	TxnUpdatedBy    int        `db:"txn_updated_by" json:"txn_updated_by"`
	TxnCreatedAt    *time.Time `db:"txn_created_at" json:"txn_created_at"`
	TxnUpdatedAt    *time.Time `db:"txn_updated_at" json:"txn_updated_at"`
	IsDeleted       bool       `db:"is_deleted" json:"is_deleted"`

	FirstID           int        `db:"first_id" json:"first_id"`
	FirstAccountID    int        `db:"first_account_id" json:"first_account_id"`
	FirstDebitAmount  float64    `db:"first_debit_amount" json:"first_debit_amount"`
	FirstCreditAmount float64    `db:"first_credit_amount" json:"first_credit_amount"`
	FirstCreatedBy    int        `db:"first_created_by" json:"first_created_by"`
	FirstUpdatedBy    int        `db:"first_updated_by" json:"first_updated_by"`
	FirstCreatedAt    *time.Time `db:"first_created_at" json:"first_created_at"`
	FirstUpdatedAt    *time.Time `db:"first_updated_at" json:"first_updated_at"`

	SecondID           int        `db:"second_id" json:"second_id"`
	SecondAccountID    int        `db:"second_account_id" json:"second_account_id"`
	SecondDebitAmount  float64    `db:"second_debit_amount" json:"second_debit_amount"`
	SecondCreditAmount float64    `db:"second_credit_amount" json:"second_credit_amount"`
	SecondCreatedBy    int        `db:"second_created_by" json:"second_created_by"`
	SecondUpdatedBy    int        `db:"second_updated_by" json:"second_updated_by"`
	SecondCreatedAt    *time.Time `db:"second_created_at" json:"second_created_at"`
	SecondUpdatedAt    *time.Time `db:"second_updated_at" json:"second_updated_at"`
}
