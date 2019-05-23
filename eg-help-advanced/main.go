package main

import (
	"fmt"

	"github.com/jpillora/opts"
)

type Config struct {
	Foo string
	Bar string
}

func main() {
	c := Config{}
	//see default templates and the default template order
	//in the opts/help.go file
	o := opts.New(&c).
		Summary("some text under usage").
		DocAfter("summary", "mytext", "\nand a raw template under the summary!\n"). //add new entry
		Repo("myfoo.com/bar").
		DocSet("repo", "\nMy awesome repo:\n  {{.Repo}}"). //change existing entry
		Parse()

	fmt.Println(o.Help())
}
