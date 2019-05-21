package main

import (
	"github.com/jpillora/opts"
	"github.com/jpillora/opts-examples/eg-commands-register/foo"
)

type cmd struct{}

func main() {
	c := cmd{}
	//default name for the root command (package main) is the binary name
	opts.New(&c).
		AddCommand(foo.New()).
		Parse().
		RunFatal()
}
