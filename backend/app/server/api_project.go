package server

import (
	"fmt"
	"net/http"
	"strconv"

	"github.com/blue-monads/turnix/backend/services/database/autoquery"
	"github.com/blue-monads/turnix/backend/services/signer"
	"github.com/blue-monads/turnix/backend/xtypes/models"
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

func (a *Server) GetProjectTypeReload(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {

	err := a.cProject.GetProjectTypeReload(ctx.Param("ptype"))
	if err != nil {
		return nil, err
	}

	return nil, nil
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

// files

func (a *Server) listProjectFiles(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {

	pid, _ := strconv.ParseInt(ctx.Param("pid"), 10, 64)

	path := ctx.Query("path")

	return a.cProject.ListProjectFiles(claim.UserId, pid, path)
}

type CreateFolder struct {
	Name string `json:"name"`
	Path string `json:"path"`
}

func (a *Server) addProjectFolder(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {
	pid, _ := strconv.ParseInt(ctx.Param("pid"), 10, 64)

	data := CreateFolder{}
	err := ctx.Bind(&data)
	if err != nil {
		return nil, err
	}

	return a.cProject.AddProjectFolder(claim.UserId, pid, data.Path, data.Name)
}

func (a *Server) addProjectFile(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {
	pid, _ := strconv.ParseInt(ctx.Param("pid"), 10, 64)

	name := ctx.Query("name")
	path := ctx.Query("path")

	if name == "" {
		return nil, fmt.Errorf("name is required")
	}

	return a.cProject.AddProjectFile(claim.UserId, pid, name, path, ctx.Request.ContentLength, ctx.Request.Body)
}

func (a *Server) getProjectFile(ctx *gin.Context) {

	claim, err := a.withAccess(ctx)
	if err != nil {
		return
	}

	pid, _ := strconv.ParseInt(ctx.Param("pid"), 10, 64)
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)

	err = a.cProject.GetProjectFile(claim.UserId, pid, id, ctx)
	if err != nil {
		ctx.AbortWithError(http.StatusBadRequest, err)
		return
	}

}

func (a *Server) removeProjectFile(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {
	pid, _ := strconv.ParseInt(ctx.Param("pid"), 10, 64)
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)

	err := a.cProject.DeleteProjectFile(claim.UserId, pid, id)

	return nil, err
}

type ProjectSQLExec struct {
	Input string `json:"input"`
	Name  string `json:"name"`
	Data  []any  `json:"data"`
}

func (a *Server) runProjectSQL(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {
	pid, _ := strconv.ParseInt(ctx.Param("pid"), 10, 64)

	data := ProjectSQLExec{}
	err := ctx.Bind(&data)
	if err != nil {
		return nil, err
	}

	res, err := a.cProject.RunQuerySQL(claim.UserId, pid, data.Input, data.Name, data.Data)

	return res, err
}

type ProjectSQLExec2 struct {
	QStr string `json:"qstr"`
	Data []any  `json:"data"`
}

func (a *Server) runProjectSQL2(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {
	pid, _ := strconv.ParseInt(ctx.Param("pid"), 10, 64)

	data := ProjectSQLExec2{}
	err := ctx.Bind(&data)
	if err != nil {
		return nil, err
	}

	res, err := a.cProject.RunQuerySQL2(claim.UserId, pid, data.QStr, data.Data)
	if err != nil {
		return nil, err
	}

	return res, err
}

func (a *Server) listProjectTables(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {
	pid, _ := strconv.ParseInt(ctx.Param("pid"), 10, 64)

	return a.cProject.ListProjectTables(claim.UserId, pid)
}

func (a *Server) listProjectTableColumns(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {
	pid, _ := strconv.ParseInt(ctx.Param("pid"), 10, 64)
	table := ctx.Param("table")

	return a.cProject.ListProjectTableColumns(claim.UserId, pid, table)
}

func (a *Server) autoQueryProjectTable(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {
	pid, _ := strconv.ParseInt(ctx.Param("pid"), 10, 64)
	table := ctx.Query("table")

	if table == "" {
		return nil, fmt.Errorf("table is required")
	}

	opts := autoquery.LoaderParams{}
	err := ctx.Bind(&opts)
	if err != nil {
		return nil, err
	}

	return a.cProject.AutoQueryProjectTable(claim.UserId, pid, table, opts)
}

// plugins
func (a *Server) listProjectPlugins(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {

	pid, _ := strconv.ParseInt(ctx.Param("pid"), 10, 64)

	return a.cProject.ListProjectPlugins(claim.UserId, pid)
}

func (a *Server) addProjectPlugin(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {
	pid, _ := strconv.ParseInt(ctx.Param("pid"), 10, 64)

	data := models.ProjectPlugin{}
	err := ctx.Bind(&data)
	if err != nil {
		return nil, err
	}

	return a.cProject.AddProjectPlugin(claim.UserId, pid, &data)
}

func (a *Server) removeProjectPlugin(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {

	pid, _ := strconv.ParseInt(ctx.Param("pid"), 10, 64)
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)

	err := a.cProject.RemoveProjectPlugin(claim.UserId, pid, id)
	return nil, err
}

func (a *Server) updateProjectPlugin(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {

	pid, _ := strconv.ParseInt(ctx.Param("pid"), 10, 64)
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)

	data := make(map[string]any)
	err := ctx.Bind(&data)
	if err != nil {
		return nil, err
	}

	err = a.cProject.UpdateProjectPlugin(claim.UserId, pid, id, data)

	return nil, err
}

func (a *Server) getProjectPlugin(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {

	pid, _ := strconv.ParseInt(ctx.Param("pid"), 10, 64)
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)

	return a.cProject.GetProjectPlugin(claim.UserId, pid, id)
}

type InstallOptions struct {
	Url string `json:"url,omitempty"`
}

func (a *Server) installProjectType(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {
	data := InstallOptions{}
	err := ctx.Bind(&data)
	if err != nil {
		return nil, err
	}

	err = a.cProject.InstallProjectType(claim.UserId, data.Url)
	if err != nil {
		return nil, err
	}

	return nil, nil

}
