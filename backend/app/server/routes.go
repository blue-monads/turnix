package server

import (
	"net/http"
	"strings"

	"github.com/blue-monads/turnix/backend/utils/libx/httpx"
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
		ctx.Data(http.StatusOK, httpx.CtypeJS, []byte(`console.log("global.js")`))
	})

	projectRoute := e.Group("/z/project")

	a.engine.MountSpaces(projectRoute)

	e.NoRoute(a.noRoute)

	root.Any("/projects/:ptype", func(ctx *gin.Context) {
		ptype := ctx.Param("ptype")
		a.engine.ServeSpace(ptype, ctx)
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
	apiv1.GET("/self/self", a.accessMiddleware(a.getSelfSelf))

	apiv1.GET("/self/messages", a.accessMiddleware(a.listUserMessages))

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

	apiv1.GET("/file/shared/:file", a.getSharedFile)
	apiv1.POST("/file/shared/:file", a.accessMiddleware(a.sharedFile))
	apiv1.DELETE("/file/shared/:file", a.accessMiddleware(a.deleteShareFile))

	apiv1.GET("/file/shortKey/:shortkey", a.GetFileWithShortKey)
	apiv1.POST("/file/:fid/shortkey", a.accessMiddleware(a.GetFileShortKey))

	// user profile
	apiv1.GET("/user/:uid", a.accessMiddleware(a.userProfile))

	// project

	apiv1.GET("/project", a.accessMiddleware(a.listProjects))
	apiv1.POST("/project", a.accessMiddleware(a.addProject))
	apiv1.POST("/project/:pid", a.accessMiddleware(a.updateProject))
	apiv1.GET("/project/:pid", a.accessMiddleware(a.getProject))
	apiv1.DELETE("/project/:pid", a.accessMiddleware(a.removeProject))

	apiv1.POST("/project/:pid/user", a.accessMiddleware(a.inviteUserToPoject))     // invite
	apiv1.DELETE("/project/:pid/user", a.accessMiddleware(a.removeUserFromPoject)) // remove from project

	apiv1.POST("/project_type_install", a.accessMiddleware(a.installProjectType))

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
	apiv1.POST("/project/:pid/sqlexec2", a.accessMiddleware(a.runProjectSQL2))
	apiv1.GET("/project/:pid/tables", a.accessMiddleware(a.listProjectTables))
	apiv1.GET("/project/:pid/tables/:table/columns", a.accessMiddleware(a.listProjectTableColumns))
	apiv1.POST("/project/:pid/autoquery", a.accessMiddleware(a.autoQueryProjectTable))

	// project plugins

	apiv1.GET("/project/:pid/plugins", a.accessMiddleware(a.listProjectPlugins))
	apiv1.POST("/project/:pid/plugins", a.accessMiddleware(a.addProjectPlugin))
	apiv1.DELETE("/project/:pid/plugins/:id", a.accessMiddleware(a.removeProjectPlugin))
	apiv1.POST("/project/:pid/plugins/:id", a.accessMiddleware(a.updateProjectPlugin))
	apiv1.GET("/project/:pid/plugins/:id", a.accessMiddleware(a.getProjectPlugin))

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

var pingResponse = []byte(`{ "message": "pong" }`)

const JsonContentType = "application/json"

func (a *Server) ping(ctx *gin.Context) {
	ctx.Writer.WriteHeader(200)
	ctx.Data(http.StatusOK, JsonContentType, pingResponse)
}
