package foo

import (
	"fmt"

	"github.com/jpillora/opts"
	"github.com/jpillora/opts-examples/eg-commands-register/bar"
)

func Register(parent opts.Opts) {
	c := cmd{}
	//default name for a subcommand is its package name ("foo")
	o := opts.New(&c).Call(bar.Register)
	parent.AddCommand(o)
}

type cmd struct {
	Ping string
	Pong string
}

func (f *cmd) Run() error {
	fmt.Printf("foo: %+v\n", f)
	return nil
}
