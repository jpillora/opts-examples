package main

import (
	"fmt"

	"github.com/jpillora/opts"
)

func main() {
	type config struct {
		Files []string `opts:"help=a set of files to show"`
	}
	c := config{}
	opts.Parse(&c)
	fmt.Printf("%+v\n", c)
}
