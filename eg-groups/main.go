package main

import (
	"fmt"

	"github.com/jpillora/opts"
)

type Config struct {
	Fizz string
	Buzz bool
	//Foo has an implicit "mode=embedded,group=Foo", it
	//could be be merged into the default group by
	//unsetting group "group=".
	Foo
	Ping, Pong int `opts:"group=More"`
}

type Foo struct {
	Bar  int
	Bazz int
}

func main() {
	c := Config{}
	opts.Parse(&c)
	fmt.Printf("%+v\n", c)
}
