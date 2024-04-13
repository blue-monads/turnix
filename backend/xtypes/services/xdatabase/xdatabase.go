package xdatabase

import (
	"github.com/bornjre/trunis/backend/xtypes/models"
	"github.com/upper/db/v4"
)

type Database interface {
	Init() error
	GetSession() db.Session
	Close() error
	RunSeed() error

	AddUser(data *models.User) (int64, error)
	GetUser(id int64) (*models.User, error)
	UpdateUser(id int64, data map[string]any) error
	GetUserByEmail(email string) (*models.User, error)
	ListUser() ([]models.User, error)
	DeleteUser(id int64) error

	AddUserToProject(ownerid int64, userId int64, projectId int64) error
	AddProject(data *models.Project) (int64, error)
	GetProject(id int64) (*models.Project, error)
	GetProjectByOwner(id int64, ownerId int64) (*models.Project, error)
	RemoveUserFromPoject(ownerid int64, userId int64, projectId int64) error
	ListOwnProjects(ownerId int64, ptype string) ([]models.Project, error)
	ListThirdPartyProjects(userid int64, ptype string) ([]models.Project, error)
	ListUserByOwner(owner int64) ([]models.User, error)
	RemoveProject(id int64, ownerId int64) error
	UpdateProject(id int64, ownerId int64, data map[string]any) error
}