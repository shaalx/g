package main

import (
	"flag"
	"fmt"
	"github.com/everfore/exc"
	// "github.com/shaalx/goutils"
)

var (
	gb = false
	gr = ""

	gs = false
)

func init() {
	flag.BoolVar(&gb, "b", false, "-b [true]")
	flag.StringVar(&gr, "r", "main.go", "-r main.go")
	flag.BoolVar(&gs, "s", false, "-s [true]")
}

func main() {
	flag.Parse()
	cmd := exc.NewCMD("git status").Debug()
	if gs {
		cmd.Execute()
		return
	}
	if gb {
		cmd.Reset("go build").Execute()
		return
	} else {
		cmd.Reset(fmt.Sprintf("go run %s", gr)).Execute()
		return
	}
}
