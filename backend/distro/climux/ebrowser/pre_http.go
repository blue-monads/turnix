package ebrowser

import (
	"fmt"
	"net/http"

	"github.com/bornjre/turnix/backend/app/server/assets"
	xutils "github.com/bornjre/turnix/backend/utils"
	"github.com/gin-gonic/gin"
)

func (e *EbrowserApp) runPreHttpServer() {

	prort, err := xutils.GetFreePort()
	if err != nil {
		panic(err)
	}

	engine := gin.Default()

	rfunc := assets.PagesRoutesServer()

	engine.GET("/z/pages", rfunc)
	engine.GET("/z/pages/*files", rfunc)

	err = http.ListenAndServe(fmt.Sprintf(":%d", prort), engine)
	if err != nil {
		panic(err)
	}

}
