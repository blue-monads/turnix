package database

import (
	"errors"
	"fmt"
	"log"
	"strings"
	"time"

	"github.com/blue-monads/turnix/backend/utils/libx/dbutils"
	"github.com/blue-monads/turnix/backend/xtypes/models"
	"github.com/k0kubun/pp"
	"github.com/upper/db/v4"
)

var errNotFound = errors.New("no found")

func (d *DB) AddProject(data *models.Project) (int64, error) {
	r, err := d.projectTable().Insert(data)
	if err != nil {
		return 0, err
	}

	return r.ID().(int64), nil

}

func (d *DB) RemoveProject(id int64, ownerId int64) error {
	return d.projectTable().Find(db.Cond{"id": id, "owned_by": ownerId}).Delete()
}

func (d *DB) GetProject(id int64) (*models.Project, error) {
	data := &models.Project{}

	err := d.projectTable().Find(db.Cond{"id": id}).One(data)
	if err != nil {
		return nil, err
	}

	return data, nil

}

func (d *DB) GetProjectByOwner(id int64, ownerId int64) (*models.Project, error) {
	data := &models.Project{}

	err := d.projectTable().Find(db.Cond{"id": id, "owned_by": ownerId}).One(data)
	if err != nil {
		return nil, err
	}

	return data, nil

}

func (d *DB) UpdateProject(id int64, ownerId int64, data map[string]any) error {
	return d.projectTable().Find(db.Cond{"id": id, "owned_by": ownerId}).Update(data)
}

