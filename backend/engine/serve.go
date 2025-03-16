package engine

import (
	"fmt"
	"io"
	"net/http"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"
)

func (e *Engine) MountProjectType(g *gin.RouterGroup) {

	e.pLock.RLock()
	defer e.pLock.RUnlock()

	for _, def := range e.projects {
		if def.def.OnAPIMount == nil {
			continue
		}

		def.def.OnAPIMount(g.Group(fmt.Sprintf("/%s", def.def.Slug)))
	}

}

func (e *Engine) ServeProjectType(name string, ctx *gin.Context) {
	e.pLock.RLock()
	def := e.projects[name]
	e.pLock.RUnlock()

	def.def.OnProjectRequest(ctx)
}

func (e *Engine) ServeProjectTypeFile(name string, ctx *gin.Context) {

	e.pLock.RLock()
	def := e.projects[name]
	e.pLock.RUnlock()

	if def == nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	if def.def.OnPageRequest != nil {
		def.def.OnPageRequest(ctx)
		return
	}

	ppath := ctx.Param("file")
	ppath = strings.TrimPrefix(ppath, "/")
	ppath = strings.TrimSuffix(ppath, "/")

	if ppath == "" {
		ppath = "index.html"
	}

	if def.def.AssetData == nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	if ppath == "" {
		ppath = "index.html"
	}

	pp.Println("@ext/10", ppath)

	if def.def.AssetData == nil {
		pp.Println("@ext/11", ppath)
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	file, err := def.def.AssetData.Open(path.Join(def.def.AssetDataPrefix, ppath))
	if err != nil {
		pp.Println("@open_err", err.Error())
		return
	}

	defer file.Close()

	io.Copy(ctx.Writer, file)

}
