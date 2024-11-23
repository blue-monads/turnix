package ebrowser

import (
	_ "embed"
	"fmt"
	"os/exec"
	"sync"

	"github.com/blue-monads/turnix/backend/distro/climux"
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

	e.webview.Run()

}

func (e *EbrowserApp) __BindEbrowserRPC(name string, opts map[string]string) {
	pp.Println("@ctx", name, opts)

	if name == "local-navigate" {
		e.webview.Navigate(opts["url"])
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
