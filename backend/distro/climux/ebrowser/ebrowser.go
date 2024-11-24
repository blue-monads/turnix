package ebrowser

import (
	_ "embed"
	"fmt"
	"net/url"
	"os/exec"
	"strings"
	"sync"

	"github.com/blue-monads/turnix/backend/distro/climux"
	"github.com/blue-monads/turnix/backend/mesh/lpweb"
	"github.com/gin-gonic/gin"
	"github.com/k0kubun/pp"

	webview "github.com/webview/webview_go"
)

// p2p-eproxy

type EbrowserApp struct {
	webview    webview.WebView
	clictx     climux.Context
	port       int
	apiToken   string
	config     *climux.ConfigModel
	configurer *climux.Configued

	cmd   *exec.Cmd
	cLock sync.Mutex

	mesh          *lpweb.LPWebMesh
	meshBuildLock sync.Mutex

	engine *gin.Engine
}

func New(clictx climux.Context) *EbrowserApp {

	w := webview.New(true)
	w.SetSize(900, 700, webview.HintNone)

	c := climux.NewConfigued("")

	return &EbrowserApp{
		webview:    w,
		clictx:     clictx,
		port:       0,
		apiToken:   "",
		config:     c.GetConfig(),
		configurer: c,
		cmd:        nil,
		cLock:      sync.Mutex{},
	}
}

func (e *EbrowserApp) Run() {

	e.webview.Bind("__ebrowser_rpc__", e.__BindEbrowserRPC)

	e.webview.SetTitle("Turnix Start")
	e.webview.Navigate(fmt.Sprintf("http://localhost:%d/z/pages/startpage", e.port))

	go e.startMesh()

	e.webview.Run()

}

func (e *EbrowserApp) __BindEbrowserRPC(name string, opts map[string]string) {
	pp.Println("@ctx", name, opts)

	switch name {
	case "local-navigate":
		e.webview.Navigate(opts["url"])
	case "remote-navigate":
		e.remoteNavigate(opts["url"])
	default:
		panic(fmt.Errorf("unknown rpc name %s", name))
	}

}

func (e *EbrowserApp) remoteNavigate(urlstr string) {

	pp.Println("@remoteNavigate/1", urlstr)

	u, err := url.Parse(urlstr)
	if err != nil {
		panic(err)
	}

	hostName := u.Hostname()

	pp.Println("@remoteNavigate/2", hostName)

	if strings.HasSuffix(hostName, ".lpweb") {
		keyHash := strings.TrimSuffix(hostName, ".lpweb")
		u.Scheme = "http"
		u.Host = fmt.Sprintf("%s-lpweb.localhost:%d", keyHash, e.port)

		pp.Println("@remoteNavigate/3", u.String())

		e.webview.Navigate(u.String())
	} else {
		pp.Println("@remoteNavigate/4", urlstr)
		e.webview.Navigate(urlstr)
	}

}

func (e *EbrowserApp) Close() {

	e.cLock.Lock()
	defer e.cLock.Unlock()

	if e.cmd != nil {
		e.cmd.Process.Kill()
		e.cmd = nil
	}

	e.webview.Destroy()
}
