package books

import (
	"strings"

	"github.com/bornjre/trunis/backend/xtypes/services/xdatabase"
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

	return b.txnTable(pid).Find(db.Cond{"id": id}).Update(data)
}

func (b *BookModule) dbOpGetTxn(pid, uid, aid int64) (*Transaction, error) {

	err := b.userHasScope(pid, uid, "read")
	if err != nil {
		return nil, err
	}

	data := &Transaction{}
	table := b.txnTable(pid)

	err = table.Find(db.Cond{"id": aid}).One(data)
	if err != nil {
		return nil, err
	}

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

	lineTable := b.txnLineTable(pid)

	err = lineTable.Find(db.Cond{
		"txn_id >":   offset,
		"account_id": accId,
	}).Limit(200).All(&record.Lines)
	if err != nil {
		return nil, err
	}

	ids := make([]int64, 0, len(record.Lines))

	for _, line := range record.Lines {
		ids = append(ids, line.TxnID)
	}

	txtable := b.txnTable(pid)

	err = txtable.Find(db.Cond{"id IN": ids}).All(&record.Transactions)
	if err != nil {
		return nil, err
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

	table := b.txnLineTable(pid)

	r, err := table.Insert(data)
	if err != nil {
		return 0, err
	}

	return r.ID().(int64), nil
}

func (b *BookModule) dbOpUpdateTxnLine(pid, uid, id int64, data map[string]any) error {
	err := b.userHasScope(pid, uid, "write")
	if err != nil {
		return err
	}

	return b.txnLineTable(pid).Find(db.Cond{"id": id}).Update(data)
}

// utils

func (b *BookModule) txnTable(pid int64) db.Collection {
	return b.db.Table("Transactions")
}

func (b *BookModule) accountsTable(pid int64) db.Collection {
	return b.db.Table("Accounts")
}

func (b *BookModule) txnLineTable(pid int64) db.Collection {
	return b.db.Table("TransactionLines")
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
