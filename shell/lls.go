package shell

import (
	"io/ioutil"
	"os"

	"github.com/abiosoft/ishell"
)

func llsCmd(ctx *ShellCtxt) *ishell.Cmd {
	return &ishell.Cmd{
		Name:      "lls",
		Help:      "list local directory",
		Completer: createDirCompleter(ctx),
		Func: func(c *ishell.Context) {

			c.Println()
			current_dir, err := os.Getwd()
			if err != nil {
				return
			}
			files, _ := ioutil.ReadDir(current_dir)

			for _, e := range files {
				Display(e, c)
			}

			c.Println()
			c.SetPrompt(ctx.prompt())
		},
	}
}
