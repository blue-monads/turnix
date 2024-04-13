package books

import "time"

type Account struct {
	ID          int        `json:"id" db:"id,omitempty"`
	Name        string     `json:"name" db:"name"`
	Info        string     `json:"info" db:"info"`
	AccType     string     `json:"acc_type" db:"acc_type"`
	TotalDebit  int        `json:"total_debit" db:"total_debit"`
	TotalCredit int        `json:"total_credit" db:"total_credit"`
	ContactID   int        `json:"contact_id" db:"contact_id"`
	CreatedAt   *time.Time `json:"created_at" db:"created_at,omitempty"`
	UpdatedAt   *time.Time `json:"updated_at" db:"updated_at,omitempty"`
	IsDeleted   bool       `json:"is_deleted" db:"is_deleted,omitempty"`
}

type Transaction struct {
	ID              int        `json:"id" db:"id,omitempty"`
	Title           string     `json:"title" db:"title"`
	Notes           string     `json:"notes" db:"notes"`
	AccountID       int        `json:"account_id" db:"account_id,omitempty"`
	TxnGroupID      int        `json:"txn_group_id" db:"txn_group_id,omitempty"`
	DebitAmount     int        `json:"debit_amount" db:"debit_amount,omitempty"`
	CreditAmount    int        `json:"credit_amount" db:"credit_amount,omitempty"`
	LinkedSalesID   int        `json:"linked_sales_id" db:"linked_sales_id,omitempty"`
	LinkedInvoiceID int        `json:"linked_invoice_id" db:"linked_invoice_id,omitempty"`
	ReferenceID     string     `json:"reference_id" db:"reference_id,omitempty"`
	Attachments     string     `json:"attachments" db:"attachments,omitempty"`
	CreatedBy       int        `json:"created_by" db:"created_by"`
	UpdatedBy       int        `json:"updated_by" db:"updated_by"`
	CreatedAt       *time.Time `json:"created_at" db:"created_at,omitempty"`
	UpdatedAt       *time.Time `json:"updated_at" db:"updated_at,omitempty"`
	IsDeleted       bool       `json:"is_deleted" db:"is_deleted,omitempty"`
}

type ReportTemplate struct {
	ID            int    `json:"id" db:"id"`
	Name          string `json:"name" db:"name"`
	InputOptions  string `json:"input_options" db:"input_options"`
	OutputOptions string `json:"output_options" db:"output_options"`
	QueryTemplate string `json:"query_template" db:"query_template"`
	FilterScript  string `json:"filter_script" db:"filter_script"`
}

type Report struct {
	ID         int        `json:"id" db:"id,omitempty"`
	TemplateID int        `json:"template_id" db:"template_id"`
	Result     string     `json:"result" db:"result"`
	CreatedBy  int        `json:"created_by" db:"created_by"`
	UpdatedBy  int        `json:"updated_by" db:"updated_by"`
	CreatedAt  *time.Time `json:"created_at" db:"created_at,omitempty"`
	UpdatedAt  *time.Time `json:"updated_at" db:"updated_at,omitempty"`
	IsDeleted  bool       `json:"is_deleted" db:"is_deleted,omitempty"`
}
