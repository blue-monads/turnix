package app

import (
	"strconv"

	"github.com/bornjre/trunis/backend/database"
	"github.com/bornjre/trunis/backend/token"
	"github.com/gin-gonic/gin"
)

func (a *App) listProjects(claim *token.AccessClaim, ctx *gin.Context) (any, error) {

	ownpjs, err := a.db.ListOwnProjects(claim.UserId)
	if err != nil {
		return nil, err
	}

	tprojs, err := a.db.ListThirdPartyProjects(claim.UserId)
	if err != nil {
		return nil, err
	}

	tprojs = append(tprojs, ownpjs...)

	return tprojs, nil

}

func (a *App) addProject(claim *token.AccessClaim, ctx *gin.Context) (any, error) {
	data := database.Project{}
	err := ctx.Bind(&data)
	if err != nil {
		return nil, err
	}

	data.OwnerID = (claim.UserId)

	return a.db.AddProject(&data)
}

func (a *App) removeProject(claim *token.AccessClaim, ctx *gin.Context) (any, error) {

	id, err := strconv.ParseInt(ctx.Param("pid"), 10, 64)
	if err != nil {
		return nil, err
	}

	err = a.db.RemoveProject(id, claim.UserId)
	return nil, err
}

func (a *App) updateProject(claim *token.AccessClaim, ctx *gin.Context) (any, error) {

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

func (a *App) getProject(claim *token.AccessClaim, ctx *gin.Context) (any, error) {
	id, err := strconv.ParseInt(ctx.Param("pid"), 10, 64)
	if err != nil {
		return nil, err
	}

	return a.db.GetProjectByOwner(id, claim.UserId)
}

func (a *App) inviteUserToPoject(claim *token.AccessClaim, ctx *gin.Context) (any, error) {

	return nil, nil
}

func (a *App) removeUserFromPoject(claim *token.AccessClaim, ctx *gin.Context) (any, error) {
	return nil, nil
}
