package main

import (
	"log"

	"github.com/jpillora/opts-v1"
)

type Config struct{}

func main() {
	opts.New(&Config{}).
		AddCommand(
			opts.New(&Foo{}).
				AddCommand(
					opts.New(&Bar{}),
				),
		).
		Parse().
		Run()
}

type Foo struct {
	Ping string
	Pong string
}

func (f *Foo) Run() {
	log.Printf("foo: %+v", f)
}

type Bar struct {
	Zip string
	Zop string
}

func (b *Bar) Run() {
	log.Printf("bar: %+v", b)
}
