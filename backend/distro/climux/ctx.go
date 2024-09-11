package climux

import (
	"os"
	"os/exec"

	"github.com/temphia/temphia/code/backend/libx/easyerr"
)

type Context struct {
	Args []string
	R    Registery
}

func (c *Context) RunSubProcessCLI(name string, args []string) error {

	_, ok := c.R[name]
	if !ok {
		return easyerr.NotFound("cli")
	}

	ebin, err := os.Executable()
	if err != nil {
		return err
	}

	cmd := exec.Command(ebin, name)
	cmd.Args = append(cmd.Args, args...)

	return cmd.Run()

}

func (c *Context) RunCLI(name string, args []string) error {

	cli, ok := c.R[name]
	if !ok {
		return easyerr.NotFound("cli")
	}

	return cli.Func(Context{
		R:    c.R,
		Args: args,
	})

}
