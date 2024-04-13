package books

import (
	"strings"

	"github.com/bornjre/trunis/backend/xtypes/services/xdatabase"
	"github.com/upper/db/v4"
)

// account

func (b *BookModule) addAccount(pid, uid int64, data *Account) (int64, error) {

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

func (b *BookModule) updateAccount(pid, uid, id int64, data map[string]any) error {
	err := b.userHasScope(pid, uid, "write")
	if err != nil {
		return err

	}

	return b.accountsTable(pid).Find(db.Cond{"id": id}).Update(data)
}

func (b *BookModule) getAccount(pid, uid, aid int64) (*Account, error) {

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

func (b *BookModule) listAccount(pid, uid int64) ([]Account, error) {

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

// transactions

func (b *BookModule) addTxn(pid, uid int64, data *Transaction) (int64, error) {

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

func (b *BookModule) updateTxn(pid, uid, id int64, data map[string]any) error {
	err := b.userHasScope(pid, uid, "write")
	if err != nil {
		return err

	}

	return b.txnTable(pid).Find(db.Cond{"id": id}).Update(data)
}

func (b *BookModule) getTxn(pid, uid, aid int64) (*Transaction, error) {

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

func (b *BookModule) listTxn(pid, uid int64) ([]Transaction, error) {

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

// utils

func (b *BookModule) txnTable(pid int64) db.Collection {
	return b.db.Table("Accounts")
}

func (b *BookModule) accountsTable(pid int64) db.Collection {
	return b.db.Table("Transactions")
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
