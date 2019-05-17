## groups example

<!--tmpl,code=go:cat main.go -->
``` go 
package main

import (
	"log"

	"github.com/jpillora/opts"
)

type Config struct {
	Fizz string
	Buzz bool
	//Foo has an implicit `opts:"mode=embedded,group=Foo"`.
	//Could be be merged with config by unsetting group `opts:"group="`.
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
	log.Printf("%+v", c)
}
```
<!--/tmpl-->

Group order in the help text is first-use order

```
$ eg-groups --help
```

<!--tmpl,code=plain:go build -o eg-groups && ./eg-groups --help ; rm eg-groups -->
``` plain 

  Usage: eg-groups [options]

  Options:
  --fizz, -f
  --buzz, -b
  --help, -h  display help

  Foo options:
  --bar
  --bazz

  More options:
  --ping, -p
  --pong

```
<!--/tmpl-->
