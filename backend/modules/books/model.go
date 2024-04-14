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
