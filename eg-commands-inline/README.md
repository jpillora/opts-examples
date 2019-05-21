## cmds example

<!--tmpl,code=go:cat main.go -->
``` go 
package main

import (
	"fmt"

	"github.com/jpillora/opts"
)

type Config struct {
	//register commands by including them in the parent struct
	Foo  `opts:"mode=cmd,help=This text also becomes commands summary text"`
	*Bar `opts:"mode=cmd,help=command two of two"`
}

func main() {
	c := Config{}
	opts.Parse(&c).Run()
}

type Foo struct {
	Ping string
	Pong string
}

func (f *Foo) Run() error {
	fmt.Printf("foo: %+v\n", f)
	return nil
}

type Bar struct {
	Zip string
	Zap string
}

func (b *Bar) Run() error {
	fmt.Printf("bar: %+v\n", b)
	return nil
}
```
<!--/tmpl-->

```
$ cmds bar --zip hello --zap world
```

<!--tmpl,code=plain:go run main.go bar --zip hello --zap world -->
``` plain 
bar: &{Zip:hello Zap:world}
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
  · bar  command two of two
  · foo  This text also becomes commands summary text

```
<!--/tmpl-->
