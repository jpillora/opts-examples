package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/jpillora/opts"
)

type Config struct {
	Mmm   []MagicInt
	Bar   time.Duration
	Zee   bool
	Files []File
	Dir   Dir
}

func main() {
	c := Config{}
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
