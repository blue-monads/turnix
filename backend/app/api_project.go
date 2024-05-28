package app

import (
	"strconv"

	"github.com/bornjre/trunis/backend/services/signer"
	"github.com/bornjre/trunis/backend/xtypes/models"
	"github.com/gin-gonic/gin"
)

func (a *App) ListProjectTypes(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {

	pdefs := make([]models.ProjectTypes, 0)

	for _, pdef := range a.ptypeDefs {
		pdefs = append(pdefs, models.ProjectTypes{
			Name:       pdef.Name,
			Ptype:      pdef.Slug,
			Icon:       pdef.Icon,
			Info:       pdef.Info,
			IsExternal: pdef.AssetData != nil,
			EventTypes: pdef.EventTypes,
		})

	}

	return pdefs, nil
}

func (a *App) GetProjectType(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {

	for _, pdef := range a.ptypeDefs {

		if pdef.Slug == ctx.Param("ptype") {
			return &models.ProjectTypes{
				Name:       pdef.Name,
				Ptype:      pdef.Slug,
				Slug:       pdef.Slug,
				Info:       pdef.Info,
				Icon:       pdef.Icon,
				IsExternal: pdef.AssetData != nil,
				EventTypes: pdef.EventTypes,
			}, nil
		}

	}

	return nil, nil
}

func (a *App) GetProjectTypeForm(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {

	for _, pdef := range a.ptypeDefs {

		if pdef.Slug == ctx.Param("ptype") {
			return pdef.NewFormSchemaFields, nil
		}

	}

	return nil, nil
}

func (a *App) listProjects(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {

	ptype := ctx.Query("ptype")

	ownpjs, err := a.db.ListOwnProjects(claim.UserId, ptype)
	if err != nil {
		return nil, err
	}

	tprojs, err := a.db.ListThirdPartyProjects(claim.UserId, ptype)
	if err != nil {
		return nil, err
	}

	tprojs = append(tprojs, ownpjs...)

	return tprojs, nil

}

func (a *App) addProject(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {
	data := models.Project{}
	err := ctx.Bind(&data)
	if err != nil {
		return nil, err
	}

	data.OwnerID = (claim.UserId)

	return a.db.AddProject(&data)
}

func (a *App) removeProject(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {

	id, err := strconv.ParseInt(ctx.Param("pid"), 10, 64)
	if err != nil {
		return nil, err
	}

	err = a.db.RemoveProject(id, claim.UserId)
	return nil, err
}

func (a *App) updateProject(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {

	id, err := strconv.ParseInt(ctx.Param("pid"), 10, 64)
	if err != nil {
		return nil, err
	}

	data := make(map[string]any)
	err = ctx.Bind(&data)
	if err != nil {
		return nil, err
	}

	err = a.db.UpdateProject(id, claim.UserId, data)

	return nil, err

}

func (a *App) getProject(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {
	id, err := strconv.ParseInt(ctx.Param("pid"), 10, 64)
	if err != nil {
		return nil, err
	}

	return a.db.GetProjectByOwner(id, claim.UserId)
}

func (a *App) inviteUserToPoject(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {

	return nil, nil
}

func (a *App) removeUserFromPoject(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {
	return nil, nil
}

// hooks

func (a *App) listProjectHooks(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {

	pid, _ := strconv.ParseInt(ctx.Param("pid"), 10, 64)

	return a.db.ListProjectHooks(claim.UserId, pid)
}

func (a *App) addProjectHook(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {
	pid, _ := strconv.ParseInt(ctx.Param("pid"), 10, 64)

	data := &models.ProjectHook{}
	err := ctx.Bind(&data)
	if err != nil {
		return nil, err
	}

	return a.db.AddProjectHook(claim.UserId, pid, data)
}

func (a *App) removeProjectHook(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {

	pid, _ := strconv.ParseInt(ctx.Param("pid"), 10, 64)
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)

	err := a.db.RemoveProjectHook(claim.UserId, pid, id)
	return nil, err

}

func (a *App) updateProjectHook(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {

	pid, _ := strconv.ParseInt(ctx.Param("pid"), 10, 64)
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)

	data := make(map[string]any)
	err := ctx.Bind(&data)
	if err != nil {
		return nil, err
	}

	err = a.db.UpdateProjectHook(claim.UserId, pid, id, data)

	return nil, err

}

func (a *App) getProjectHook(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {

	pid, _ := strconv.ParseInt(ctx.Param("pid"), 10, 64)
	id, _ := strconv.ParseInt(ctx.Param("id"), 10, 64)

	return a.db.GetProjectHook(claim.UserId, pid, id)

}
