package shell

import (
	"os"
	
	"github.com/abiosoft/ishell"
)

func lcdCmd(ctx *ShellCtxt) *ishell.Cmd {
	return &ishell.Cmd{
		Name:      "lcd",
		Help:      "change directory on local system",
		Completer: createDirCompleter(ctx),
		Func: func(c *ishell.Context) {
			if len(c.Args) == 0 {
				return
			}

			target := c.Args[0]

			current_dir, err := os.Getwd()
			if err != nil { return }
			
			os.Chdir(target)
			
			current_dir, _ = os.Getwd()
			ctx.local_path = current_dir

			c.Println()
			c.SetPrompt(ctx.prompt())
		},
	}
}
