package main

import (
	"flag"
	"fmt"

	"github.com/everfore/exc"
	// "github.com/toukii/goutils"
)

var (
	gb = false
	gT = false
	gr = ""
	gi = false
	gs = false

	gt  = ""
	gtb = ""
	gm  = ""
)

func init() {
	flag.BoolVar(&gb, "b", false, "-b [true] : go build")
	flag.BoolVar(&gi, "i", false, "-i [true] : go install")
	flag.BoolVar(&gT, "T", false, "-T [true] : go test -v")
	flag.StringVar(&gt, "t", "", "go test -v -test.run XX")
	flag.StringVar(&gtb, "tb", ".", "go test -bench XX")
	flag.StringVar(&gr, "r", "main.go", "-r main.go : go run file.go")
	flag.StringVar(&gm, "m", "", "-m commit : git add -A;git commit -m --")

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
	if gi {
		cmd.Reset("go install").Execute()
		return
	}
	if gT {
		testCMD := "go test -v"
		cmd.Reset(testCMD).Execute()
		return
	}
	if gt != "" {
		testCMD := fmt.Sprintf("go test -v -test.run %s", gt)
		cmd.Reset(testCMD).Execute()
		return
	}
	if gtb != "." {
		testCMD := fmt.Sprintf("go test -bench=%s", gtb)
		cmd.Reset(testCMD).Execute()
		return
	}
	if len(gm) > 0 {
		_, err := cmd.Reset("git add -A").DoNoTime()
		if nil != err {
			return
		}
		cmd.Reset(fmt.Sprintf("git commit -m %s", gm)).Execute()
		return
	}
	cmd.Reset(fmt.Sprintf("go run %s", gr)).Execute()
}
