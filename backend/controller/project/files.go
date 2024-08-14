package project

import (
	"fmt"
	"time"

	"github.com/bornjre/turnix/backend/services/database"
)

func (a *ProjectController) ListProjectFiles(userId int64, pid int64) ([]database.File, error) {
	return a.db.ListFilesByProject(pid, "")
}

func (a *ProjectController) AddProjectFile(userId int64, pid int64, name string, data []byte) (int64, error) {

	now := time.Now()

	file := &database.File{
		Name:      name,
		Path:      "",
		OwnerUser: userId,
		OwnerProj: pid,
		FType:     "project",
		IsPublic:  false,
		Size:      int64(len(data)),
		CreatedAt: &now,
	}

	return a.db.AddFile(file, data)
}

func (a *ProjectController) GetProjectFile(userId int64, pid int64, id int64) ([]byte, error) {
	// fixme check if user has access to file

	fmeta, err := a.db.GetFileMeta(id)
	if err != nil {
		return nil, err
	}

	if fmeta.OwnerProj != pid {
		return nil, fmt.Errorf("file not found")
	}

	if fmeta.FType != "project" {
		return nil, fmt.Errorf("file is not found")
	}

	return a.db.GetFileBlob(id)
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
