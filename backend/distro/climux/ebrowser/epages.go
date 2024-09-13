package ebrowser

import (
	_ "embed"
	"html/template"
	"strings"
)

//go:embed start.tpl.html
var StartTemplate string

type TurnixeOptions struct {
	LocalExists  bool
	LocalRunning bool
	LocalFile    string
}

func RenderPage(opts TurnixeOptions) (string, error) {

	tpl, err := template.New("start_page").Parse(StartTemplate)
	if err != nil {
		return "", err
	}

	var b strings.Builder
	err = tpl.Execute(&b, &opts)
	if err != nil {
		return "", err
	}

	return b.String(), nil
}
