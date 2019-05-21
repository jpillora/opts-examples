## Commands example (individual packages)

Here, we're adding `Opts` instances from other packages into our root instance:

_`main.go`_

<!--tmpl,code=go:cat main.go -->
``` go 
package main

import (
	"github.com/jpillora/opts"
	"github.com/jpillora/opts-examples/eg-commands-register/foo"
)

type cmd struct{}

func main() {
	c := cmd{}
	//default name for the root command (package main) is the binary name
	opts.New(&c).
		AddCommand(foo.New()).
		Parse().
		RunFatal()
}
```
<!--/tmpl-->

_`foo/cmd.go`_

<!--tmpl,code=go:cat foo/cmd.go -->
``` go 
package foo

import (
	"fmt"

	"github.com/jpillora/opts"
	"github.com/jpillora/opts-examples/eg-commands-register/bar"
)

func New() opts.Opts {
	c := cmd{}
	//default name for a subcommand is its package name ("foo")
	return opts.New(&c).AddCommand(bar.New())
}

type cmd struct {
	Ping string
	Pong string
}

func (f *cmd) Run() error {
	fmt.Printf("foo: %+v\n", f)
	return nil
}
```
<!--/tmpl-->

```sh
$ ./eg-commands-register --help
```

<!--tmpl,code=plain:go build -o eg-commands-register && ./eg-commands-register --help ; rm eg-commands-register -->
``` plain 

  Usage: eg-commands-register [options] <command>

  Options:
  --help, -h  display help

  Commands:
  · foo

```
<!--/tmpl-->
