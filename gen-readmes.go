package main

import (
	"fmt"
	"io/ioutil"
	"log"
	"path/filepath"
	"strings"

	"github.com/jpillora/md-tmpl/mdtmpl"
	"github.com/jpillora/opts"
)

func main() {
	type config struct {
		Directory string `help:"examples directory"`
		Filter    string `help:"string to filter which examples are generated"`
		Preview   bool   `help:"display all shell commands in all files"`
	}
	c := config{
		Directory: ".",
		Filter:    "",
	}
	opts.Parse(&c)
	egs, err := ioutil.ReadDir(c.Directory)
	check(err)
	for _, s := range egs {
		eg := s.Name()
		if !s.IsDir() || !strings.HasPrefix(eg, "eg-") {
			continue
		}
		if !strings.Contains(eg, c.Filter) {
			continue
		}
		check(err)
		processReadme(eg, c.Preview)
		processGo(eg)
	}
}

func processGo(eg string) {
	b, err := ioutil.ReadFile(filepath.Join(eg, "main.go"))
	if err != nil {
		log.Printf("example '%s' has no main.go file", eg)
		return
	}
	if len(b) == 0 {
		log.Fatalf("example '%s' has empty main.go file", eg)
	}
}

func processReadme(eg string, preview bool) {
	fp := filepath.Join(eg, "README.md")
	b, err := ioutil.ReadFile(fp)
	if err != nil {
		log.Printf("example '%s' has no README.md file", eg)
		return
	}
	if preview {
		for i, c := range mdtmpl.Commands(b) {
			fmt.Printf("%18s#%d %s\n", eg, i+1, c)
		}
		return
	}
	b = mdtmpl.ExecuteIn(b, eg)
	check(ioutil.WriteFile(fp, b, 0655))
	log.Printf("executed templates and rewrote '%s'", eg)
}

func check(err error) {
	if err != nil {
		log.Fatal(err)
	}
}
