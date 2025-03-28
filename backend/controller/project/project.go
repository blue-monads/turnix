package project

import (
	"fmt"
	"os"
	"path"

	"github.com/blue-monads/turnix/backend/engine"
	"github.com/blue-monads/turnix/backend/services/database"
	xutils "github.com/blue-monads/turnix/backend/utils"
	"github.com/blue-monads/turnix/backend/xtypes/models"
	"github.com/blue-monads/turnix/backend/xtypes/xproject"
)

// this layer (controller) should not have claim and gin context

type ProjectController struct {
	db          *database.DB
	engine      *engine.Engine
	installPath string
}

func NewProjectController(db *database.DB, engine *engine.Engine) *ProjectController {

	return &ProjectController{
		db:          db,
		engine:      engine,
		installPath: engine.InstallPath(),
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

func (a *ProjectController) GetProjectTypeReload(ptype string) error {
	return a.engine.GetProjectTypeReload(ptype)
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

// project_type install

type ManifestMini struct {
	Slug string `json:"slug"`
}

func (a *ProjectController) InstallProjectType(userId int64, url string) error {

	os.MkdirAll(a.installPath, 0766)

	randstr, err := xutils.GenerateRandomString(10)
	if err != nil {
		return err
	}

	tempFileName := fmt.Sprint(a.installPath, "/__downloading_", randstr)

	err = engine.DownloadFile(url, tempFileName)
	if err != nil {
		os.Remove(tempFileName)
		return err
	}

	mf := ManifestMini{}

	err = engine.ReadManifestFromZip(tempFileName, &mf)
	if err != nil {
		os.Remove(tempFileName)
		return err
	}

	finalFileName := fmt.Sprintf("%s.zip", path.Join(a.installPath, mf.Slug))
	err = os.Rename(tempFileName, finalFileName)
	if err != nil {
		os.Remove(tempFileName)
		os.Remove(finalFileName)
		return err
	}

	a.engine.InformPtypeAdded(mf.Slug)

	return nil
}
