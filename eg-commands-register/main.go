package main

import (
	"github.com/jpillora/opts-examples/eg-commands-register/foo"
	"github.com/jpillora/opts-v1"
)

type cmd struct{}

func main() {
	c := cmd{}
	o := opts.New(&c)
	foo.Register(o)
	o.Parse().RunFatal()
}
