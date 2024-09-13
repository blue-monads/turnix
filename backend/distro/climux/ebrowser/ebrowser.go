package ebrowser

import (
	_ "embed"
	"encoding/json"
	"fmt"
	"time"

	"github.com/bornjre/turnix/backend/distro/climux"

	webview "github.com/webview/webview_go"

	"github.com/k0kubun/pp"
)

// p2p-eproxy

type EbrowserApp struct {
	webview webview.WebView
	clictx  climux.Context
}

func New(clictx climux.Context) *EbrowserApp {

	w := webview.New(true)
	w.SetSize(900, 700, webview.HintNone)

	return &EbrowserApp{
		webview: w,
		clictx:  clictx,
	}
}

func (e *EbrowserApp) RunWithStartPage(opts TurnixeOptions) {

	e.webview.Bind("__ebrowser_rpc__", e.__BindEbrowserRPC)

	shtml, err := RenderPage(opts)
	if err != nil {
		panic(err)
	}

	go func() {
		time.Sleep(time.Second * 2)
		pp.Println(e.__sendRPC("turnix_start", map[string]any{}))
	}()

	e.webview.SetTitle("Turnix Start")
	e.webview.SetHtml(shtml)
	e.webview.Run()

}

func (e *EbrowserApp) __BindEbrowserRPC(name string, opts map[string]string) {
	pp.Println("@ctx", name, opts)

	go func() {

		switch name {
		case "connect_local":

			if opts["init_instance"] == "true" {
				err := e.clictx.RunCLI("app", []string{"app", "init"})
				if err != nil {
					pp.Println("@cannot_init", err)
					return
				}
			}

			if opts["start_instance"] == "true" || opts["init_instance"] == "true" {
				go func() {
					err := e.clictx.RunCLI("app", []string{"app", "start"})
					if err != nil {
						pp.Println("@cannot_start", err)
						return
					}
				}()
			}

			time.Sleep(time.Second * 5)

			e.NavigateLocal("FIXME FILE")

		case "connect_remote":

		}

	}()

}

func (e *EbrowserApp) NavigateLocal(file string) error {

	url := fmt.Sprintf("http://localhost%s/z/pages", ":7777")

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
