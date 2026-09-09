package cloudy

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func placeholderHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello, World!",
	})
}

func (a *CloudyApp) registerBaseRouter(router *gin.Engine) {

	rgroup := router.Group("/zz/cloudy/")

	rgroup.GET("/", placeholderHandler)
	rgroup.GET("/health", placeholderHandler)
	rgroup.GET("/sign-up", placeholderHandler)
	rgroup.POST("/sign-up", placeholderHandler)

}

func (a *CloudyApp) redirrectToApp(c *gin.Context) {}

func (a *CloudyApp) loadApp(c *gin.Context) {

}
