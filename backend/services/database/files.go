package database

import (
	"bytes"
	"database/sql"
	"fmt"
	"io"
	"os"
	"time"

	"github.com/jaevor/go-nanoid"
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
	StoreType int64      `db:"storeType" json:"storeType"`
	OwnerUser int64      `db:"owner_user_id" json:"owner_user_id"`
	OwnerProj int64      `db:"owner_project_id" json:"owner_project_id"`
	CreatedAt *time.Time `db:"created_at" json:"created_at"`
}

type FilePartedBlob struct {
	Id     int64 `db:"id" json:"id"`
	FileId int64 `db:"file_id" json:"file_id"`
	Size   int64 `db:"size" json:"size"`
	PartId int64 `db:"part_id" json:"part_id"`
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

func (d *DB) AddFileStreaming(file *File, stream io.Reader) (id int64, err error) {

	pp.Println("@add_file_streaming/1", file.FType, file.Path)

	filetable := d.filesTable()
	partstable := d.filePartedBlobsTable()

	exists, _ := filetable.Find(db.Cond{
		"name":             file.Name,
		"path":             file.Path,
		"owner_user_id":    file.OwnerUser,
		"owner_project_id": file.OwnerProj,
	}).Exists()

	pp.Println("@add_file_streaming/2", exists)

	if exists {
		pp.Println("@add_file_streaming/3")

		return 0, fmt.Errorf("file already exists")
	}

	if d.externalFileMode {
		file.StoreType = 1
	} else if file.Size > d.minFileMultiPartSize {
		file.StoreType = 2
	}

	pp.Println("@add_file_streaming/3", file.StoreType)

	rid, err := filetable.Insert(file)
	if err != nil {
		return 0, err
	}

	id = rid.ID().(int64)

	pp.Println("@add_file_streaming/4")

	defer func() {
		if err != nil {
			filetable.Find(db.Cond{"id": id}).Delete()
			partstable.Find(db.Cond{"file_id": id}).Delete()
		}
	}()

	pp.Println("@add_file_streaming/5")

	pidOrUid := file.OwnerUser
	if file.FType == "project" {
		pidOrUid = file.OwnerProj
	}

	pp.Println("@add_file_streaming/6", pidOrUid)

	if d.externalFileMode {

		pp.Println("@add_file_streaming/7")

		file, err := os.Create(fmt.Sprintf("files/%s/%d/%s", file.FType, pidOrUid, file.Path))
		if err != nil {
			return 0, err
		}

		pp.Println("@add_file_streaming/8")

		defer file.Close()

		_, err = io.Copy(file, stream)
		if err != nil {
			return 0, err
		}

		pp.Println("@add_file_streaming/9")

		file.Sync()

		pp.Println("@add_file_streaming/10")

		return id, nil
	}

	driver := d.sess.Driver().(*sql.DB)

	if file.StoreType == 0 {
		pp.Println("@add_file_streaming/11")

		data, err := io.ReadAll(stream)
		if err != nil {
			return 0, err
		}

		pp.Println("@add_file_streaming/12")

		_, err = driver.Exec("UPDATE Files SET blob = ? WHERE id = ?", data, id)
		return id, err

	} else if file.StoreType == 2 {

		pp.Println("@add_file_streaming/13")

		partId := 0

		buf := make([]byte, d.minFileMultiPartSize)

		for {

			pp.Println("@add_file_streaming/14", partId)

			n, err := stream.Read(buf)
			if err != nil && err != io.EOF {
				return 0, err
			}

			if n == 0 {
				break
			}

			pp.Println("@add_file_streaming/15", n)

			_, err = driver.Exec("INSERT INTO FilePartedBlobs (file_id, size, part_id, blob) VALUES (?, ?, ?, ?)", id, n, partId, buf[:n])
			if err != nil {
				return 0, err
			}

			pp.Println("@add_file_streaming/16")

			partId++

		}

		pp.Println("@add_file_streaming/17")

	}

	pp.Println("@add_file_streaming/18")

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
	buf := bytes.Buffer{}

	err := d.GetFileBlobStreaming(id, &buf)
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil

}

func (d *DB) GetFileBlobStreaming(id int64, w io.Writer) error {

	pp.Println("@get_file_blob_streaming/1", id)

	table := d.filesTable()

	file := File{}
	err := table.Find(db.Cond{"id": id}).One(&file)
	if err != nil {
		return err
	}

	pp.Println("@get_file_blob_streaming/2", file.FType, file.Path)

	if file.StoreType == 0 {
		data := make([]byte, file.Size)

		pp.Println("@get_file_blob_streaming/3")

		row, err := d.sess.SQL().Select("blob").From("Files").Where(db.Cond{"id": id}).QueryRow()
		if err != nil {
			pp.Println("@get_file_blob_streaming/4", err)
			return err
		}

		pp.Println("@get_file_blob_streaming/5")

		err = row.Scan(&data)
		if err != nil {
			pp.Println("@get_file_blob_streaming/6", err)
			return err
		}

		pp.Println("@get_file_blob_streaming/7")

		written := int64(0)

		for written < file.Size || written == 0 {

			pp.Println("@get_file_blob_streaming/8", written, file.Size)

			n, err := w.Write(data[written:])
			if err != nil {
				return err
			}
			written += int64(n)

		}

		return err
	} else if file.StoreType == 1 {

		pp.Println("@get_file_blob_streaming/9")

		pidOrUid := file.OwnerUser
		if file.FType == "project" {
			pidOrUid = file.OwnerProj
		}

		pp.Println("@get_file_blob_streaming/10", pidOrUid)

		out, err := os.ReadFile(fmt.Sprintf("files/%s/%d/%s", file.FType, pidOrUid, file.Path))
		if err != nil {
			return err
		}

		pp.Println("@get_file_blob_streaming/11")

		_, err = w.Write(out)
		if err != nil {
			return err
		}

	} else if file.StoreType == 2 {
		parts := make([]FilePartedBlob, 0)

		pp.Println("@get_file_blob_streaming/12")

		err := d.filePartedBlobsTable().Find(db.Cond{"file_id": id}).
			Select("id", "size", "part_id").
			OrderBy("part_id").
			All(&parts)
		if err != nil {
			pp.Println("@get_file_blob_streaming/13", err.Error())
			return err
		}

		pp.Println("@get_file_blob_streaming/14", len(parts))

		for _, part := range parts {

			pp.Println("@get_file_blob_streaming/15")

			blob, err := d.sess.SQL().Select("blob").From("FilePartedBlobs").Where(db.Cond{"id": part.Id}).QueryRow()
			if err != nil {
				pp.Println("@get_file_blob_streaming/16", err.Error())
				return err
			}

			pp.Println("@get_file_blob_streaming/17")

			data := make([]byte, part.Size)
			err = blob.Scan(&data)
			if err != nil {
				pp.Println("@get_file_blob_streaming/18", err.Error())
				return err
			}

			pp.Println("@get_file_blob_streaming/19")

			written := int64(0)

			for written <= part.Size {
				pp.Println("@get_file_blob_streaming/20")

				_, err = w.Write(data[written:])
				if err != nil {
					return err
				}

				written += int64(len(data))
			}

			pp.Println("@get_file_blob_streaming/21")

		}

		pp.Println("@get_file_blob_streaming/22")

		return nil

	}

	pp.Println("@get_file_blob_streaming/23")

	return nil

}

func (d *DB) DeleteFile(id int64) error {
	table := d.filesTable()

	file := File{}

	record := table.Find(db.Cond{"id": id})

	err := record.One(&file)
	if err != nil {
		return err
	}

	if file.IsFolder {
		// fixme => delete all files in folder
		pp.Println("@delete_files_form_folder/1")
	}

	if file.StoreType == 1 {

		pidOrUid := file.OwnerUser
		if file.FType == "project" {
			pidOrUid = file.OwnerProj
		}

		err = os.Remove(fmt.Sprintf("files/%s/%d/%s", file.FType, pidOrUid, file.Path))
		if err != nil {
			return err
		}

	} else if file.StoreType == 2 {
		d.filePartedBlobsTable().Find(db.Cond{"file_id": id}).Delete()
	}

	return record.Delete()

}

type FileShare struct {
	ID        string     `json:"id" db:"id,omitempty"`
	FileID    int64      `json:"file_id" db:"file_id"`
	UserID    int64      `json:"user_id" db:"user_id"`
	CreatedAt *time.Time `json:"created_at" db:"created_at"`
}

var generator, _ = nanoid.CustomASCII("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789", 12)

func (d *DB) AddFileShare(fileId, userId int64) (int64, error) {
	table := d.fileSharesTable()

	t := &time.Time{}

	data := &FileShare{
		ID:        generator(),
		FileID:    fileId,
		UserID:    userId,
		CreatedAt: t,
	}

	rid, err := table.Insert(data)
	if err != nil {
		return 0, err
	}

	id := rid.ID().(int64)

	return id, nil
}

func (d *DB) ListFileShares(fileId int64) ([]FileShare, error) {
	table := d.fileSharesTable()

	shares := make([]FileShare, 0)

	err := table.Find(db.Cond{"file_id": fileId}).All(&shares)
	if err != nil {
		return nil, err
	}
	return shares, nil

}

func (d *DB) DeleteFileShare(userId int64, id string) error {
	table := d.fileSharesTable()

	return table.Find(db.Cond{"id": id, "user_id": userId}).Delete()
}

func (d *DB) GetSharedFile(id string) ([]byte, error) {
	table := d.fileSharesTable()
	file := &FileShare{}

	err := table.Find(db.Cond{"id": id}).One(file)
	if err != nil {
		return nil, err
	}

	return d.GetFileBlob(file.FileID)
}

func (d *DB) fileSharesTable() db.Collection {
	return d.Table("FileShares")
}

func (d *DB) filesTable() db.Collection {
	return d.Table("Files")
}

func (d *DB) filePartedBlobsTable() db.Collection {
	return d.Table("FilePartedBlobs")
}
