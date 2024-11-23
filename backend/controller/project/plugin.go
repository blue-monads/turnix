package project

import (
	"github.com/blue-monads/turnix/backend/xtypes/models"
)

func (a *ProjectController) ListProjectPlugins(userId int64, pid int64) ([]models.ProjectPlugin, error) {
	return a.db.ListProjectPlugins(userId, pid)
}

func (a *ProjectController) AddProjectPlugin(userId int64, pid int64, data *models.ProjectPlugin) (int64, error) {
	return a.db.AddProjectPlugin(pid, userId, data)
}

func (a *ProjectController) RemoveProjectPlugin(userId int64, pid int64, id int64) error {
	return a.db.RemoveProjectPlugin(userId, pid, id)
}

func (a *ProjectController) UpdateProjectPlugin(userId int64, pid int64, id int64, data map[string]any) error {
	return a.db.UpdateProjectPlugin(userId, pid, id, data)
}

func (a *ProjectController) GetProjectPlugin(userId int64, pid int64, id int64) (*models.ProjectPlugin, error) {
	return a.db.GetProjectPlugin(userId, pid, id)
}
