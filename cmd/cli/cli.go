package cli

import (
	"fmt"
	"os"

	"github.com/blue-monads/turnix/backend/distro/climux"
)

func RunCLI() {

	if len(os.Args) == 1 {
		os.Args = []string{os.Args[0], climux.DefaultCLI}
	}

	if os.Args[1] == "help" || os.Args[1] == "--help" {
		PrintHelpText()
		return
	}

	if os.Args[1] == "version" || os.Args[1] == "--version" {
		fmt.Printf("turnix %s", "0.0.1-dev")
		return
	}

	clis := climux.GetRegistry()
	acli, ok := clis[os.Args[1]]
	if !ok {
		fmt.Println("not found cli :", os.Args)
		os.Exit(1)
		return
	}

	err := acli.Func(climux.Context{
		Args: os.Args[1:],
		R:    clis,
	})
	if err != nil {
		panic(err)
	}
}

func PrintHelpText() {
	clis := climux.GetRegistry()

	fmt.Printf(`Turnix is a platform for apps.
Usage:
	
	turnix <command> [arguments]

The commands are:
`)

	for _, v := range clis {
		fmt.Printf("\t%s  \t\t%s\n", v.Name, v.Help)
	}

	fmt.Printf(`	version		print turnix version
	help  		this help page

Use "turnix <command> help" for more information about a command.
`)
}
