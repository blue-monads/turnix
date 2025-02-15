package zblog

import (
	"fmt"
	"regexp"
	"strings"
	"time"

	"github.com/k0kubun/pp"
	"github.com/upper/db/v4"
)

// posts

type PostModal struct {
	ID          int64      `db:"id,omitempty" json:"id"`
	Title       string     `db:"title" json:"title"`
	Slug        string     `db:"slug" json:"slug"`
	Excerpt     string     `db:"excerpt" json:"excerpt"`
	Content     string     `db:"content" json:"content"`
	AuthorID    int64      `db:"author_id" json:"author_id"`
	CreatedAt   *time.Time `db:"created_at" json:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at" json:"updated_at"`
	IsPublished bool       `db:"is_published" json:"is_published"`
}

var slugRegex = regexp.MustCompile(`^[a-z0-9-]+$`)

func slugify(s string) string {
	s = strings.ReplaceAll(s, " ", "-")
	s = strings.ToLower(s)

	return slugRegex.ReplaceAllString(s, "")
}

func (z *ZBlogModule) dbAddPost(pid, uid int64, data *PostModal) (int64, error) {

	table := z.postsTable(pid)

	t := time.Now()
	data.CreatedAt = &t
	data.AuthorID = uid
	data.UpdatedAt = &t

	if data.Slug == "" {
		data.Slug = slugify(data.Title)
	}

	if len(data.Content) > 100 {
		data.Excerpt = data.Content[:100]
	} else {
		data.Excerpt = data.Content
	}

	pp.Println("@data", data)

	r, err := table.Insert(data)
	if err != nil {
		return 0, err
	}

	return r.ID().(int64), nil

}

func (z *ZBlogModule) dbUpdatePost(pid, uid, id int64, data map[string]any) error {

	t := time.Now()
	data["updated_at"] = &t

	content, ok := data["content"]
	if ok {

		if len(content.(string)) > 100 {
			data["excerpt"] = content.(string)[:100]
		} else {
			data["excerpt"] = content.(string)
		}
	}

	return z.postsTable(pid).Find(db.Cond{"id": id}).Update(data)
}

func (z *ZBlogModule) dbGetPost(pid, uid, aid int64) (*PostModal, error) {

	data := &PostModal{}
	table := z.postsTable(pid)

	err := table.Find(db.Cond{"id": aid}).One(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

var selectColumns = []any{
	"id",
	"title",
	"excerpt",
	"author_id",
	"created_at",
	"updated_at",
	"is_published",
}

func (z *ZBlogModule) dbListPost(pid, uid, offset int64) ([]PostModal, error) {

	datas := make([]PostModal, 0)
	table := z.postsTable(pid)

	err := table.Find(db.Cond{"id >": offset}).Select(selectColumns...).All(&datas)
	if err != nil {
		return nil, err
	}

	return datas, nil

}

func (z *ZBlogModule) dbDeletePost(pid, uid, id int64) error {
	return z.postsTable(pid).Find(db.Cond{"id": id}).Delete()
}

// site

type SiteModal struct {
	ID             int64      `db:"id,omitempty" json:"id"`
	ApiKey         string     `db:"api_key" json:"api_key"`
	Provider       string     `db:"provider" json:"provider"`
	Domain         string     `db:"domain" json:"domain"`
	BasePath       string     `db:"base_path" json:"base_path"`
	CreatedAt      *time.Time `db:"created_at" json:"created_at"`
	LastDeployedAt *time.Time `db:"last_deployed_at" json:"last_deployed_at"`
	DeployWebhook  string     `db:"deploy_webhook" json:"deploy_webhook"`
	DeployBranch   string     `db:"deploy_branch" json:"deploy_branch"`
}

func (z *ZBlogModule) dbAddSite(pid, uid int64, data *SiteModal) (int64, error) {

	table := z.sitesTable(pid)

	t := time.Now()

	data.CreatedAt = &t

	r, err := table.Insert(data)
	if err != nil {
		return 0, err
	}

	return r.ID().(int64), nil

}

func (z *ZBlogModule) dbUpdateSite(pid, uid, id int64, data map[string]any) error {

	return z.sitesTable(pid).Find(db.Cond{"id": id}).Update(data)
}

func (z *ZBlogModule) dbGetSite(pid, uid, id int64) (*SiteModal, error) {

	data := &SiteModal{}

	err := z.sitesTable(pid).Find(db.Cond{"id": id}).One(data)
	if err != nil {
		return nil, err
	}

	return data, nil
}

func (z *ZBlogModule) dbListSite(pid, uid int64) ([]SiteModal, error) {

	datas := make([]SiteModal, 0)

	err := z.sitesTable(pid).Find().All(&datas)
	if err != nil {
		return nil, err
	}

	return datas, nil
}

func (z *ZBlogModule) dbDeleteSite(pid, uid, id int64) error {
	return z.postsTable(pid).Find(db.Cond{"id": id}).Delete()
}

// private

func (z *ZBlogModule) sitesTable(pid int64) db.Collection {
	return z.db.Table(sitesTableName(pid))
}

func sitesTableName(pid int64) string {
	return fmt.Sprintf("z_%d_Sites", pid)
}

func (z *ZBlogModule) postsTable(pid int64) db.Collection {
	return z.db.Table(postsTableName(pid))
}

func postsTableName(pid int64) string {
	return fmt.Sprintf("z_%d_Posts", pid)
}
