package app

import (
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/bornjre/trunis/backend/xtypes/libx/httpx"
	"github.com/bornjre/trunis/frontend"
	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"
)

func (a *App) bindRoutes(e *gin.Engine) {

	root := e.Group("/z")

	root.POST("/auth/signup/direct", a.signUpDirect)
	root.POST("/auth/signup/invite", a.signUpInvite)
	root.POST("/auth/login", a.login)

	a.apiRoutes(root)
	a.pages(root)

	root.GET("/global.js", func(ctx *gin.Context) {

		ctx.Data(http.StatusOK, httpx.CtypeJS, a.globalJS)
	})

	e.NoRoute(a.noRoute)

	e.GET("/ping", a.ping)

}

func (a *App) apiRoutes(root *gin.RouterGroup) {

	apiv1 := root.Group("/api/v1")

	apiv1.GET("/self", a.accessMiddleware(a.selfInfo))
	apiv1.GET("/self/change_password", a.accessMiddleware(a.selfChangePassword))

	apiv1.GET("/self/users", a.accessMiddleware(a.selfUsers))
	apiv1.POST("/self/users", a.accessMiddleware(a.selfAddUser))
	apiv1.GET("/self/users/:uid", a.accessMiddleware(a.selfGetUser))
	apiv1.POST("/self/users/:uid", a.accessMiddleware(a.selfUpdateUser))
	apiv1.DELETE("/self/users/:uid", a.accessMiddleware(a.selfDeleteUser))

	apiv1.GET("/project", a.accessMiddleware(a.listProjects))
	apiv1.POST("/project", a.accessMiddleware(a.addProject))
	apiv1.POST("/project/:pid", a.accessMiddleware(a.updateProject))
	apiv1.GET("/project/:pid", a.accessMiddleware(a.getProject))
	apiv1.DELETE("/project/:pid", a.accessMiddleware(a.removeProject))

	apiv1.POST("/project/:pid/user", a.accessMiddleware(a.inviteUserToPoject))     // invite
	apiv1.DELETE("/project/:pid/user", a.accessMiddleware(a.removeUserFromPoject)) // remove from project

	// project type

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

}

func (a *App) noRoute(ctx *gin.Context) {

	if strings.HasPrefix(ctx.Request.URL.Path, "/z/") {
		pparts := strings.Split(ctx.Request.URL.Path, "/")

		switch pparts[2] {
		case "portal":
			ctx.Redirect(http.StatusFound, "/z/pages/portal")
			return
		case "auth":
			ctx.Redirect(http.StatusFound, "/z/pages/auth/login")
			return
		default:
			return
		}
	}

}

// during dev we just proxy to dev vite server running otherwise serve files from build folder
func (s *App) pages(z *gin.RouterGroup) {
	rfunc := s.pagesRoutes()

	z.GET("/pages", rfunc)
	z.GET("/pages/*files", rfunc)

}

const NoPreBuildFiles = false

func (s *App) pagesRoutes() gin.HandlerFunc {
	var proxy *httputil.ReverseProxy
	pserver := os.Getenv("FRONTEND_DEV_SERVER")

	if pserver != "" && !NoPreBuildFiles {
		url, err := url.Parse(pserver)
		if err != nil {
			panic(err)
		}
		pp.Println("@using_dev_proxy", pserver)

		proxy = httputil.NewSingleHostReverseProxy(url)
		return func(ctx *gin.Context) {
			pp.Println("[PROXY]", ctx.Request.URL.String())
			proxy.ServeHTTP(ctx.Writer, ctx.Request)
		}

	}
	pp.Println("@not_using_dev_proxy")

	return func(ctx *gin.Context) {

		ppath := strings.TrimSuffix(strings.TrimPrefix(ctx.Request.URL.Path, "/z/pages"), "/")

		if ppath == "" {
			ppath = "index.html"
		}

		pitems := strings.Split(ppath, "/")
		lastpath := pitems[len(pitems)-1]

		if !strings.Contains(lastpath, ".") {
			ppath = ppath + ".html"
		}

		pp.Println("@FILE ==>", ppath)

		out, err := frontend.BuildProd.ReadFile(path.Join("build", ppath))
		if err != nil {
			pp.Println("@open_err", err.Error())
			return
		}

		httpx.WriteFile(ppath, out, ctx)
	}
}

var pingResponse = []byte(`{ "message": "pong" }`)

const JsonContentType = "application/json"

func (a *App) ping(ctx *gin.Context) {
	ctx.Writer.WriteHeader(200)
	ctx.Data(http.StatusOK, JsonContentType, pingResponse)
}
