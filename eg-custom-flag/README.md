## customtypes example

<!--tmpl,code=go:cat main.go -->
``` go 
package main

import (
	"fmt"
	"net"
	"os"
	"strconv"
	"time"

	"github.com/jpillora/opts"
)

type Config struct {
	Mmm   MagicInt
	Bar   time.Duration
	Zee   bool
	IP    net.IP
	Files []File
	Dir   Dir
}

func main() {
	c := Config{
		Mmm: 1,
		IP:  net.ParseIP("1.1.1.1"),
	}
	opts.Parse(&c)
	fmt.Printf("%+v\n", c)
}

//MagicInt is a valid custom type since it implements the flag.Value interface
type MagicInt int

func (b MagicInt) String() string {
	return "{" + strconv.Itoa(int(b)) + "}"
}

func (b *MagicInt) Set(s string) error {
	n, err := strconv.Atoi(s)
	if err != nil {
		return err
	}
	*b = MagicInt(n + 42)
	return nil
}

//File is a string flag.Value which also performs
//an os.Stat on itself during Set, confirming it
//references a file
type File string

func (f File) String() string {
	return string(f)
}

func (f *File) Set(s string) error {
	if info, err := os.Stat(s); os.IsNotExist(err) {
		return fmt.Errorf("'%s' does not exist", s)
	} else if err != nil {
		return err
	} else if info.IsDir() {
		return fmt.Errorf("'%s' is a directory", s)
	}
	*f = File(s)
	return nil
}

//Dir is a string flag.Value which also performs
//an os.Stat on itself during Set, confirming it
//references a directory
type Dir string

func (d Dir) String() string {
	return string(d)
}

func (d *Dir) Set(s string) error {
	if info, err := os.Stat(s); os.IsNotExist(err) {
		return fmt.Errorf("'%s' does not exist", s)
	} else if err != nil {
		return err
	} else if !info.IsDir() {
		return fmt.Errorf("'%s' is a file", s)
	}
	*d = Dir(s)
	return nil
}
```
<!--/tmpl-->

```sh
#NOTE: 5 + 42 = 47
$ eg-custom-flag --foo 2m --bar 5 --bazz 5
```

<!--tmpl,code=plain:go run main.go --foo 2m --bar 5 --bazz 5 -->
``` plain 

  Usage: main [options]

  Options:
  --mmm, -m   default {1}
  --bar, -b
  --zee, -z
  --ip, -i    default 1.1.1.1
  --file, -f  allows multiple
  --dir, -d
  --help, -h  display help

  Error:
    flag provided but not defined: -foo

```
<!--/tmpl-->

```
$ eg-custom-flag --help
```

<!--tmpl,code=plain:go install && eg-custom-flag --help ; rm $(which eg-custom-flag) -->
``` plain 

  Usage: eg-custom-flag [options]

  Options:
  --mmm, -m   default {1}
  --bar, -b
  --zee, -z
  --ip, -i    default 1.1.1.1
  --file, -f  allows multiple
  --dir, -d
  --help, -h  display help

```
<!--/tmpl-->
