## Completion example with custom completer 

<!--tmpl,code=go:cat main.go -->
``` go 
package main

import (
	"fmt"
	"strings"

	"github.com/jpillora/opts"
)

type Config struct {
	Alpha string
	Beta  myEnum
	Delta `opts:"mode=cmd"`
	Echo  `opts:"mode=cmd"`
}

type Delta struct {
	Zip bool
	Zop string
	Dir string
}

type Echo struct {
	Ping  bool
	Pong  string
	Files []string
}

func main() {
	c := Config{}
	opts.New(&c).
		Complete().
		Parse()
	fmt.Printf("%+v\n", c)
}

//myEnum implements completer and setter
//to provide an enum-like field
type myEnum string

var valid = []string{"a", "b", "c"}

func (myEnum) Complete(s string) []string {
	return valid
}

func (e *myEnum) Set(s string) error {
	for _, v := range valid {
		if s == v {
			*e = myEnum(s)
			return nil
		}
	}
	return fmt.Errorf("must be one of: %v", strings.Join(valid, ", "))
}
```
<!--/tmpl-->


```
$ eg-complete --help
```

<!--tmpl,code=plain:go build -o eg-complete && ./eg-complete --help -->
``` plain 

  Usage: eg-complete [options] <command>

  Options:
  --alpha, -a
  --beta, -b
  --help, -h       display help

  Completion options:
  --install, -i    install fish-completion
  --uninstall, -u  uninstall fish-completion

  Commands:
  · delta
  · echo

```
<!--/tmpl-->

The help above shows your prefered shell but installation actually inserts completers for all known shells (bash, zsh and fish).

### Install the completion scripts
`./eg-complete -i`
<!--tmpl,code=plain:go build -o eg-complete && ./eg-complete -i -->
``` plain 
Installed
```
<!--/tmpl-->

Attempt second install

`./eg-complete -i`
<!--tmpl,code=plain:./eg-complete -i -->
``` plain 
already installed in /Users/jpillora/.bashrc
already installed at /Users/jpillora/.config/fish/completions/eg-complete.fish
```
<!--/tmpl-->


*Tip: You'll need to re-source your shell to load the new completion*

### Completion of commands

When you press `TAB` your shell will set the `COMP_LINE` variable and run your binary, expecting a list of strings which represent the next valid arguments.

By default, you'll see available commands:

`COMP_LINE="./eg-complete " ./eg-complete`
<!--tmpl,code=plain:COMP_DEBUG= COMP_LINE="./eg-complete " ./eg-complete -->
``` plain 
delta
echo
```
<!--/tmpl-->

### Completion of flags

If you type `-` (dash) `TAB`, you'll see available flags:

`COMP_LINE="./eg-complete -" ./eg-complete `
<!--tmpl,code=plain:COMP_DEBUG= COMP_LINE="./eg-complete -" ./eg-complete -->
``` plain 
-b
-i
--uninstall
-u
--alpha
-a
--beta
--help
-h
--install
```
<!--/tmpl-->


### Custom completion

The `--beta` flag implements a custom completer:

`COMP_LINE="./eg-complete -b " ./eg-complete `
<!--tmpl,code=plain:COMP_DEBUG= COMP_LINE="./eg-complete -b " ./eg-complete -->
``` plain 
a
b
c
```
<!--/tmpl-->

### Completion of sub-commands

Sub-command completion works as expected:

<!--/tmpl-->
`COMP_LINE="./eg-complete delta -" ./eg-complete`
<!--tmpl,code=plain:COMP_DEBUG= COMP_LINE="./eg-complete delta -" ./eg-complete -->
``` plain 
--zip
-z
--zop
--dir
-d
--help
-h
```
<!--/tmpl-->

### Uninstall
`./eg-complete -u`
<!--tmpl,code=plain:./eg-complete -u && rm eg-complete -->
``` plain 
Uninstalled
```
<!--/tmpl-->


*Tip: Write your debug logs to a file:*

```go
//Since your binary is doing the completion, debugging your
//completions can be tricky. You can use this `debugf` function
//to write to a file instead of the terminal during development:
func debugf(f string, a ...interface{}) {
	l, err := os.OpenFile("/tmp/opts.log", os.O_CREATE|os.O_APPEND|os.O_WRONLY, 0600)
	if err == nil {
		fmt.Fprintf(l, f+"\n", a...)
		l.Close()
	}
}
//then you can `tail -f /tmp/opts/log`
```