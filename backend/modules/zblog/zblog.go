package zblog

import (
	"fmt"
	"strconv"

	_ "embed"

	"github.com/blue-monads/turnix/backend/modules"
	"github.com/blue-monads/turnix/backend/services/database"
	"github.com/blue-monads/turnix/backend/xtypes"
	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"
	"github.com/upper/db/v4"
)

//go:embed zblog.sql
var Schema string

type ZBlogModule struct {
	app  xtypes.App
	sess db.Session
	db   *database.DB
}

func (z *ZBlogModule) Init(pid int64) error {

	schema := modules.ParamaterizedSchema(Schema, pid)

	pp.Println("@parameterizedSchema")
	fmt.Println(schema)

	err := z.db.RunDDL(schema)
	if err != nil {
		pp.Println("@err", err.Error())
	}

	return err

}

func (z *ZBlogModule) Register(rg *gin.RouterGroup) {

	w := z.app.AsApiAction

	api := rg.Group("/api")

	api.POST("/post", w("addPosts", z.addPosts))
	api.GET("/post", w("listPosts", z.listPosts))
	api.GET("/post/:id", w("getPosts", z.getPosts))
	api.PUT("/post/:id", w("updatePosts", z.updatePosts))
	api.DELETE("/post/:id", w("deletePosts", z.deletePosts))

}

func (z *ZBlogModule) listPosts(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()
	offset, _ := strconv.ParseInt(ctx.Http.Query("offset"), 10, 64)
	return z.dbListPost(pid, ctx.Claim.UserId, offset)
}

func (z *ZBlogModule) getPosts(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()
	id := ctx.ParamInt64("id")
	return z.dbGetPost(pid, ctx.Claim.UserId, id)
}

func (z *ZBlogModule) addPosts(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()

	data := &PostModal{}
	err := ctx.Http.Bind(data)
	if err != nil {
		return nil, err
	}

	return z.dbAddPost(pid, ctx.Claim.UserId, data)
}

func (z *ZBlogModule) updatePosts(ctx xtypes.ContextPlus) (any, error) {
	pid := ctx.ProjectId()
	id := ctx.ParamInt64("id")

	data := make(map[string]any)
	err := ctx.Http.Bind(data)
	if err != nil {
		return nil, err
	}

	return nil, z.dbUpdatePost(pid, ctx.Claim.UserId, id, data)
}

func (z *ZBlogModule) deletePosts(ctx xtypes.ContextPlus) (any, error) {

	pid := ctx.ProjectId()
	id := ctx.ParamInt64("id")

	return nil, z.dbDeletePost(pid, ctx.Claim.UserId, id)
}
