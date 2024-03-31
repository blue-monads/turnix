package database

import (
	"errors"
	"log"

	"github.com/upper/db/v4"
)

var errNotFound = errors.New("no found")

func (d *DB) AddProject(data *Project) (int64, error) {
	r, err := d.projectTable().Insert(data)
	if err != nil {
		return 0, err
	}

	return r.ID().(int64), nil

}

func (d *DB) RemoveProject(id int64, ownerId int64) error {
	return d.projectTable().Find(db.Cond{"id": id, "owner": ownerId}).Delete()
}

func (d *DB) GetProject(id int64) (*Project, error) {
	data := &Project{}

	err := d.projectTable().Find(db.Cond{"id": id}).One(data)
	if err != nil {
		return nil, err
	}

	return data, nil

}

func (d *DB) GetProjectByOwner(id int64, ownerId int64) (*Project, error) {
	data := &Project{}

	err := d.projectTable().Find(db.Cond{"id": id, "owner": ownerId}).One(data)
	if err != nil {
		return nil, err
	}

	return data, nil

}

func (d *DB) UpdateProject(id int64, ownerId int64, data map[string]any) error {
	return d.projectTable().Find(db.Cond{"id": id, "owner": ownerId}).Update(data)
}

func (d *DB) ListOwnProjects(ownerId int64, ptype string) ([]Project, error) {
	datas := make([]Project, 0)

	cond := db.Cond{"owner": ownerId}

	if ptype != "" {
		cond["ptype"] = ptype
	}

	err := d.projectTable().Find(cond).All(&datas)
	if err != nil {
		return nil, err
	}

	return datas, nil
}

type TPProjects struct {
	Project int64 `db:"projectId"`
}

func (d *DB) ListThirdPartyProjects(userid int64, ptype string) ([]Project, error) {

	cond := db.Cond{
		"userId": userid,
	}

	projs := make([]TPProjects, 0)

	if ptype != "" {
		cond["ptype"] = ptype
	}

	err := d.projectUserTable().Find(cond).Select("projectId").All(&projs)

	if err != nil {
		return nil, err
	}

	projIds := make([]int64, 0, len(projs))

	for _, p := range projs {
		projIds = append(projIds, p.Project)
	}

	datas := make([]Project, 0, len(projs))

	err = d.projectTable().Find(db.Cond{
		"userId": userid,
		"id IN":  projIds,
	}).All(&datas)

	if err != nil {
		return nil, err
	}

	return datas, nil
}

func (d *DB) AddUserToProject(ownerid int64, userId int64, projectId int64) error {

	if !d.isOwner(ownerid, projectId) {
		return errNotFound
	}

	data := &ProjectUser{
		ProjectID: (projectId),
		UserID:    (userId),
	}

	_, err := d.projectUserTable().Insert(data)

	return err
}

func (d *DB) RemoveUserFromPoject(ownerid int64, userId int64, projectId int64) error {
	if !d.isOwner(ownerid, projectId) {
		return errNotFound
	}

	return d.projectUserTable().Find(db.Cond{"projectId": projectId, "userId": userId}).Delete()
}

// private

func (d *DB) isOwner(ownerid int64, projId int64) bool {
	exist, err := d.projectTable().Find(db.Cond{"owner": ownerid, "id": projId}).Exists()

	if err != nil {
		log.Println("owner check error", err)
		return false
	}

	return exist

}

func (d *DB) userHasAccess(userId int64, projId int64) bool {
	exist, err := d.projectUserTable().Find(db.Cond{"userId": userId, "projectId": projId}).Exists()
	if err != nil {
		log.Println("owner check error", err)
		return false
	}

	return exist

}

func (d *DB) projectTable() db.Collection {
	return d.table("Projects")
}

func (d *DB) projectUserTable() db.Collection {
	return d.table("ProjectUsers")
}
