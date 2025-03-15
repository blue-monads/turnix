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

func (e *Engine) MountSpaces(g *gin.RouterGroup) {

	e.pLock.RLock()
	defer e.pLock.RUnlock()

	for _, def := range e.projects {
		if def.OnAPIMount == nil {
			continue
		}

		def.OnAPIMount(g.Group(fmt.Sprintf("/%s", def.Slug)))
	}

}

func (e *Engine) ServeSpace(name string, ctx *gin.Context) {
	e.pLock.RLock()
	def := e.projects[name]
	e.pLock.RUnlock()

	def.OnProjectRequest(ctx)
}

func (e *Engine) ServeSpaceFile(name string, ctx *gin.Context) {

	e.pLock.RLock()
	def := e.projects[name]
	e.pLock.RUnlock()

	if def == nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	if def.OnPageRequest != nil {
		def.OnPageRequest(ctx)
		return
	}

	ppath := ctx.Param("file")
	ppath = strings.TrimPrefix(ppath, "/")
	ppath = strings.TrimSuffix(ppath, "/")

	if ppath == "" {
		ppath = "index.html"
	}

	if def.AssetData == nil {
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	if ppath == "" {
		ppath = "index.html"
	}

	pp.Println("@ext/10", ppath)

	if def.AssetData == nil {
		pp.Println("@ext/11", ppath)
		ctx.AbortWithStatus(http.StatusNotFound)
		return
	}

	file, err := def.AssetData.Open(path.Join(def.AssetDataPrefix, ppath))
	if err != nil {
		pp.Println("@open_err", err.Error())
		return
	}

	defer file.Close()

	io.Copy(ctx.Writer, file)

}
