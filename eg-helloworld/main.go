package main

import (
	"log"

	"github.com/jpillora/opts"
)

func main() {
	type config struct {
		File  string `opts:"help=file to load"`
		Lines int    `opts:"help=number of lines to show"`
	}
	c := config{}
	opts.Parse(&c)
	log.Printf("%+v", c)
}
