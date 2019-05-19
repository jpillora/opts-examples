package bar

import (
	"fmt"

	"github.com/jpillora/opts"
)

func Register(parent opts.Opts) {
	c := cmd{}
	//default name for a subcommand is its package name ("bar")
	o := opts.New(&c)
	parent.AddCommand(o)
}

type cmd struct {
	Zip string
	Zop string
}

func (b *cmd) Run() error {
	fmt.Printf("bar: %+v\n", b)
	return nil
}
