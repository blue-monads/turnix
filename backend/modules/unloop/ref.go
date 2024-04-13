package unloop

/*


import (
	"github.com/k0kubun/pp"
	"github.com/upper/db/v4"
)

// queue messages

func (d *DB) AddQueueMessage(data *PTLactionQueueMessage) (int64, error) {

	if !d.userHasAccess(int64(data.Submitter), int64(data.ProjectID)) {
		if !d.isOwner(int64(data.Submitter), int64(data.ProjectID)) {
			pp.Println("@DD", data)

			return 0, errNotFound
		}
	}

	r, err := d.queueMessageTable().Insert(data)
	if err != nil {
		return 0, err
	}

	return r.ID().(int64), nil

}

func (d *DB) RemoveQueueMessage(ownerId int64, projectId int64, id int64) error {
	if !d.isOwner(ownerId, projectId) {
		return errNotFound
	}

	return d.queueMessageTable().Find(db.Cond{"projectId": projectId, "id": id}).Delete()
}

func (d *DB) UpdateQueueMessage(ownerId, projectId, id int64, data map[string]any) error {
	if !d.isOwner(ownerId, projectId) {
		return errNotFound
	}

	return d.queueMessageTable().Find(db.Cond{"projectId": projectId, "id": id}).Update(data)
}

func (d *DB) GetQueueMessage(ownerId, projectId, id int64) (*PTLactionQueueMessage, error) {
	if !d.isOwner(ownerId, projectId) {
		return nil, errNotFound
	}

	data := &PTLactionQueueMessage{}

	err := d.queueMessageTable().Find(db.Cond{"projectId": projectId, "id": id}).One(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

type QQueryOptions struct {
	LastId     int64
	TemplateId int64
}

func (d *DB) QueryQueueMessage(ownerId int64, projectId int64, opts QQueryOptions) ([]PTLactionQueueMessage, error) {
	if !d.isOwner(ownerId, projectId) {
		return nil, errNotFound
	}

	datas := make([]PTLactionQueueMessage, 0)

	cond := db.Cond{"projectId": projectId, "id >": opts.LastId}
	if opts.TemplateId != 0 {
		cond["templateId"] = opts.TemplateId
	}

	err := d.queueMessageTable().Find(cond).All(&datas)
	if err != nil {
		return nil, err
	}

	return datas, nil
}

// template

func (d *DB) AddTemplate(ownerId int64, data *PTLactionTemplate) (int64, error) {
	if !d.isOwner(ownerId, int64(data.ProjectID)) {
		return 0, errNotFound
	}

	r, err := d.templatesTable().Insert(data)
	if err != nil {
		return 0, err
	}

	return r.ID().(int64), nil

}

func (d *DB) GetTemplate(ownerId int64, projectId int64, id int64) (*PTLactionTemplate, error) {
	if !d.isOwner(ownerId, projectId) {
		return nil, errNotFound
	}

	data := &PTLactionTemplate{}

	err := d.templatesTable().Find(db.Cond{"id": id, "projectId": projectId}).One(data)
	if err != nil {
		return nil, err
	}

	return data, nil

}

func (d *DB) RemoveTemplate(ownerId int64, projectId int64, id int64) error {
	if !d.isOwner(ownerId, projectId) {
		return errNotFound
	}

	return d.templatesTable().Find(db.Cond{"id": id, "projectId": projectId}).Delete()
}

func (d *DB) UpdateTemplate(ownerId int64, id int64, projectId int64, data map[string]any) error {
	if !d.isOwner(ownerId, projectId) {
		return errNotFound
	}

	return nil
}

func (d *DB) ListTemplate(ownerId int64, projectId int64) ([]PTLactionTemplate, error) {
	pp.Println("@DD", ownerId)

	if !d.isOwner(ownerId, projectId) {
		return nil, errNotFound
	}

	datas := make([]PTLactionTemplate, 0)

	err := d.templatesTable().Find(db.Cond{"projectId": projectId}).All(&datas)
	if err != nil {
		return nil, err
	}

	return datas, nil
}

// private

func (d *DB) templatesTable() db.Collection {
	return d.table("pt_onloop_Templates")
}

// private

func (d *DB) queueMessageTable() db.Collection {
	return d.table("pt_onloop_QueueMessages")
}


*/

/*


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



*/

/*

	apiv1.GET("/pt/laction/template/:pid", a.accessMiddleware(a.listTemplates))
	apiv1.POST("/pt/laction/template/:pid", a.accessMiddleware(a.addTemplate))
	apiv1.POST("/pt/laction/template/:pid/:tid", a.accessMiddleware(a.updateTemplate))
	apiv1.GET("/pt/laction/template/:pid/:tid", a.accessMiddleware(a.getTemplate))
	apiv1.DELETE("/pt/laction/template/:pid/:tid", a.accessMiddleware(a.removeTemplate))

	apiv1.POST("/pt/laction/template/:pid/:tid/push", a.accessMiddleware(a.pushQueueMessage))

	apiv1.GET("/pt/laction/queue/:pid", a.accessMiddleware(a.listQueueMessages))
	apiv1.POST("/pt/laction/queue/:pid", a.accessMiddleware(a.addQueueMessage))
	apiv1.POST("/pt/laction/queue/:pid/:qid", a.accessMiddleware(a.updateQueueMessage))
	apiv1.GET("/pt/laction/queue/:pid/:qid", a.accessMiddleware(a.getQueueMessage))
	apiv1.DELETE("/pt/laction/queue/:pid/:qid", a.accessMiddleware(a.removeUpdateQueueMessage))

	apiv1.POST("/pt/laction/query/:pid", a.accessMiddleware(a.queryQueueMessage))

*/
