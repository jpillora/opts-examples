package main

import (
	"log"

	"github.com/jpillora/opts"
)

type Config struct {
	//register commands by including them
	//in the parent struct
	Foo  `opts:"mode=cmd,help=This text also becomes commands summary text"`
	*Bar `opts:"mode=cmd,help=command two of two"`
}

func main() {
	c := Config{}
	opts.New(&c).
		Parse().
		Run()
}

type Foo struct {
	Ping string
	Pong string
}

func (f *Foo) Run() error {
	log.Printf("foo: %+v", f)
	return nil
}

type Bar struct {
	Zip string
	Zap string
}

func (b *Bar) Run() error {
	log.Printf("bar: %+v", b)
	return nil
}
