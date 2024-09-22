package ebrowser

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"

	"github.com/bornjre/turnix/backend/app/server"
	"github.com/bornjre/turnix/backend/app/server/assets"
	xutils "github.com/bornjre/turnix/backend/utils"
	"github.com/gin-gonic/gin"
)

func (e *EbrowserApp) runPreHttpServer() {

	engine := gin.Default()

	rfunc := assets.PagesRoutesServer()

	engine.GET("/z/global.js", func(ctx *gin.Context) {

	})
	engine.GET("/z/pages", rfunc)
	engine.GET("/z/pages/*files", rfunc)

	eapi := engine.Group("/z/eapi")

	eapi.GET("/status", e.status)

	port, err := xutils.GetFreePort()
	if err != nil {
		panic(err)
	}

	e.port = port

	err = http.ListenAndServe(fmt.Sprintf(":%d", port), engine)
	if err != nil {
		panic(err)
	}

}

func (e *EbrowserApp) status(ctx *gin.Context) {
	conn, err := net.Dial("unix", e.config.LocalSocket)
	if err != nil {
		ctx.JSON(200, gin.H{
			"is_running": false,
			"status":     err.Error(),
		})
		return
	}

	defer conn.Close()

	out, err := io.ReadAll(conn)
	if err != nil {
		ctx.JSON(200, gin.H{
			"is_running": false,
			"status":     err.Error(),
		})
		return
	}

	msg := server.LocalStatus{}

	err = json.Unmarshal(out, &msg)
	if err != nil {
		ctx.JSON(200, gin.H{
			"is_running": false,
			"status":     err.Error(),
		})
		return
	}

	ctx.JSON(200, gin.H{
		"is_running": true,
		"port":       e.port,
		"status":     "ok",
	})

}
