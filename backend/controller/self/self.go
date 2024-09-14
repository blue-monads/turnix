package self

import (
	"fmt"
	"io"
	"time"

	"github.com/bornjre/turnix/backend/services/database"
	"github.com/gin-gonic/gin"
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

func (a *SelfController) AddSelfFile(userId int64, name, path string, size int64, stream io.Reader) (int64, error) {

	now := time.Now()

	file := &database.File{
		Name:      name,
		Path:      path,
		OwnerUser: userId,
		FType:     "user",
		Size:      size,
		CreatedAt: &now,
	}

	return a.db.AddFileStreaming(file, stream)
}

func (a *SelfController) AddSelfFolder(userId int64, path, name string) (int64, error) {
	return a.db.AddFolder(0, userId, "user", path, name)
}

func (a *SelfController) GetSelfFile(userId int64, id int64, ctx *gin.Context) error {
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

	return a.db.GetFileBlobStreaming(id, ctx.Writer)
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

// file shares

func (a *SelfController) ListFileShares(userId int64, fileId int64) ([]database.FileShare, error) {
	return a.db.ListFileShares(fileId)
}

func (a *SelfController) AddFileShare(userId int64, fileId int64) (int64, error) {
	return a.db.AddFileShare(fileId, userId)
}

func (a *SelfController) DeleteFileShare(userId int64, Id string) error {
	return a.db.DeleteFileShare(userId, Id)
}
