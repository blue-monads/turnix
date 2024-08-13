package server

import (
	"strconv"

	"github.com/bornjre/turnix/backend/services/signer"
	"github.com/bornjre/turnix/backend/xtypes/models"
	"github.com/gin-gonic/gin"
)

func (a *Server) ListProjectTypes(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {
	return a.cProject.ListProjectTypes()
}

func (a *Server) GetProjectType(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {

	return a.cProject.GetProjectType(ctx.Param("ptype"))
}

func (a *Server) GetProjectTypeForm(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {

	return a.cProject.GetProjectTypeForm(ctx.Param("ptype"))
}

func (a *Server) listProjects(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {

	return a.cProject.ListProjects(claim.UserId, ctx.Query("ptype"))

}

func (a *Server) addProject(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {
	data := models.Project{}
	err := ctx.Bind(&data)
	if err != nil {
		return nil, err
	}

	data.OwnerID = (claim.UserId)

	return a.cProject.AddProject(claim.UserId, &data)
}

func (a *Server) removeProject(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {

	id, err := strconv.ParseInt(ctx.Param("pid"), 10, 64)
	if err != nil {
		return nil, err
	}

	err = a.cProject.RemoveProject(claim.UserId, id)
	return nil, err

}

func (a *Server) updateProject(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {

	id, err := strconv.ParseInt(ctx.Param("pid"), 10, 64)
	if err != nil {
		return nil, err
	}

	data := make(map[string]any)
	err = ctx.Bind(&data)
	if err != nil {
		return nil, err
	}

	err = a.cProject.UpdateProject(claim.UserId, id, data)
	return nil, err

}

func (a *Server) getProject(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {
	id, err := strconv.ParseInt(ctx.Param("pid"), 10, 64)
	if err != nil {
		return nil, err
	}

	return a.cProject.GetProject(claim.UserId, id)
}

func (a *Server) inviteUserToPoject(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {

	return nil, nil
}

func (a *Server) removeUserFromPoject(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {
	return nil, nil
}

// hooks

func (a *Server) listProjectHooks(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {

	pid, _ := strconv.ParseInt(ctx.Param("pid"), 10, 64)

	return a.cProject.ListProjectHooks(claim.UserId, pid)
}

func (a *Server) addProjectHook(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {
	pid, _ := strconv.ParseInt(ctx.Param("pid"), 10, 64)

	data := &models.ProjectHook{}
	err := ctx.Bind(&data)
	if err != nil {
		return nil, err
	}

	return a.cProject.AddProjectHook(claim.UserId, pid, data)
}

func (a *Server) removeProjectHook(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {

	pid, _ := strconv.ParseInt(ctx.Param("pid"), 10, 64)
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)

	err := a.cProject.RemoveProjectHook(claim.UserId, pid, id)
	return nil, err

}

func (a *Server) updateProjectHook(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {

	pid, _ := strconv.ParseInt(ctx.Param("pid"), 10, 64)
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)

	data := make(map[string]any)
	err := ctx.Bind(&data)
	if err != nil {
		return nil, err
	}

	err = a.cProject.UpdateProjectHook(claim.UserId, pid, id, data)

	return nil, err

}

func (a *Server) getProjectHook(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {

	pid, _ := strconv.ParseInt(ctx.Param("pid"), 10, 64)
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)

	return a.cProject.GetProjectHook(claim.UserId, pid, id)

}
