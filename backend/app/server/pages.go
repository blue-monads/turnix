package server

import (
	"fmt"
	"io"
	"net/http"
	"net/http/httputil"
	"net/url"
	"os"
	"strings"

	"github.com/bornjre/turnix/backend/app/server/assets"
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

		pname := ctx.Param("pname")

		if s.devMode {
			proxy := proxyAddrs[pname]
			if proxy != nil {
				proxy.ServeHTTP(ctx.Writer, ctx.Request)
				return
			}
		}

		def := s.projects[pname]
		if def == nil {
			ctx.AbortWithStatus(http.StatusNotFound)
			return
		}

		if def.OnPageRequest != nil {
			def.OnPageRequest(ctx)
			return
		}

		ppath := strings.TrimPrefix(ctx.Request.URL.Path, fmt.Sprintf("/z/x/%s/", pname))
		ppath = strings.TrimSuffix(ppath, "/")
		ppath = strings.TrimPrefix(ppath, "/")

		if ppath == "" {
			ppath = "index.html"
		}

		file, err := def.AssetData.Open(ppath)
		if err != nil {
			pp.Println("@open_err", err.Error())
			return
		}

		defer file.Close()

		io.Copy(ctx.Writer, file)
	}

}
