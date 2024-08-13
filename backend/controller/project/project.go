package project

import (
	"github.com/bornjre/turnix/backend/services/database"
	"github.com/bornjre/turnix/backend/xtypes/xproject"
)

type ProjectController struct {
	db       *database.DB
	projects map[string]*xproject.Defination
}

func NewProjectController(db *database.DB) *ProjectController {

	return &ProjectController{
		db: db,
	}
}
