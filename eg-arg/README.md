## arg example

<!--tmpl,code=go:cat main.go -->
``` go 
package main

import (
	"fmt"

	"github.com/jpillora/opts"
)

type Config struct {
	Foo string `opts:"mode=arg,help=<foo> is a very important argument"`
	Bar string
}

func main() {
	c := Config{}
	opts.New(&c).Parse()
	fmt.Printf("%+v\n", c)
}
```
<!--/tmpl-->

```
$ eg-arg --foo hello --bar world
```

<!--tmpl,code=plain:go run main.go --foo hello --bar world -->
``` plain 

  Usage: main [options] <foo>

  <foo> is a very important argument

  Options:
  --bar, -b
  --help, -h  display help

  Error:
    flag provided but not defined: -foo

```
<!--/tmpl-->

```
$ eg-arg --help
```

<!--tmpl,code=plain:go build -o eg-arg && ./eg-arg --help ; rm eg-arg -->
``` plain 

  Usage: eg-arg [options] <foo>

  <foo> is a very important argument

  Options:
  --bar, -b
  --help, -h  display help

```
<!--/tmpl-->
