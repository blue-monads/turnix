package app

import (
	"strconv"

	"github.com/bornjre/trunis/backend/database"
	"github.com/bornjre/trunis/backend/services/signer"
	"github.com/gin-gonic/gin"
)

func (a *App) listQueueMessages(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {

	id, err := strconv.ParseInt(ctx.Param("pid"), 10, 64)
	if err != nil {
		return nil, err
	}

	return a.db.QueryQueueMessage(claim.UserId, id, database.QQueryOptions{
		LastId: 0,
	})

}

func (a *App) addQueueMessage(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {
	data := &database.PTLactionQueueMessage{}

	err := ctx.Bind(data)
	if err != nil {
		return nil, err
	}

	data.Submitter = claim.UserId

	return a.db.AddQueueMessage(data)
}

func (a *App) getQueueMessage(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {
	pid, err := strconv.ParseInt(ctx.Param("pid"), 10, 64)
	if err != nil {
		return nil, err
	}

	qid, err := strconv.ParseInt(ctx.Param("qid"), 10, 64)
	if err != nil {
		return nil, err
	}

	return a.db.GetQueueMessage(claim.UserId, pid, qid)
}

func (a *App) removeUpdateQueueMessage(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {

	pid, err := strconv.ParseInt(ctx.Param("pid"), 10, 64)
	if err != nil {
		return nil, err
	}

	qid, err := strconv.ParseInt(ctx.Param("qid"), 10, 64)
	if err != nil {
		return nil, err
	}

	err = a.db.RemoveQueueMessage(claim.UserId, pid, qid)
	return nil, err
}

func (a *App) updateQueueMessage(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {
	pid, err := strconv.ParseInt(ctx.Param("pid"), 10, 64)
	if err != nil {
		return nil, err
	}

	qid, err := strconv.ParseInt(ctx.Param("qid"), 10, 64)
	if err != nil {
		return nil, err
	}

	data := make(map[string]any)

	err = ctx.Bind(&data)
	if err != nil {
		return nil, err
	}

	err = a.db.UpdateQueueMessage(claim.UserId, pid, qid, data)

	return nil, err
}

func (a *App) pushQueueMessage(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {

	pid, err := strconv.ParseInt(ctx.Param("pid"), 10, 64)
	if err != nil {
		return nil, err
	}

	tid, err := strconv.ParseInt(ctx.Param("tid"), 10, 64)
	if err != nil {
		return nil, err
	}

	data := &database.PTLactionQueueMessage{}
	err = ctx.Bind(data)
	if err != nil {
		return nil, err
	}

	data.ProjectID = (pid)
	data.TemplateID = (tid)

	return a.db.AddQueueMessage(data)

}

func (a *App) queryQueueMessage(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {

	id, err := strconv.ParseInt(ctx.Param("pid"), 10, 64)
	if err != nil {
		return nil, err
	}

	query := database.QQueryOptions{}

	err = ctx.Bind(&query)
	if err != nil {
		return nil, err
	}

	return a.db.QueryQueueMessage(claim.UserId, id, query)

}

// templates

func (a *App) listTemplates(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {
	id, err := strconv.ParseInt(ctx.Param("pid"), 10, 64)
	if err != nil {
		return nil, err
	}

	return a.db.ListTemplate(claim.UserId, id)
}

func (a *App) addTemplate(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {

	id, err := strconv.ParseInt(ctx.Param("pid"), 10, 64)
	if err != nil {
		return nil, err
	}

	data := &database.PTLactionTemplate{}
	err = ctx.Bind(data)
	if err != nil {
		return nil, err
	}

	data.ProjectID = (id)

	return a.db.AddTemplate(claim.UserId, data)
}

func (a *App) getTemplate(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {
	pid, err := strconv.ParseInt(ctx.Param("pid"), 10, 64)
	if err != nil {
		return nil, err
	}

	tid, err := strconv.ParseInt(ctx.Param("tid"), 10, 64)
	if err != nil {
		return nil, err
	}

	return a.db.GetTemplate(claim.UserId, pid, tid)
}

func (a *App) removeTemplate(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {

	pid, err := strconv.ParseInt(ctx.Param("pid"), 10, 64)
	if err != nil {
		return nil, err
	}

	tid, err := strconv.ParseInt(ctx.Param("tid"), 10, 64)
	if err != nil {
		return nil, err
	}

	err = a.db.RemoveTemplate(claim.UserId, pid, tid)
	return nil, err
}

func (a *App) updateTemplate(claim *signer.AccessClaim, ctx *gin.Context) (any, error) {
	pid, err := strconv.ParseInt(ctx.Param("pid"), 10, 64)
	if err != nil {
		return nil, err
	}

	tid, err := strconv.ParseInt(ctx.Param("tid"), 10, 64)
	if err != nil {
		return nil, err
	}

	data := make(map[string]any)
	err = ctx.Bind(data)
	if err != nil {
		return nil, err
	}

	err = a.db.UpdateTemplate(claim.UserId, tid, pid, data)
	return nil, err
}
