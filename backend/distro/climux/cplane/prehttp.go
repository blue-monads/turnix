package cplane

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path"
	"strings"
	"time"

	"github.com/blue-monads/turnix/backend/app/server"
	"github.com/blue-monads/turnix/backend/app/server/assets"
	"github.com/blue-monads/turnix/backend/extras/mesh/lpweb"
	xutils "github.com/blue-monads/turnix/backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"
)

func (e *Cplane) runPreHttpServer() {

	engine := gin.Default()

	rfunc := assets.PagesRoutesServer()

	engine.GET("/z/global.js", func(ctx *gin.Context) {})
	engine.GET("/z/pages", rfunc)
	engine.GET("/z/pages/*files", rfunc)

	eapi := engine.Group("/z/eapi", e.Auth)

	eapi.GET("/status", e.statusPage)
	eapi.POST("/start", e.startInstance)
	eapi.POST("/navigate/local", e.handleNavigateLocal)
	eapi.POST("/navigate/remote", e.handleNavigateRemote)

	port, err := xutils.GetFreePort()
	if err != nil {
		panic(err)
	}

	e.port = port
	e.engine = engine

	err = http.ListenAndServe(fmt.Sprintf(":%d", port), e)
	if err != nil {
		panic(err)
	}

}

func (e *Cplane) Auth(ctx *gin.Context) {
	apiToken := ctx.GetHeader("X-API-TOKEN")
	if apiToken != e.apiToken {
		ctx.JSON(401, gin.H{
			"error": "unauthorized",
		})
		return
	}

	ctx.Next()
}

func (e *Cplane) handleNavigateLocal(ctx *gin.Context) {

	url := ctx.PostForm("url")
	pp.Println("@navigateLocal", url)

	if e.browser != nil {
		e.browser.Navigate(url)
	}

	ctx.JSON(200, gin.H{
		"success": true,
	})

}

func (e *Cplane) handleNavigateRemote(ctx *gin.Context) {

	url := ctx.PostForm("url")
	routeToPages := ctx.PostForm("routeToPages")
	pp.Println("@navigateRemote", url, routeToPages)

	e.remoteNavigate(url, routeToPages == "true")
	ctx.JSON(200, gin.H{
		"success": true,
	})

}

func (e *Cplane) startInstance(ctx *gin.Context) {

	e.cLock.Lock()
	defer e.cLock.Unlock()

	if e.cmd != nil {
		ctx.JSON(400, gin.H{
			"error": "already running",
		})
		return
	}

	err := e.configurer.InitPath()
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	selfbinary, err := os.Executable()
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
	}

	cmd := exec.Command(selfbinary, "node", "start")
	cmd.Env = append(cmd.Env, os.Environ()...)
	cmd.Env = append(cmd.Env, fmt.Sprintf("TURNIX_MASTER_SECRET=%s", e.config.MasterKey))
	cmd.Env = append(cmd.Env, fmt.Sprintf("TURNIX_HTTP_PORT=:%d", e.config.HttpServerPort))

	cmd.Env = append(cmd.Env, fmt.Sprintf("TURNIX_LOCAL_SOCKET=%s", path.Join(e.configurer.BasePath, e.config.LocalSocket)))
	cmd.Env = append(cmd.Env, fmt.Sprintf("TURNIX_DATABASE_FILE=%s", path.Join(e.configurer.BasePath, e.config.DatabaseFile)))

	cmd.Stdout = os.Stdout
	cmd.Stderr = os.Stderr

	err = cmd.Start()
	if err != nil {
		ctx.JSON(400, gin.H{
			"error": err.Error(),
		})
		return
	}

	e.cmd = cmd

	msg, err := e.getStatus()
	if err != nil {
		msg, err = e.getStatus()
		if err != nil {
			time.Sleep(time.Second * 5)
		}
	}

	ctx.JSON(200, gin.H{
		"pid":  cmd.Process.Pid,
		"port": msg.Port,
	})

}

func (e *Cplane) statusPage(ctx *gin.Context) {

	meshAddr := ""

	counter := 0

	for {
		pp.Println("@status/1", counter)
		if e.mesh == nil {
			pp.Println("@status/2/sleep", counter)
			time.Sleep(time.Second * 2)
		} else {
			pp.Println("@status/3/break", counter)
			break
		}

		if counter > 10 {
			pp.Println("@status/4/break", counter)
			break
		}

		pp.Println("@status/5", counter)

		counter++

	}

	if e.mesh != nil {
		meshAddr = e.mesh.GetMeshAddr()
	}

	msg, err := e.getStatus()
	if err != nil {
		ctx.JSON(200, gin.H{
			"is_running": false,
			"status":     err.Error(),
			"mesh_addr":  meshAddr,
		})
		return
	}

	pp.Println("status/4")

	ctx.JSON(200, gin.H{
		"is_running":  true,
		"port":        msg.Port,
		"status":      "ok",
		"working_dir": e.configurer.BasePath,
		"mesh_addr":   meshAddr,
	})

}

func (e *Cplane) getStatus() (*server.LocalStatus, error) {

	localSocket := path.Join(e.configurer.BasePath, e.config.LocalSocket)

	conn, err := net.Dial("unix", localSocket)
	pp.Println("status/1.5")

	if err != nil {
		return nil, err
	}

	pp.Println("status/2")

	defer conn.Close()

	out, err := io.ReadAll(conn)
	if err != nil {
		return nil, err
	}

	pp.Println("status/3")

	msg := server.LocalStatus{}

	err = json.Unmarshal(out, &msg)
	if err != nil {
		return nil, err
	}

	return &msg, nil

}

func (e *Cplane) startMesh() {
	e.meshBuildLock.Lock()
	defer e.meshBuildLock.Unlock()

	if e.mesh != nil {
		return
	}

	time.Sleep(time.Second * 5)

	pp.Println("@e.config.MasterKey", e.config)

	e.mesh = lpweb.New(e.config.MasterKey, e.config.MeshPort, e.config.HttpServerPort)

}

func (e *Cplane) ServeHTTP(rw http.ResponseWriter, r *http.Request) {

	host := r.Host

	pp.Println("@EbrowserApp/ServeHTTP", r.Host, r.URL)

	if strings.Contains(host, "-lpweb.localhost:") {

		r.URL.Host = strings.Replace(host, "-lpweb", "", 1)
		url, err := r.URL.Parse(r.URL.String())
		if err != nil {
			panic(err)
		}

		r.URL = url

		if r.Header.Get("Upgrade") == "websocket" {
			e.mesh.HandleWsIn(r, rw)
		} else {
			e.mesh.HandleHttpIn(r, rw)
		}

		return
	}

	e.engine.ServeHTTP(rw, r)
}
