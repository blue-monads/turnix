package server

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/blue-monads/turnix/backend/app/server/assets"
	output "github.com/blue-monads/turnix/frontend"
	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"
)

// during dev we just proxy to dev vite server running otherwise serve files from build folder
func (s *Server) pages(z *gin.RouterGroup) {
	rfunc := assets.PagesRoutesServer()
	extFunc := s.externalAssets()

	z.GET("/pages", rfunc)
	z.GET("/pages/*files", rfunc)
	z.GET("/x/:pname/*files", extFunc)
	z.GET("/x/:pname", extFunc)
	z.GET("/lib/*file", func(ctx *gin.Context) {

		pp.Println("@lib/1")

		file := ctx.Param("file")

		file = strings.TrimPrefix(file, "/")

		fout, err := output.BuildProd.Open(path.Join("output", file))
		if err != nil {
			pp.Println("@lib/2", err.Error())
			return
		}

		pp.Println("@lib/3")

		defer fout.Close()

		io.Copy(ctx.Writer, fout)

	})

}

func (s *Server) externalAssets() gin.HandlerFunc {

	proxyAddrs := map[string]*httputil.ReverseProxy{}

	if s.devMode {
		for pname := range s.projects {
			addr := os.Getenv(fmt.Sprintf("TURNIX_DEV_%s_SERVER", strings.ToUpper(pname)))
			if addr == "" {
				continue
			}

			url, err := url.Parse(addr)
			if err != nil {
				panic(err)
			}

			proxy := httputil.NewSingleHostReverseProxy(url)
			proxyAddrs[pname] = proxy
		}
	}

	return func(ctx *gin.Context) {

		pp.Println("@ext/1")

		pname := ctx.Param("pname")

		if s.devMode {
			pp.Println("@ext/2")
			proxy := proxyAddrs[pname]
			if proxy != nil {
				pp.Println("@ext/3")
				proxy.ServeHTTP(ctx.Writer, ctx.Request)
				return
			}
			pp.Println("@ext/4")
		}

		pp.Println("@ext/5")

		def := s.projects[pname]
		if def == nil {
			pp.Println("@ext/6")
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		pp.Println("@ext/7")

		if def.OnPageRequest != nil {
			def.OnPageRequest(ctx)
			return
		}

		pp.Println("@ext/8", ctx.Request.URL.Path)

		prefix := fmt.Sprintf("/z/x/%s", pname)
		ppath := strings.TrimPrefix(ctx.Request.URL.Path, prefix)
		ppath = strings.TrimSuffix(ppath, "/")
		ppath = strings.TrimPrefix(ppath, "/")

		pp.Println("@ext/9")

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

}
