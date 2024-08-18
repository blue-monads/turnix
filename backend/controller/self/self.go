package self

import (
	"fmt"
	"time"

	"github.com/bornjre/turnix/backend/services/database"
)

type SelfController struct {
	db *database.DB
}

func NewSelfController(db *database.DB) *SelfController {
	return &SelfController{
		db: db,
	}
}

func (a *SelfController) ListSelfFiles(userId int64, path string) ([]database.File, error) {
	return a.db.ListFilesByUser(userId, path)
}

func (a *SelfController) AddSelfFile(userId int64, name, path string, data []byte) (int64, error) {

	now := time.Now()

	file := &database.File{
		Name:      name,
		Path:      path,
		OwnerUser: userId,
		FType:     "user",
		IsPublic:  false,
		Size:      int64(len(data)),
		CreatedAt: &now,
	}

	return a.db.AddFile(file, data)
}

func (a *SelfController) AddSelfFolder(userId int64, path, name string) (int64, error) {
	return a.db.AddFolder(0, userId, "user", path, name)
}

func (a *SelfController) GetSelfFile(userId int64, id int64) ([]byte, error) {
	// fixme check if user has access to file

	fmeta, err := a.db.GetFileMeta(id)
	if err != nil {
		return nil, err
	}

	if fmeta.FType != "user" {
		return nil, fmt.Errorf("file is not found")
	}

	if fmeta.OwnerUser != userId {
		return nil, fmt.Errorf("file not found")
	}

	return a.db.GetFileBlob(id)
}

func (a *SelfController) DeleteSelfFile(userId int64, id int64) error {
	// fixme check if user has access to file

	fmeta, err := a.db.GetFileMeta(id)
	if err != nil {
		return err
	}

	if fmeta.FType != "user" {
		return fmt.Errorf("file is not found")
	}

	if fmeta.OwnerUser != userId {
		return fmt.Errorf("file not found")
	}

	return a.db.DeleteFile(id)
}
