package dbops

import (
	"fmt"
	"strings"

	"github.com/bornjre/turnix/backend/services/database"
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
	return b.db.Table(fmt.Sprintf("z_%d_Transactions", pid))
}

func (b *DbOps) accountsTable(pid int64) db.Collection {
	return b.db.Table(fmt.Sprintf("z_%d_Accounts", pid))
}

func (b *DbOps) txnLineTable(pid int64) db.Collection {
	return b.db.Table(fmt.Sprintf("z_%d_TransactionLines", pid))
}

func (b *DbOps) catagoryTable(pid int64) db.Collection {
	return b.db.Table(fmt.Sprintf("z_%d_Catagories", pid))
}

func (b *DbOps) productTable(pid int64) db.Collection {
	return b.db.Table(fmt.Sprintf("z_%d_Products", pid))
}

func (b *DbOps) productStockInTable(pid int64) db.Collection {
	return b.db.Table(fmt.Sprintf("z_%d_ProductStockIn", pid))
}

func (b *DbOps) productStockInLineTable(pid int64) db.Collection {
	return b.db.Table(fmt.Sprintf("z_%d_ProductStockInLines", pid))
}

func (b *DbOps) contactTable(pid int64) db.Collection {
	return b.db.Table(fmt.Sprintf("z_%d_Contacts", pid))
}

func (b *DbOps) taxTable(pid int64) db.Collection {
	return b.db.Table(fmt.Sprintf("z_%d_Tax", pid))
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
