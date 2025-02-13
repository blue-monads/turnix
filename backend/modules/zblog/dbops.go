package zblog

import (
	"fmt"
	"time"

	"github.com/upper/db/v4"
)

// posts

type PostModal struct {
	ID          int64      `db:"id,omitempty"`
	Title       string     `db:"title"`
	Excerpt     string     `db:"excerpt"`
	Content     string     `db:"content"`
	AuthorID    int64      `db:"author_id"`
	CreatedAt   *time.Time `db:"created_at"`
	UpdatedAt   *time.Time `db:"updated_at"`
	IsPublished bool       `db:"is_published"`
}

func (z *ZBlogModule) dbAddPost(pid, uid int64, data *PostModal) (int64, error) {

	table := z.postsTable(pid)

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

var selectColumns = []string{
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

	err := table.Find(db.Cond{"id >": offset}).Select(selectColumns).All(&datas)
	if err != nil {
		return nil, err
	}

	return datas, nil

}

func (z *ZBlogModule) dbDeletePost(pid, uid, id int64) error {
	return z.postsTable(pid).Find(db.Cond{"id": id}).Delete()
}

// private

func (z *ZBlogModule) postsTable(pid int64) db.Collection {
	return z.db.Table(postsTableName(pid))
}

func postsTableName(pid int64) string {
	return fmt.Sprintf("z_%d_Posts", pid)
}
