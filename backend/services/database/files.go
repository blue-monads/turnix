package database

import (
	"fmt"
	"os"

	"github.com/upper/db/v4"
)

type File struct {
	ID        int64  `db:"id"`
	Name      string `db:"name"`
	FType     string `db:"ftype"`
	Path      string `db:"path"`
	Size      int64  `db:"size"`
	Mime      string `db:"mime"`
	Hash      string `db:"hash"`
	External  bool   `db:"external"`
	IsPublic  bool   `db:"is_public"`
	OwnerUser int64  `db:"owner_user_id"`
	OwnerProj int64  `db:"owner_project_id"`
	CreatedAt string `db:"created_at"`
}

func (d *DB) AddFile(file *File, data []byte) (id int64, err error) {
	table := d.filesTable()
	rid, err := table.Insert(file)
	if err != nil {
		return 0, err
	}

	id = rid.ID().(int64)

	defer func() {
		if err != nil {
			table.Find(db.Cond{"id": id}).Delete()
		}
	}()

	if !d.externalFileMode {
		_, err = d.sess.SQL().
			InsertInto("Files").
			Columns("data").
			Values(data).
			Exec()

		return
	}

	pidOrUid := file.OwnerUser
	if file.FType == "project" {
		pidOrUid = file.OwnerProj
	}

	err = os.WriteFile(fmt.Sprintf("files/%s/%d/%s", file.FType, pidOrUid, file.Path), data, 0644)

	return
}

func (d *DB) ListFilesByProject(pid int64, path string) ([]File, error) {
	table := d.filesTable()

	cond := db.Cond{
		"ftype":            "project",
		"owner_project_id": pid,
		"path LIKE":        fmt.Sprintf("%s%%", path),
	}

	files := make([]File, 0)
	err := table.Find(cond).All(&files)
	if err != nil {
		return nil, err
	}

	return files, nil

}

func (d *DB) ListFilesByUser(uid int64, path string) ([]File, error) {
	table := d.filesTable()

	cond := db.Cond{
		"ftype":         "user",
		"owner_user_id": uid,
		"path LIKE":     fmt.Sprintf("%s%%", path),
	}

	files := make([]File, 0)
	err := table.Find(cond).All(&files)
	if err != nil {
		return nil, err
	}

	return files, nil
}

func (d *DB) GetFileBlob(id int64) ([]byte, error) {
	table := d.filesTable()

	file := File{}
	err := table.Find(db.Cond{"id": id}).One(&file)
	if err != nil {
		return nil, err
	}

	if !file.External {
		data := make([]byte, file.Size)

		row, err := d.sess.SQL().Select("blob").From("Files").Where(db.Cond{"id": id}).QueryRow()
		if err != nil {
			return nil, err
		}

		err = row.Scan(&data)
		if err != nil {
			return nil, err
		}

		return data, nil
	}

	pidOrUid := file.OwnerUser
	if file.FType == "project" {
		pidOrUid = file.OwnerProj
	}

	return os.ReadFile(fmt.Sprintf("files/%s/%d/%s", file.FType, pidOrUid, file.Path))

}

func (d *DB) DeleteFile(id int64) error {
	table := d.filesTable()

	file := File{}

	record := table.Find(db.Cond{"id": id})

	err := record.One(&file)
	if err != nil {
		return err
	}

	if file.External {

		pidOrUid := file.OwnerUser
		if file.FType == "project" {
			pidOrUid = file.OwnerProj
		}

		err = os.Remove(fmt.Sprintf("files/%s/%d/%s", file.FType, pidOrUid, file.Path))
		if err != nil {
			return err
		}

	}

	return record.Delete()

}

func (d *DB) filesTable() db.Collection {
	return d.Table("Files")
}
