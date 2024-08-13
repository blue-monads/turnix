package controller

import (
	"github.com/bornjre/turnix/backend/controller/auth"
	"github.com/bornjre/turnix/backend/controller/project"
	"github.com/bornjre/turnix/backend/controller/self"
	"github.com/bornjre/turnix/backend/services/database"
)

type RootController struct {
	auth    *auth.AuthController
	project *project.ProjectController
	self    *self.SelfController
}

func New(db *database.DB) *RootController {
	return &RootController{
		auth:    auth.NewAuthController(db),
		project: project.NewProjectController(db),
		self:    self.NewSelfController(db),
	}
}
