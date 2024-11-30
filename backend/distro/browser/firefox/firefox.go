package firefox

import (
	"github.com/k0kubun/pp"
	gofirefox "github.com/unikiosk/go-firefox"
)

type Firefox struct {
	gfhandle gofirefox.UI
}

func New() *Firefox {
	return &Firefox{
		gfhandle: nil,
	}
}

func (f *Firefox) Start() error {

	gfhandle, err := gofirefox.New(
		"data:text/html,<html>Turnix Control Plane</html>",
		[]string{"--new-instance", "--profile", "./tmp"},
		nil)
	if err != nil {
		pp.Println("gfhandle init error", err)
		panic(err)
	}

	f.gfhandle = gfhandle

	return nil
}

func (f *Firefox) Navigate(url string) error {

	f.gfhandle.Load(url)

	return nil
}
