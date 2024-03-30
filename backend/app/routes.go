package app

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func (a *App) bindRoutes(e *gin.Engine) {

	e.POST("/auth/signup/direct", a.signUpDirect)
	e.POST("/auth/signup/invite", a.signUpInvite)
	e.POST("/auth/login", a.login)

	e.GET("/ping", a.ping)

	apiv1 := e.Group("/api/v1")

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

var pingResponse = []byte(`{ "message": "pong" }`)

const JsonContentType = "application/json"

func (a *App) ping(ctx *gin.Context) {
	ctx.Writer.WriteHeader(200)
	ctx.Data(http.StatusOK, JsonContentType, pingResponse)
}
