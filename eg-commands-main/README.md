## Commands example (main package)

Here, we're using `AddCommand` to add other `Opts` instances into our root instance:

<!--tmpl,code=go:cat main.go -->
``` go 
package main

import (
	"fmt"

	"github.com/jpillora/opts"
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
	fmt.Printf("foo: %+v\n", f)
}

type Bar struct {
	Zip string
	Zop string
}

func (b *Bar) Run() {
	fmt.Printf("bar: %+v\n", b)
}
```
<!--/tmpl-->

```
$ eg-commands-inline foo bar --zip 2
```

<!--tmpl,code=plain:go run main.go foo bar --zip 2 -->
``` plain 
bar: &{Zip:2 Zop:}
```
<!--/tmpl-->

```
$ eg-commands-inline --help
```

<!--tmpl,code=plain:go build -o eg-commands-inline && ./eg-commands-inline --help ; rm eg-commands-inline -->
``` plain 

  Usage: eg-commands-inline [options] <command>

  Options:
  --help, -h  display help

  Commands:
  · foo

```
<!--/tmpl-->
