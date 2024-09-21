package ebrowser

import (
	"fmt"
	"net/http"

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
