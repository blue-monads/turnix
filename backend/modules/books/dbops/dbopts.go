package dbops

import (
	"fmt"
	"strings"

	"github.com/blue-monads/turnix/backend/services/database"
	"github.com/upper/db/v4"
)

type DbOps struct {
	db *database.DB
}

func New(db *database.DB) DbOps {
	return DbOps{
		db: db,
	}
}

// utils

func (b *DbOps) txnTable(pid int64) db.Collection {
	return b.db.Table(txnTableName(pid))
}

func txnTableName(pid int64) string {
	return fmt.Sprintf("z_%d_Transactions", pid)
}

func (b *DbOps) accountsTable(pid int64) db.Collection {
	return b.db.Table(accountsTableName(pid))
}

func accountsTableName(pid int64) string {
	return fmt.Sprintf("z_%d_Accounts", pid)
}

func (b *DbOps) txnLineTable(pid int64) db.Collection {
	return b.db.Table(txnLineTableName(pid))
}

func txnLineTableName(pid int64) string {
	return fmt.Sprintf("z_%d_TransactionLines", pid)
}

func (b *DbOps) catagoryTable(pid int64) db.Collection {
	return b.db.Table(catagoryTableName(pid))
}

func catagoryTableName(pid int64) string {
	return fmt.Sprintf("z_%d_Catagories", pid)
}

func (b *DbOps) productTable(pid int64) db.Collection {
	return b.db.Table(productTableName(pid))
}

func productTableName(pid int64) string {
	return fmt.Sprintf("z_%d_Products", pid)
}

func (b *DbOps) productStockInTable(pid int64) db.Collection {
	return b.db.Table(productStockInTableName(pid))
}

func productStockInTableName(pid int64) string {
	return fmt.Sprintf("z_%d_ProductStockIn", pid)
}

func (b *DbOps) productStockInLineTable(pid int64) db.Collection {
	return b.db.Table(productStockInLineTableName(pid))
}

func productStockInLineTableName(pid int64) string {
	return fmt.Sprintf("z_%d_ProductStockInLines", pid)
}

func (b *DbOps) contactTable(pid int64) db.Collection {
	return b.db.Table(contactTableName(pid))
}

func contactTableName(pid int64) string {
	return fmt.Sprintf("z_%d_Contacts", pid)
}

func (b *DbOps) taxTable(pid int64) db.Collection {
	return b.db.Table(taxTableName(pid))
}

func taxTableName(pid int64) string {
	return fmt.Sprintf("z_%d_Tax", pid)
}

func (b *DbOps) productTaxTable(pid int64) db.Collection {
	return b.db.Table(productTaxTableName(pid))
}

func productTaxTableName(pid int64) string {
	return fmt.Sprintf("z_%d_ProductTaxes", pid)
}

func (b *DbOps) notepadTable(pid int64) db.Collection {
	return b.db.Table(notepadTableName(pid))
}

func (b *DbOps) estimateTable(pid int64) db.Collection {
	return b.db.Table(estimateTableName(pid))
}

func estimateTableName(pid int64) string {
	return fmt.Sprintf("z_%d_Estimates", pid)
}

func (b *DbOps) estimateLineTable(pid int64) db.Collection {
	return b.db.Table(estimateLineTableName(pid))
}

func estimateLineTableName(pid int64) string {
	return fmt.Sprintf("z_%d_EstimateLines", pid)
}

func notepadTableName(pid int64) string {
	return fmt.Sprintf("z_%d_Notepads", pid)
}

func (b *DbOps) savedReportTable(pid int64) db.Collection {
	return b.db.Table(savedReportTableName(pid))
}

func savedReportTableName(pid int64) string {
	return fmt.Sprintf("z_%d_SavedReports", pid)
}

func (b *DbOps) reportTemplateTable(pid int64) db.Collection {
	return b.db.Table(reportTemplateTableName(pid))
}

func reportTemplateTableName(pid int64) string {
	return fmt.Sprintf("z_%d_ReportTemplates", pid)
}

func (b *DbOps) userHasScope(pid, uid int64, scope string) error {
	dbscope, err := b.db.GetProjectUserScope(uid, pid)
	if err != nil {
		return err
	}

	if dbscope != database.ScopeOwner && strings.Contains(dbscope, scope) {
		return database.ErrUserNoScope
	}

	return nil
}
