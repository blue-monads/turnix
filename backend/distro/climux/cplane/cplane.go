package cplane

import (
	_ "embed"
	"fmt"
	"net/url"
	"os/exec"
	"strings"
	"sync"

	"github.com/blue-monads/turnix/backend/distro/browser"
	"github.com/blue-monads/turnix/backend/distro/climux"
	"github.com/blue-monads/turnix/backend/mesh/lpweb"
	xutils "github.com/blue-monads/turnix/backend/utils"
	_ "github.com/chromedp/chromedp"
	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"
)

type Cplane struct {
	clictx     climux.Context
	port       int
	apiToken   string
	config     *climux.ConfigModel
	configurer *climux.Configued

	cmd   *exec.Cmd
	cLock sync.Mutex

	mesh          *lpweb.LPWebMesh
	meshBuildLock sync.Mutex

	engine  *gin.Engine
	browser browser.Browser
}

func New(clictx climux.Context) *Cplane {

	c := climux.NewConfigued("")
	config := c.GetConfig()

	apiToken, err := xutils.GenerateRandomString(32)
	if err != nil {
		panic(err)
	}

	return &Cplane{

		clictx:     clictx,
		port:       0,
		mesh:       nil,
		apiToken:   apiToken,
		config:     config,
		configurer: c,
		cmd:        nil,
		cLock:      sync.Mutex{},
		browser:    browser.InstanceBrowser,
	}
}

func (e *Cplane) Run() {

	rurl := fmt.Sprintf("http://localhost:%d/z/pages/startpage", e.port)

	if e.browser != nil {
		err := e.browser.Navigate(rurl)
		if err != nil {
			pp.Println("gfhandle load error", err)
			panic(err)
		}
	}

	e.startMesh()

}

func (e *Cplane) remoteNavigate(urlstr string, routeToPages bool) {

	pp.Println("@remoteNavigate/1", urlstr)

	u, err := url.Parse(urlstr)
	if err != nil {
		panic(err)
	}

	hostName := u.Hostname()

	if routeToPages {
		u.Path = "/z/pages"
	}

	pp.Println("@remoteNavigate/2", hostName)

	if strings.HasSuffix(hostName, ".lpweb") {
		keyHash := strings.TrimSuffix(hostName, ".lpweb")
		u.Scheme = "http"
		u.Host = fmt.Sprintf("%s-lpweb.localhost:%d", keyHash, e.port)

		pp.Println("@remoteNavigate/3", u.String())

		if e.browser != nil {
			e.browser.Navigate(u.String())
		}

	} else {
		pp.Println("@remoteNavigate/4", urlstr)
		if e.browser != nil {
			e.browser.Navigate(u.String())
		}

	}

}

func (e *Cplane) Close() {

	e.cLock.Lock()
	defer e.cLock.Unlock()

	if e.cmd != nil {
		e.cmd.Process.Kill()
		e.cmd = nil
	}

}
