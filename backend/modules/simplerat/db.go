package simplerat

import (
	"fmt"
	"strings"

	"github.com/blue-monads/turnix/backend/services/database"
	"github.com/upper/db/v4"
)

func (b *ECPServer) userHasScope(pid, uid int64, scope string) error {
	dbscope, err := b.db.GetProjectUserScope(uid, pid)
	if err != nil {
		return err
	}

	if dbscope != database.ScopeOwner && strings.Contains(dbscope, scope) {
		return database.ErrUserNoScope
	}

	return nil
}

func (b *ECPServer) deviceTable(pid int64) db.Collection {
	return b.db.Table(deviceTableName(pid))
}

func deviceTableName(pid int64) string {
	return fmt.Sprintf("z_%d_Devices", pid)
}

func (b *ECPServer) dbAddDevice(pid, uid int64, data *Device) (int64, error) {
	err := b.userHasScope(pid, uid, "write")
	if err != nil {
		return 0, err
	}

	table := b.deviceTable(pid)

	r, err := table.Insert(data)
	if err != nil {
		return 0, err
	}

	return r.ID().(int64), nil
}

func (b *ECPServer) dbUpdateDevice(pid, uid, id int64, data map[string]any) error {
	err := b.userHasScope(pid, uid, "write")
	if err != nil {
		return err
	}

	return b.deviceTable(pid).Find(db.Cond{"id": id}).Update(data)
}

func (b *ECPServer) dbGetDevice(pid, uid, id int64) (*Device, error) {
	err := b.userHasScope(pid, uid, "read")
	if err != nil {
		return nil, err
	}

	data := &Device{}
	table := b.deviceTable(pid)

	err = table.Find(db.Cond{"id": id}).One(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (b *ECPServer) dbListDevice(pid, uid, offset int64) ([]Device, error) {
	err := b.userHasScope(pid, uid, "read")
	if err != nil {
		return nil, err
	}

	datas := make([]Device, 0)
	table := b.deviceTable(pid)

	err = table.Find(db.Cond{"id >": offset}).Limit(100).All(&datas)
	if err != nil {
		return nil, err
	}

	return datas, nil
}

func (b *ECPServer) dbDeleteDevice(pid, uid, id int64) error {
	err := b.userHasScope(pid, uid, "write")
	if err != nil {
		return err
	}

	return b.deviceTable(pid).Find(db.Cond{"id": id}).Update(db.Cond{"status": "deleted"})
}
