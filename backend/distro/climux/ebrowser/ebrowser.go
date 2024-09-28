package ebrowser

import (
	_ "embed"
	"encoding/json"
	"fmt"

	"github.com/bornjre/turnix/backend/distro"
	"github.com/bornjre/turnix/backend/distro/climux"

	webview "github.com/webview/webview_go"

	"github.com/k0kubun/pp"
)

// p2p-eproxy

type EbrowserApp struct {
	webview    webview.WebView
	clictx     climux.Context
	port       int
	apiToken   string
	config     *climux.ConfigModel
	configurer *climux.Configued
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
}

func (e *EbrowserApp) NavigateLocal(file string) error {

	url := fmt.Sprintf("http://localhost%s/z/pages", distro.DefaultOption)

	pp.Println("@opening_url", url)
	e.webview.Navigate(url)
	pp.Println("@should_open_newLurl")

	return nil

}

func (e *EbrowserApp) __sendRPC(name string, data any) error {

	out, err := json.Marshal(data)
	if err != nil {

		return err
	}

	e.webview.Eval(fmt.Sprintf(`__handle_rpc__("%s", "%s" )`, name, out))
	return nil
}

func (e *EbrowserApp) Close() {
	e.webview.Destroy()
}
