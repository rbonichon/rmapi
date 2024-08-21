package shell

import (
	"os"
	"io/ioutil"
	
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
			if err != nil { return }
			files, _ := ioutil.ReadDir(current_dir)

			for _, e := range files {
				eType := "f"
				if e.IsDir() {
					eType = "d"
				}
				c.Printf("[%s]\t%s\n", eType, e.Name())
			}
			
			c.Println()
			c.SetPrompt(ctx.prompt())
		},
	}
}
