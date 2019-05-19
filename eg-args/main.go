package main

import (
	"fmt"

	"github.com/jpillora/opts"
)

type Config struct {
	Shark  string   `opts:"mode=arg"`
	Octopi []string `opts:"mode=arg,min=1"`
}

func main() {
	c := Config{}
	opts.New(&c).Parse()
	fmt.Printf("%+v\n", c)
}
