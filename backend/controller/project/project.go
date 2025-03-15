package project

import (
	"github.com/blue-monads/turnix/backend/engine"
	"github.com/blue-monads/turnix/backend/services/database"
	"github.com/blue-monads/turnix/backend/xtypes/models"
	"github.com/blue-monads/turnix/backend/xtypes/xproject"
)

// this layer (controller) should not have claim and gin context

type ProjectController struct {
	db     *database.DB
	engine *engine.Engine
}

func NewProjectController(db *database.DB, projects map[string]*xproject.Defination) *ProjectController {

	return &ProjectController{
		db: db,
	}
}

func (a *ProjectController) ListProjectTypes() ([]models.ProjectTypes, error) {
	return a.engine.ListProjectTypes()
}

func (a *ProjectController) GetProjectType(ptype string) (*models.ProjectTypes, error) {
	return a.engine.GetProjectType(ptype)
}

func (a *ProjectController) GetProjectTypeForm(ptype string) ([]xproject.PTypeField, error) {
	return a.engine.GetProjectTypeForm(ptype)
}

func (a *ProjectController) ListProjects(userId int64, ptype string) ([]models.Project, error) {
	ownpjs, err := a.db.ListOwnProjects(userId, ptype)
	if err != nil {
		return nil, err
	}

	tprojs, err := a.db.ListThirdPartyProjects(userId, ptype)
	if err != nil {
		return nil, err
	}

	tprojs = append(tprojs, ownpjs...)
	return tprojs, nil

}

func (a *ProjectController) AddProject(userId int64, data *models.Project) (int64, error) {
	data.OwnerID = userId

	id, err := a.db.AddProject(data)
	if err != nil {
		return 0, err
	}

	err = a.engine.OnInit(data.Ptype, id)
	if err != nil {
		return 0, err
	}

	return id, nil

}

func (a *ProjectController) RemoveProject(userId int64, pid int64) error {
	return a.db.RemoveProject(pid, userId)
}

func (a *ProjectController) UpdateProject(userId int64, pid int64, data map[string]any) error {
	return a.db.UpdateProject(pid, userId, data)
}

func (a *ProjectController) GetProject(userId int64, pid int64) (*models.Project, error) {
	return a.db.GetProjectByOwner(pid, userId)
}

func (a *ProjectController) InviteUserToPoject(userId int64, pid int64, data any) (int64, error) {
	panic("implement me")
}

func (a *ProjectController) RemoveUserFromPoject(userId int64, pid int64, uid int64) (int64, error) {
	panic("implement me")
}

// hooks

func (a *ProjectController) ListProjectHooks(userId int64, pid int64) ([]models.ProjectHook, error) {
	return a.db.ListProjectHooksByUser(userId, pid)
}

func (a *ProjectController) AddProjectHook(userId int64, pid int64, data *models.ProjectHook) (int64, error) {
	return a.db.AddProjectHook(userId, pid, data)
}

func (a *ProjectController) RemoveProjectHook(userId int64, pid int64, id int64) error {
	return a.db.RemoveProjectHook(userId, pid, id)
}

func (a *ProjectController) UpdateProjectHook(userId int64, pid int64, id int64, data map[string]any) error {
	return a.db.UpdateProjectHook(userId, pid, id, data)
}

func (a *ProjectController) GetProjectHook(userId int64, pid int64, id int64) (*models.ProjectHook, error) {
	return a.db.GetProjectHook(userId, pid, id)
}
