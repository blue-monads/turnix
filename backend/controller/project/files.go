package project

import (
	"fmt"
	"io"
	"time"

	"github.com/bornjre/turnix/backend/services/database"
	"github.com/gin-gonic/gin"
)

func (a *ProjectController) ListProjectFiles(userId int64, pid int64, path string) ([]database.File, error) {
	return a.db.ListFilesByProject(pid, path)
}

func (a *ProjectController) AddProjectFile(userId int64, pid int64, name, path string, size int64, stream io.Reader) (int64, error) {

	now := time.Now()

	file := &database.File{
		Name:      name,
		Path:      path,
		OwnerUser: userId,
		OwnerProj: pid,
		FType:     "project",
		Size:      size,
		CreatedAt: &now,
	}

	return a.db.AddFileStreaming(file, stream)
}

func (a *ProjectController) AddProjectFolder(userId int64, pid int64, path, name string) (int64, error) {
	return a.db.AddFolder(pid, userId, "project", path, name)
}

func (a *ProjectController) GetProjectFile(userId int64, pid int64, id int64, ctx *gin.Context) error {
	// fixme check if user has access to file

	fmeta, err := a.db.GetFileMeta(id)
	if err != nil {
		return err
	}

	if fmeta.OwnerProj != pid {
		return fmt.Errorf("file not found")
	}

	if fmeta.FType != "project" {
		return fmt.Errorf("file is not found")
	}

	return a.db.GetFileBlobStreaming(id, ctx.Writer)
}

func (a *ProjectController) DeleteProjectFile(userId int64, pid int64, id int64) error {
	// fixme check if user has access to file

	fmeta, err := a.db.GetFileMeta(id)
	if err != nil {
		return err
	}

	if fmeta.OwnerProj != pid {
		return fmt.Errorf("file not found")
	}

	if fmeta.FType != "project" {
		return fmt.Errorf("file is not found")
	}

	return a.db.DeleteFile(id)
}
