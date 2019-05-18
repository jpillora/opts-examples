package foo

import (
	"log"

	"github.com/jpillora/opts-examples/eg-commands-register/bar"
	"github.com/jpillora/opts"
)

func Register(parent opts.Opts) {
	c := cmd{}
	//default name for a subcommand is its package name ("foo")
	o := opts.New(&c)
	bar.Register(o)
	parent.AddCommand(o)
}

type cmd struct {
	Ping string
	Pong string
}

func (f *cmd) Run() error {
	log.Printf("foo: %+v", f)
	return nil
}
