## helloworld example

<!--tmpl,code=go:cat main.go -->
``` go 
package main

import (
	"fmt"

	"github.com/jpillora/opts"
)

func main() {
	type config struct {
		File  string `opts:"help=file to load"`
		Lines int    `opts:"help=number of lines to show"`
	}
	c := config{}
	opts.Parse(&c)
	fmt.Printf("%+v\n", c)
}
```
<!--/tmpl-->

```
$ eg-helloworld --file zip.txt --lines 42
```

<!--tmpl,code=plain:go run main.go --file zip.txt --lines 42 -->
``` plain 
{File:zip.txt Lines:42}
```
<!--/tmpl-->

```
$ eg-helloworld --help
```

<!--tmpl,code=plain:go build -o eg-helloworld && ./eg-helloworld --help ; rm eg-helloworld -->
``` plain 

  Usage: eg-helloworld [options]

  Options:
  --file, -f   file to load
  --lines, -l  number of lines to show
  --help, -h   display help

```
<!--/tmpl-->
