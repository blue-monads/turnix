package database

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

	err := d.queueMessageTable().Find(db.Cond{"projectId": projectId, "id": id}).Update(data)
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
	return d.table("PTLactionTemplates")
}

// private

func (d *DB) queueMessageTable() db.Collection {
	return d.table("PTLactionQueueMessages")
}
