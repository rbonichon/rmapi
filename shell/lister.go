package shell

import "github.com/abiosoft/ishell"

type Lister interface {
	IsDir() bool
	Name() string
}

func Display(e Lister, c *ishell.Context) {
	eType := "f"
	if e.IsDir() {
		eType = "d"
	}
	c.Printf("[%s]\t%s\n", eType, e.Name())
}
