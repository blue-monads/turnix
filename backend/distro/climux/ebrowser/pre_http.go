package ebrowser

import (
	"encoding/json"
	"fmt"
	"io"
	"net"
	"net/http"
	"os"
	"os/exec"
	"path"
	"time"

	"github.com/bornjre/turnix/backend/app/server"
	"github.com/bornjre/turnix/backend/app/server/assets"
	xutils "github.com/bornjre/turnix/backend/utils"
	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"
)

func (e *EbrowserApp) runPreHttpServer() {

	engine := gin.Default()

	rfunc := assets.PagesRoutesServer()

	engine.GET("/z/global.js", func(ctx *gin.Context) {

	})
	engine.GET("/z/pages", rfunc)
	engine.GET("/z/pages/*files", rfunc)

	eapi := engine.Group("/z/eapi")

	eapi.GET("/status", e.statusPage)
	eapi.POST("/start", e.startInstance)

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

func (e *EbrowserApp) startInstance(ctx *gin.Context) {

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

	time.Sleep(time.Second * 20)

	ctx.JSON(200, gin.H{
		"pid":  cmd.Process.Pid,
		"port": e.port,
	})

	go func() {
		e.cmd.Wait()

		e.cLock.Lock()
		defer e.cLock.Unlock()
		e.cmd = nil

	}()

}

func (e *EbrowserApp) statusPage(ctx *gin.Context) {

	localSocket := path.Join(e.configurer.BasePath, e.config.LocalSocket)

	pp.Println("status/1", localSocket)

	conn, err := net.Dial("unix", localSocket)
	pp.Println("status/1.5")

	if err != nil {
		ctx.JSON(200, gin.H{
			"is_running": false,
			"status":     err.Error(),
		})
		return
	}

	pp.Println("status/2")

	defer conn.Close()

	out, err := io.ReadAll(conn)
	if err != nil {
		ctx.JSON(200, gin.H{
			"is_running": false,
			"status":     err.Error(),
		})
		return
	}

	pp.Println("status/3")

	msg := server.LocalStatus{}

	err = json.Unmarshal(out, &msg)
	if err != nil {
		ctx.JSON(200, gin.H{
			"is_running": false,
			"status":     err.Error(),
		})
		return
	}

	pp.Println("status/4")

	ctx.JSON(200, gin.H{
		"is_running":  true,
		"port":        msg.Port,
		"status":      "ok",
		"working_dir": e.configurer.BasePath,
	})

}