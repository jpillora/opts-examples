package main

import (
	"fmt"

	"github.com/jpillora/opts"
)

func main() {
	type config struct {
		File  string `opts:"help=file to load"`
		Lines int    `opts:"help=number of lines to show"`
	}
	c := config{}
	opts.Parse(&c)
	fmt.Printf("%+v\n", c)
}
