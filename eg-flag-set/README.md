## simple example

<!--tmpl,code=go:cat main.go -->
``` go 
package main

import (
	"github.com/golang/glog"
	"github.com/jpillora/opts"
)

func main() {
	opts.New(&app{}).
		EmbedGlobalFlagSet().
		Complete().
		SetLineWidth(90).
		Parse().
		RunFatal()
}

type app struct {
}

func (a *app) Run() {
	glog.Infof("hello from app via glog\n")
}
```
<!--/tmpl-->

```
$ eg-flag-set -a
```

<!--tmpl,code=plain:go run main.go -a -->
``` plain 
I0519 22:53:44.607289   69530 main.go:21] hello from app via glog
```
<!--/tmpl-->

```
$ eg-flag-set --help
```

<!--tmpl,code=plain:go build -o eg-flag-set && ./eg-flag-set --help ; rm eg-flag-set -->
``` plain 

  Usage: eg-flag-set [options]

  Options:
  --alsologtostderr, -a   log to standard error as well as files (default false)
  --log_backtrace_at, -l  when logging hits line file:N, emit a stack trace (default :0)
  --log_dir               If non-empty, write log files in this directory
  --logtostderr           log to standard error instead of files (default false)
  --stderrthreshold, -s   logs at or above this threshold go to stderr (default 0)
  --v                     log level for V logs (default 0)
  --vmodule               comma-separated list of pattern=N settings for file-filtered logging
  --help, -h              display help

  Completion options:
  --install, -i           install fish-completion
  --uninstall, -u         uninstall fish-completion

```
<!--/tmpl-->
