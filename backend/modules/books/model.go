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
	ID              int64      `json:"id" db:"id,omitempty"`
	Title           string     `json:"title" db:"title"`
	Notes           string     `json:"notes" db:"notes"`
	LinkedSalesID   int64      `json:"linked_sales_id" db:"linked_sales_id,omitempty"`
	LinkedInvoiceID int64      `json:"linked_invoice_id" db:"linked_invoice_id,omitempty"`
	ReferenceID     string     `json:"reference_id" db:"reference_id,omitempty"`
	Attachments     string     `json:"attachments" db:"attachments,omitempty"`
	CreatedBy       int64      `json:"created_by" db:"created_by"`
	UpdatedBy       int64      `json:"updated_by" db:"updated_by"`
	TxnDate         *time.Time `json:"txn_date" db:"txn_date,omitempty"`
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

type TransactionWithLine struct {
	Txn        Transaction      `db:"txn" json:"txn"`
	FirstLine  *TransactionLine `json:"first_line"`
	SecondLine *TransactionLine `json:"second_line"`
}

type TransactionResult struct {
	Id              int64      `db:"id" json:"id"`
	Title           string     `db:"title" json:"title"`
	Notes           string     `db:"notes" json:"notes"`
	LinkedSalesID   int64      `db:"linked_sales_id" json:"linked_sales_id"`
	LinkedInvoiceID int64      `db:"linked_invoice_id" json:"linked_invoice_id"`
	ReferenceID     string     `db:"reference_id" json:"reference_id"`
	Attachments     string     `db:"attachments" json:"attachments"`
	TxnCreatedBy    int64      `db:"txn_created_by" json:"txn_created_by"`
	TxnUpdatedBy    int64      `db:"txn_updated_by" json:"txn_updated_by"`
	TxnCreatedAt    *time.Time `db:"txn_created_at" json:"txn_created_at"`
	TxnUpdatedAt    *time.Time `db:"txn_updated_at" json:"txn_updated_at"`
	IsDeleted       bool       `db:"is_deleted" json:"is_deleted"`

	FirstID           int64      `db:"first_id" json:"first_id"`
	FirstAccountID    int64      `db:"first_account_id" json:"first_account_id"`
	FirstDebitAmount  float64    `db:"first_debit_amount" json:"first_debit_amount"`
	FirstCreditAmount float64    `db:"first_credit_amount" json:"first_credit_amount"`
	FirstCreatedBy    int64      `db:"first_created_by" json:"first_created_by"`
	FirstUpdatedBy    int64      `db:"first_updated_by" json:"first_updated_by"`
	FirstCreatedAt    *time.Time `db:"first_created_at" json:"first_created_at"`
	FirstUpdatedAt    *time.Time `db:"first_updated_at" json:"first_updated_at"`

	SecondID           int64      `db:"second_id" json:"second_id"`
	SecondAccountID    int64      `db:"second_account_id" json:"second_account_id"`
	SecondDebitAmount  float64    `db:"second_debit_amount" json:"second_debit_amount"`
	SecondCreditAmount float64    `db:"second_credit_amount" json:"second_credit_amount"`
	SecondCreatedBy    int64      `db:"second_created_by" json:"second_created_by"`
	SecondUpdatedBy    int64      `db:"second_updated_by" json:"second_updated_by"`
	SecondCreatedAt    *time.Time `db:"second_created_at" json:"second_created_at"`
	SecondUpdatedAt    *time.Time `db:"second_updated_at" json:"second_updated_at"`
}

const (
	ReportTypeLongLedger  = "long_ledger"
	ReportTypeShortLedger = "short_ledger"
	ReportTypeCustom      = "custom"
)

type LongLedgerRecord struct {
	TxnId        int64  `json:"txn_id" db:"txn_id"`
	Title        string `json:"title" db:"title"`
	AccountName  string `json:"account_name" db:"account_name"`
	AccountType  string `json:"acc_type" db:"acc_type"`
	DebitAmount  int64  `json:"debit_amount" db:"debit_amount"`
	CreditAmount int64  `json:"credit_amount" db:"credit_amount"`
	TotalDebit   int64  `json:"total_debit" db:"total_debit"`
	TotalCredit  int64  `json:"total_credit" db:"total_credit"`
	AccountID    int64  `json:"account_id" db:"account_id"`
}

type ShortLedgerRecord struct {
	AccountID   int64  `json:"account_id" db:"account_id"`
	AccountName string `json:"account_name" db:"account_name"`
	TotalDebit  int64  `json:"total_debit" db:"total_debit"`
	TotalCredit int64  `json:"total_credit" db:"total_credit"`
	AccountType string `json:"acc_type" db:"acc_type"`
}
