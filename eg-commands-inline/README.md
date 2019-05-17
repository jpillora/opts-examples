## cmds example

<!--tmpl,code=go:cat main.go -->
``` go 
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
```
<!--/tmpl-->

```
$ cmds bar --zip hello --zap world
```

<!--tmpl,code=plain:go run main.go bar --zip hello --zap world -->
``` plain 
2019/05/18 02:05:51 bar: &{Zip:hello Zap:world}
```
<!--/tmpl-->

```
$ cmds --help
```

<!--tmpl,code=plain:go build -o eg-commands-inline && ./eg-commands-inline --help ; rm eg-commands-inline -->
``` plain 

  Usage: eg-commands-inline [options] <command>

  Options:
  --help, -h  display help

  Commands:
  · foo - This text also becomes commands summary text
  · bar - command two of two

```
<!--/tmpl-->
