package zblog

import (
	"github.com/blue-monads/turnix/backend/registry"
	"github.com/blue-monads/turnix/backend/services/database"
	"github.com/blue-monads/turnix/backend/xtypes/xproject"
)

var perminssions = []string{"read", "write", "read/write"}

func init() {
	registry.Register("zblog", New)
}

func New(opt xproject.BuilderOption) (*xproject.Defination, error) {

	db := opt.App.GetDatabase()

	d := db.(*database.DB)

	mod := &ZBlogModule{
		sess: d.GetSession(),
		db:   d,
		app:  opt.App,
	}

	def := &xproject.Defination{
		Name:         "ZBlog",
		Slug:         "zblog",
		Info:         "Simple blogging platform",
		Icon:         "chat-bubble-bottom-center-text",
		Perminssions: perminssions,
		OnInit:       mod.Init,
		OnAPIMount:   mod.Register,
	}

	return def, nil
}

/*

artifact.zip
 - index.html
 - posts/[id-slug].html
 - tags/[tag].html
 - assets
 	   - images
 - theme_assets

theme.zip
 - theme.json
 - layout.tpl.html
 - posts.tpl.html
 - post.tpl.html
 - tags.tpl.html
 - assets
	   - css
	   - js
	   - img
	   - fonts

*/
