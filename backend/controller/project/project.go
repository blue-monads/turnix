package project

import (
	"github.com/bornjre/turnix/backend/services/database"
	"github.com/bornjre/turnix/backend/xtypes/models"
	"github.com/bornjre/turnix/backend/xtypes/xproject"
)

// this layer (controller) should not have claim and gin context

type ProjectController struct {
	db       *database.DB
	projects map[string]*xproject.Defination
}

func NewProjectController(db *database.DB, projects map[string]*xproject.Defination) *ProjectController {

	return &ProjectController{
		db:       db,
		projects: projects,
	}
}

func (a *ProjectController) ListProjectTypes() ([]models.ProjectTypes, error) {
	pdefs := make([]models.ProjectTypes, 0)

	for _, pdef := range a.projects {
		pdefs = append(pdefs, models.ProjectTypes{
			Name:       pdef.Name,
			Ptype:      pdef.Slug,
			Icon:       pdef.Icon,
			Info:       pdef.Info,
			IsExternal: pdef.AssetData != nil,
		})

	}

	return pdefs, nil
}

func (a *ProjectController) GetProjectType(ptype string) (*models.ProjectTypes, error) {

	for _, pdef := range a.projects {

		if pdef.Slug == ptype {
			return &models.ProjectTypes{
				Name:       pdef.Name,
				Ptype:      pdef.Slug,
				Slug:       pdef.Slug,
				Info:       pdef.Info,
				Icon:       pdef.Icon,
				IsExternal: pdef.AssetData != nil,
			}, nil
		}

	}

	return nil, nil
}

func (a *ProjectController) GetProjectTypeForm(ptype string) ([]xproject.PTypeField, error) {

	for _, pdef := range a.projects {

		if pdef.Slug == ptype {
			return pdef.NewFormSchemaFields, nil
		}

	}

	return nil, nil
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

	ptype := a.projects[data.Ptype]

	if ptype.OnInit != nil {
		err = ptype.OnInit(id)
		if err != nil {
			return 0, err
		}
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
