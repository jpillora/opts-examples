## env example

<!--tmpl,code=go:cat main.go -->
``` go 
package main

import (
	"fmt"

	"github.com/jpillora/opts"
)

type Config struct {
	Foo string `opts:"env=FOO"`
	Bar string `opts:"env"`
}

func main() {
	c := Config{}
	//NOTE: we could also use UseEnv(), which
	//adds 'env' to all fields.
	opts.New(&c).
		// UseEnv().
		Parse()
	fmt.Printf("%+v\n", c)
}
```
<!--/tmpl-->

```
$ export FOO=hello
$ export BAR=world
$ go run env.go
```

<!--tmpl,code=plain:(export FOO=hello && export BAR=world && go run main.go) -->
``` plain 
{Foo:hello Bar:world}
```
<!--/tmpl-->

```
$ eg-env --help
```

<!--tmpl,code=plain:go build -o eg-env && ./eg-env --help ; rm eg-env -->
``` plain 

  Usage: eg-env [options]

  Options:
  --foo, -f   env FOO
  --bar, -b   env BAR
  --help, -h  display help

```
<!--/tmpl-->
