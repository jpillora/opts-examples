package main

import (
	"fmt"

	"github.com/jpillora/opts"
)

type Config struct {
	Foo string `opts:"mode=arg,help=<foo> is a very important argument"`
	Bar string
}

func main() {
	c := Config{}
	opts.New(&c).Parse()
	fmt.Printf("%+v\n", c)
}
