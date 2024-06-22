package app

import (
	"fmt"
	"log"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/bornjre/turnix/backend/utils/libx/httpx"
	"github.com/bornjre/turnix/backend/xtypes/xproject"
	"github.com/bornjre/turnix/frontend"
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

	projectRoute := e.Group("/z/project")

	for _, pdef := range a.ptypeDefs {
		subLogger := a.rootLogger.With().Str("ptype", pdef.Slug).Logger()

		pp.Println("@mounting_ptype", pdef.Slug)

		if pdef.Builder == nil {
			continue
		}

		proj, err := pdef.Builder(xproject.BuilderOption{
			App:         a,
			Logger:      subLogger,
			RouterGroup: projectRoute.Group(fmt.Sprintf("/%s", pdef.Slug)),
		})
		if err != nil {
			log.Fatal(err)
		}

		a.projects[pdef.Slug] = proj

	}

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

	apiv1.GET("/project_types", a.accessMiddleware(a.ListProjectTypes))
	apiv1.GET("/project_types/:ptype/form", a.accessMiddleware(a.GetProjectTypeForm))
	apiv1.GET("/project_types/:ptype", a.accessMiddleware(a.GetProjectType))

	// project hook

	apiv1.GET("/project/:pid/hook", a.accessMiddleware(a.listProjectHooks))
	apiv1.POST("/project/:pid/hook", a.accessMiddleware(a.addProjectHook))
	apiv1.GET("/project/:pid/hook/:id", a.accessMiddleware(a.getProjectHook))
	apiv1.POST("/project/:pid/hook/:id", a.accessMiddleware(a.updateProjectHook))
	apiv1.DELETE("/project/:pid/hook/:id", a.accessMiddleware(a.removeProjectHook))

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
