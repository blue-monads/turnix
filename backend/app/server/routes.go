package server

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/bornjre/turnix/backend/app/server/assets"
	"github.com/bornjre/turnix/backend/utils/libx/httpx"
	"github.com/gin-gonic/gin"
)

func (a *Server) bindRoutes(e *gin.Engine) {

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

	for _, instance := range a.projects {
		if instance.OnAPIMount == nil {
			continue
		}

		instance.OnAPIMount(projectRoute.Group(fmt.Sprintf("/%s", instance.Slug)))
	}

	e.NoRoute(a.noRoute)

	root.Any("/projects/:ptype", func(ctx *gin.Context) {
		instance := a.projects[ctx.Param("ptype")]

		if instance.OnProjectRequest == nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		instance.OnProjectRequest(ctx)

	})

	e.GET("/ping", a.ping)

}

func (a *Server) apiRoutes(root *gin.RouterGroup) {

	apiv1 := root.Group("/api/v1")

	apiv1.GET("/self", a.accessMiddleware(a.selfInfo))
	apiv1.GET("/self/change_password", a.accessMiddleware(a.selfChangePassword))

	apiv1.GET("/self/users", a.accessMiddleware(a.selfUsers))
	apiv1.POST("/self/users", a.accessMiddleware(a.selfAddUser))
	apiv1.GET("/self/users/:uid", a.accessMiddleware(a.selfGetUser))
	apiv1.POST("/self/users/:uid", a.accessMiddleware(a.selfUpdateUser))
	apiv1.DELETE("/self/users/:uid", a.accessMiddleware(a.selfDeleteUser))

	apiv1.GET("/self/messages", a.accessMiddleware(a.listUserMessages))

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

	// project files

	apiv1.GET("/project/:pid/files", a.accessMiddleware(a.listProjectFiles))
	apiv1.POST("/project/:pid/files", a.accessMiddleware(a.addProjectFile))
	apiv1.PUT("/project/:pid/files", a.accessMiddleware(a.addProjectFolder))

	apiv1.GET("/project/:pid/files/:id", (a.getProjectFile))
	apiv1.DELETE("/project/:pid/files/:id", a.accessMiddleware(a.removeProjectFile))

	apiv1.POST("/project/:pid/sqlexec", a.accessMiddleware(a.runProjectSQL))

	// project plugins

	apiv1.GET("/project/:pid/plugins", a.accessMiddleware(a.listProjectPlugins))
	apiv1.POST("/project/:pid/plugins", a.accessMiddleware(a.addProjectPlugin))
	apiv1.DELETE("/project/:pid/plugins/:id", a.accessMiddleware(a.removeProjectPlugin))
	apiv1.POST("/project/:pid/plugins/:id", a.accessMiddleware(a.updateProjectPlugin))
	apiv1.GET("/project/:pid/plugins/:id", a.accessMiddleware(a.getProjectPlugin))

	// self files

	apiv1.GET("/self/files", a.accessMiddleware(a.listSelfFiles))
	apiv1.POST("/self/files", a.accessMiddleware(a.addSelfFile))
	apiv1.PUT("/self/files", a.accessMiddleware(a.addSelfFolder))

	apiv1.GET("/self/files/:id", (a.getSelfFile))
	apiv1.DELETE("/self/files/:id", a.accessMiddleware(a.removeSelfFile))
	apiv1.POST("/self/messages/:uid", a.accessMiddleware(a.messageUser))

	apiv1.GET("/self/files/:id/shares", a.accessMiddleware(a.listSelfFileShares))
	apiv1.POST("/self/files/:id/shares", a.accessMiddleware(a.addSelfFileShare))
	apiv1.DELETE("/self/files/:id/shares/:id", a.accessMiddleware(a.deleteSelfFileShare))

	// user profile
	apiv1.GET("/user/:uid", a.accessMiddleware(a.userProfile))
	apiv1.GET("/file/shared/:file", a.getSharedFile)
	apiv1.GET("/file/shortKey/:shortkey", a.GetFileWithShortKey)
	apiv1.POST("/file/:fid/shortkey", a.accessMiddleware(a.GetFileShortKey))

}

func (s *Server) noRoute(ctx *gin.Context) {

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
func (s *Server) pages(z *gin.RouterGroup) {
	rfunc := assets.PagesRoutesServer()

	z.GET("/pages", rfunc)
	z.GET("/pages/*files", rfunc)

}

var pingResponse = []byte(`{ "message": "pong" }`)

const JsonContentType = "application/json"

func (a *Server) ping(ctx *gin.Context) {
	ctx.Writer.WriteHeader(200)
	ctx.Data(http.StatusOK, JsonContentType, pingResponse)
}
