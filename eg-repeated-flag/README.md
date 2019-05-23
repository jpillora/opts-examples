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
		Files []string `opts:"help=a set of files to show"`
	}
	c := config{}
	opts.Parse(&c)
	fmt.Printf("%+v\n", c)
}
```
<!--/tmpl-->

```
$ eg-repeated-flag --file zip.txt --file zop.txt --file zap.txt
```

<!--tmpl,code=plain:go run main.go --file zip.txt --file zop.txt --file zap.txt -->
``` plain 
{Files:[zip.txt zop.txt zap.txt]}
```
<!--/tmpl-->

```
$ eg-repeated-flag --help
```

<!--tmpl,code=plain:go build -o eg-repeated-flag && ./eg-repeated-flag --help ; rm eg-repeated-flag -->
``` plain 

  Usage: eg-repeated-flag [options]

  Options:
  --file, -f  a set of files to show (allows multiple)
  --help, -h  display help

```
<!--/tmpl-->
