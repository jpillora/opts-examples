package main

import (
	"fmt"

	"github.com/jpillora/opts"
)

type Config struct {
	Foo string `opts:"env=FOO"`
	Bar string `opts:"env"`
}

func main() {
	c := Config{}
	//NOTE: we could also use UseEnv(), which
	//adds 'env' to all fields.
	opts.New(&c).
		// UseEnv().
		Parse()
	fmt.Printf("%+v\n", c)
}
