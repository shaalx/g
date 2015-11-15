package main

import (
	"flag"
	"fmt"
	"github.com/everfore/exc"
	// "github.com/shaalx/goutils"
)

var (
	gb = false
	gt = false
	gr = ""

	gs = false
)

func init() {
	flag.BoolVar(&gb, "b", false, "-b [true] : go build")
	flag.BoolVar(&gt, "t", false, "-t [true] : go test -v")
	flag.StringVar(&gr, "r", "main.go", "-r main.go : go run file.go")

	flag.BoolVar(&gs, "s", false, "-s [true] : git status")
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
	}
	if gt {
		cmd.Reset("go test -v").Execute()
		return
	}
	cmd.Reset(fmt.Sprintf("go run %s", gr)).Execute()
	return
}
