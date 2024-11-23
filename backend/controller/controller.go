package controller

import (
	"github.com/blue-monads/turnix/backend/controller/auth"
	"github.com/blue-monads/turnix/backend/controller/common"
	"github.com/blue-monads/turnix/backend/controller/project"
	"github.com/blue-monads/turnix/backend/controller/self"
	"github.com/blue-monads/turnix/backend/services/database"
	"github.com/blue-monads/turnix/backend/xtypes/xproject"
)

type RootController struct {
	auth    *auth.AuthController
	project *project.ProjectController
	self    *self.SelfController
	common  *common.CommonController
}

func New(db *database.DB, projects map[string]*xproject.Defination) *RootController {
	return &RootController{
		auth:    auth.NewAuthController(db),
		project: project.NewProjectController(db, projects),
		self:    self.NewSelfController(db),
		common:  common.New(db),
	}
}

func (c *RootController) GetAuthController() *auth.AuthController {
	return c.auth
}

func (c *RootController) GetProjectController() *project.ProjectController {
	return c.project
}

func (c *RootController) GetSelfController() *self.SelfController {
	return c.self
}

func (c *RootController) GetCommonController() *common.CommonController {
	return c.common
}
