package xdatabase

import (
	"errors"

	"github.com/bornjre/trunis/backend/xtypes/models"
	"github.com/gobuffalo/fizz"
	"github.com/upper/db/v4"
)

const (
	ScopeOwner = "owner"
)

var (
	ErrUserNoScope = errors.New("err: user doesnot have required scope")
)

type DDLContext struct {
	Tables []fizz.Table
	DDL    string
}

type Database interface {
	Init() error
	GetSession() db.Session
	Table(name string) db.Collection
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

	GetProjectUserScope(userId int64, projectId int64) (string, error)

	Vender() string

	RunDDL(ctx DDLContext) error
}
