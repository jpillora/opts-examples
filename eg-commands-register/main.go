package main

import (
	"github.com/jpillora/opts"
	"github.com/jpillora/opts-examples/eg-commands-register/foo"
)

type cmd struct{}

func main() {
	c := cmd{}
	//default name for the root command (package main) is the binary name
	o := opts.New(&c)
	foo.Register(o)
	o.Parse().RunFatal()
}