func (d *DB) ListOwnProjects(ownerId int64, ptype string) ([]models.Project, error) {
	datas := make([]models.Project, 0)

	cond := db.Cond{"owned_by": ownerId}

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

func (d *DB) RunProjectSQLQuery(pid int64, query string, data []any) ([]map[string]any, error) {

	parameterizedQuery := strings.ReplaceAll(query, "__project__", fmt.Sprintf("z_%d_", pid))

	rows, err := d.sess.SQL().Query(parameterizedQuery, data...)
	if err != nil {
		return nil, err
	}

	return dbutils.SelectScan(rows)

}

func (d *DB) ListProjectTables(pid int64) ([]string, error) {

	result := make([]string, 0)

	rows, err := d.sess.SQL().Query(fmt.Sprintf(`SELECT name FROM sqlite_master WHERE type='table' AND name NOT LIKE 'sqlite_%%' and name like 'z_%d_%%'`, pid))
	if err != nil {
		return nil, err
	}

	prefix := fmt.Sprintf("z_%d_", pid)

	for rows.Next() {
		var name string
		err := rows.Scan(&name)
		if err != nil {
			return nil, err
		}

		result = append(result, strings.TrimLeft(name, prefix))
	}

	return result, nil
}

/*

cid	name	type	notnull	dflt_value	pk

*/

func (d *DB) ListProjectTableColumns(pid int64, table string) ([]models.TableColumn, error) {

	result := make([]models.TableColumn, 0)

	finalTableName := fmt.Sprintf("z_%d_%s", pid, table)

	rows, err := d.sess.SQL().Query(fmt.Sprintf(`SELECT * FROM pragma_table_info('%s')`, finalTableName))
	if err != nil {
		return nil, err
	}

	for rows.Next() {
		var cid int64
		var name string
		var typeName string
		var notNull int
		var defaultValue *string
		var primaryKey int
		err := rows.Scan(&cid, &name, &typeName, &notNull, &defaultValue, &primaryKey)
		if err != nil {
			return nil, err
		}
		var defaultValueActual string
		if defaultValue != nil {
			defaultValueActual = *defaultValue
		}

		result = append(result, models.TableColumn{
			Cid:          cid,
			Name:         name,
			Type:         typeName,
			NotNull:      notNull == 1,
			DefaultValue: defaultValueActual,
			PrimaryKey:   primaryKey == 1,
		})
	}

	return result, nil
}

func (d *DB) ListThirdPartyProjects(userid int64, ptype string) ([]models.Project, error) {

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

	datas := make([]models.Project, 0, len(projs))

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

	data := &models.ProjectUser{
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

type ProjectUserScope struct {
	Scope string `json:"scope" db:"scope,omitempty"`
}

func (d *DB) GetProjectUserScope(userId int64, projectId int64) (string, error) {
	if d.isOwner(userId, projectId) {
		return ScopeOwner, nil
	}

	data := &ProjectUserScope{}
	err := d.projectUserTable().Find(db.Cond{"userId": userId, "projectId": projectId}).One(data)
	if err != nil {
		return "", err
	}

	return data.Scope, nil
}

func (d *DB) ListProjectHooksByUser(uid int64, pid int64) ([]models.ProjectHook, error) {

	if !d.isOwner(uid, pid) {
		return nil, errNotFound
	}

	table := d.projectHooksTable()

	hooks := []models.ProjectHook{}

	err := table.Find(db.Cond{"project_id": pid}).Select(
		"id",
		"event",
		"order_id",
		"runas_user_id",
		"hook_type",
		"hook_code",
	).All(&hooks)
	if err != nil {
		return nil, err
	}

	return hooks, nil

}

func (d *DB) ListProjectHooks(pid int64) ([]models.ProjectHook, error) {

	table := d.projectHooksTable()

	hooks := []models.ProjectHook{}

	err := table.Find(db.Cond{"project_id": pid}).All(&hooks)
	if err != nil {
		return nil, err
	}

	return hooks, nil
}

func (d *DB) ListAllProjectHooks() ([]models.ProjectHook, error) {

	table := d.projectHooksTable()

	hooks := []models.ProjectHook{}

	err := table.Find().Select("id", "event", "project_id").All(&hooks)
	if err != nil {
		return nil, err
	}

	return hooks, nil
}

func (d *DB) AddProjectHook(uid, pid int64, data *models.ProjectHook) (int64, error) {

	if !d.isOwner(uid, pid) {
		pp.Println("@not_isOwner/", uid, pid)
		return 0, errNotFound
	}

	data.ProjectID = pid

	table := d.projectHooksTable()

	res, err := table.Insert(data)
	if err != nil {
		return 0, err
	}

	return res.ID().(int64), nil

}

func (d *DB) RemoveProjectHook(uid, pid int64, hid int64) error {
	if !d.isOwner(uid, pid) {
		return errNotFound
	}

	table := d.projectHooksTable()

	return table.Find(db.Cond{"project_id": pid, "id": hid}).Delete()

}

func (d *DB) UpdateProjectHook(uid, pid, hid int64, data map[string]any) error {

	if !d.isOwner(uid, pid) {
		return errNotFound
	}

	table := d.projectHooksTable()

	return table.Find(db.Cond{"project_id": pid, "id": hid}).Update(&data)
}

func (d *DB) GetProjectHook(uid, pid, hid int64) (*models.ProjectHook, error) {

	if !d.isOwner(uid, pid) {
		return nil, errNotFound
	}

	table := d.projectHooksTable()

	hook := &models.ProjectHook{}

	err := table.Find(db.Cond{"project_id": pid, "id": hid}).One(hook)
	if err != nil {
		return nil, err
	}

	return hook, nil
}

// Plugins

func (d *DB) ListProjectPlugins(uid, pid int64) ([]models.ProjectPlugin, error) {

	if !d.isOwner(uid, pid) {
		return nil, errNotFound
	}

	table := d.projectPluginsTable()

	plugins := []models.ProjectPlugin{}

	err := table.Find(db.Cond{"project_id": pid}).
		Select("id", "name", "ptype", "created_by", "updated_by", "project_id", "created_at", "updated_at").
		All(&plugins)
	if err != nil {
		return nil, err
	}

	return plugins, nil
}

func (d *DB) AddProjectPlugin(uid, pid int64, data *models.ProjectPlugin) (int64, error) {

	if !d.isOwner(uid, pid) {
		return 0, errNotFound
	}

	t := time.Now()

	data.ProjectID = pid
	data.CreatedBy = uid
	data.UpdatedBy = uid
	data.CreatedAt = &t
	data.UpdatedAt = &t

	table := d.projectPluginsTable()

	res, err := table.Insert(data)
	if err != nil {
		return 0, err
	}

	return res.ID().(int64), nil

}

func (d *DB) RemoveProjectPlugin(uid, pid int64, hid int64) error {
	if !d.isOwner(uid, pid) {
		return errNotFound
	}

	table := d.projectPluginsTable()

	return table.Find(db.Cond{"project_id": pid, "id": hid}).Delete()

}

func (d *DB) UpdateProjectPlugin(uid, pid, hid int64, data map[string]any) error {

	if !d.isOwner(uid, pid) {
		return errNotFound
	}

	t := time.Now()

	data["updated_at"] = t
	data["updated_by"] = uid

	table := d.projectPluginsTable()

	return table.Find(db.Cond{"project_id": pid, "id": hid}).Update(&data)
}

func (d *DB) GetProjectPlugin(uid, pid, hid int64) (*models.ProjectPlugin, error) {

	if !d.isOwner(uid, pid) {
		return nil, errNotFound
	}

	table := d.projectPluginsTable()

	plugin := &models.ProjectPlugin{}

	err := table.Find(db.Cond{"project_id": pid, "id": hid}).One(plugin)
	if err != nil {
		return nil, err
	}

	return plugin, nil
}

// project config

func (d *DB) ListProjectConfigs(pid int64) ([]models.ProjectConfig, error) {
	table := d.projectConfigsTable()

	configs := make([]models.ProjectConfig, 0)
	err := table.Find(db.Cond{"project_id": pid}).All(&configs)
	if err != nil {
		return nil, err
	}

	return configs, nil
}

func (d *DB) AddProjectConfig(pid, uid int64, data *models.ProjectConfig) (int64, error) {
	table := d.projectConfigsTable()
	data.ProjectID = pid

	rid, err := table.Insert(data)
	if err != nil {
		return 0, err
	}
	id := rid.ID().(int64)
	return id, nil
}

func (d *DB) UpdateProjectConfig(pid, uid, hid int64, data map[string]any) error {
	table := d.projectConfigsTable()
	return table.Find(db.Cond{"project_id": pid, "id": hid}).Update(&data)
}

func (d *DB) GetProjectConfig(pid, uid, hid int64) (*models.ProjectConfig, error) {
	table := d.projectConfigsTable()
	config := &models.ProjectConfig{}
	err := table.Find(db.Cond{"project_id": pid, "id": hid}).One(config)
	if err != nil {
		return nil, err
	}
	return config, nil
}

func (d *DB) RemoveProjectConfig(pid, uid, hid int64) error {
	table := d.projectConfigsTable()
	return table.Find(db.Cond{"project_id": pid, "id": hid}).Delete()
}

// private

func (d *DB) isOwner(ownerid int64, projId int64) bool {
	exist, err := d.projectTable().Find(db.Cond{"owned_by": ownerid, "id": projId}).Exists()

	if err != nil {
		log.Println("owner check error", err)
		return false
	}

	return exist

}

func (d *DB) projectHooksTable() db.Collection {
	return d.Table("ProjectHooks")
}

func (d *DB) projectPluginsTable() db.Collection {
	return d.Table("ProjectPlugins")
}

func (d *DB) projectConfigsTable() db.Collection {
	return d.Table("ProjectConfigs")
}

func (d *DB) projectTable() db.Collection {
	return d.Table("Projects")
}

func (d *DB) projectUserTable() db.Collection {
	return d.Table("ProjectUsers")
}
