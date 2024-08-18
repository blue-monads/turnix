package database

import (
	"fmt"
	"os"
	"time"

	"github.com/k0kubun/pp"
	"github.com/upper/db/v4"
)

type File struct {
	ID        int64      `db:"id,omitempty" json:"id"`
	Name      string     `db:"name" json:"name"`
	FType     string     `db:"ftype" json:"ftype"`
	Path      string     `db:"path" json:"path"`
	Size      int64      `db:"size" json:"size"`
	Mime      string     `db:"mime" json:"mime"`
	Hash      string     `db:"hash" json:"hash"`
	IsFolder  bool       `db:"is_folder" json:"is_folder"`
	External  bool       `db:"external" json:"external"`
	IsPublic  bool       `db:"is_public" json:"is_public"`
	OwnerUser int64      `db:"owner_user_id" json:"owner_user_id"`
	OwnerProj int64      `db:"owner_project_id" json:"owner_project_id"`
	CreatedAt *time.Time `db:"created_at" json:"created_at"`
}

func (d *DB) AddFolder(pid, uid int64, ftype, path, name string) (int64, error) {
	table := d.filesTable()

	t := time.Now()

	file := &File{
		Name:      name,
		Path:      path,
		OwnerUser: uid,
		OwnerProj: pid,
		FType:     ftype,
		IsPublic:  false,
		Size:      0,
		CreatedAt: &t,
		IsFolder:  true,
	}

	rid, err := table.Insert(file)
	if err != nil {
		return 0, err
	}

	id := rid.ID().(int64)

	return id, nil
}

func (d *DB) AddFile(file *File, data []byte) (id int64, err error) {

	pp.Println("@file_add", file)

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
		err = table.Find(db.Cond{"id": id}).Update(db.Cond{"blob": data})
		return id, err
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
		"path":             path,
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

	pp.Println("@list_files_by_user", uid, path)

	cond := db.Cond{
		"ftype":         "user",
		"owner_user_id": uid,
		"path":          path,
	}

	files := make([]File, 0)
	err := table.Find(cond).All(&files)
	if err != nil {
		return nil, err
	}

	return files, nil
}

func (d *DB) GetFileMeta(id int64) (*File, error) {
	table := d.filesTable()

	file := File{}
	err := table.Find(db.Cond{"id": id}).One(&file)
	if err != nil {
		return nil, err
	}

	return &file, nil
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
