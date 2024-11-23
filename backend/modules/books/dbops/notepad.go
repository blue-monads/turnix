package dbops

import (
	"github.com/blue-monads/turnix/backend/modules/books/models"
	"github.com/upper/db/v4"
)

func (b *DbOps) NotepadAdd(pid, uid int64, data *models.Notepad) (int64, error) {
	err := b.userHasScope(pid, uid, "write")
	if err != nil {
		return 0, err
	}

	table := b.notepadTable(pid)

	r, err := table.Insert(data)
	if err != nil {
		return 0, err
	}

	return r.ID().(int64), nil
}

func (b *DbOps) NotepadUpdate(pid, uid, id int64, data map[string]any) error {
	err := b.userHasScope(pid, uid, "write")
	if err != nil {
		return err
	}

	return b.notepadTable(pid).Find(db.Cond{"id": id}).Update(data)
}

func (b *DbOps) NotepadGet(pid, uid, id int64) (*models.Notepad, error) {
	err := b.userHasScope(pid, uid, "read")
	if err != nil {
		return nil, err
	}

	data := &models.Notepad{}
	table := b.notepadTable(pid)

	err = table.Find(db.Cond{"id": id}).One(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (b *DbOps) NotepadList(pid, uid, offset int64) ([]models.Notepad, error) {
	err := b.userHasScope(pid, uid, "read")
	if err != nil {
		return nil, err
	}

	datas := make([]models.Notepad, 0)
	table := b.notepadTable(pid)

	err = table.Find(db.Cond{"id >": offset}).Limit(100).All(&datas)
	if err != nil {
		return nil, err
	}

	return datas, nil
}

func (b *DbOps) NotepadDelete(pid, uid, id int64) error {
	err := b.userHasScope(pid, uid, "write")
	if err != nil {
		return err
	}

	return b.notepadTable(pid).Find(db.Cond{"id": id}).Delete()
}
