package main

import (
	"flag"
	"fmt"

	"github.com/everfore/exc"
	// "github.com/toukii/goutils"
)

var (
	gb  = false
	gt  = false
	gtb = false
	gT  = false
	gr  = ""
	gi  = false
	gs  = false
	gm  = ""
)

func init() {
	flag.BoolVar(&gb, "b", false, "-b [true] : go build")
	flag.BoolVar(&gi, "i", false, "-i [true] : go install")
	flag.BoolVar(&gt, "t", false, "-t [true] : go test -v -test.run")
	flag.BoolVar(&gT, "T", false, "-T [true] : go test -v")
	flag.BoolVar(&gtb, "tb", false, "-tb [true] : go test -v -bench")
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
	if gtb {
		var benchFunc string
		fmt.Print("bench func:")
		fmt.Scanf("%s", &benchFunc)
		testCMD := "go test"
		if benchFunc != "" {
			testCMD += fmt.Sprintf(" -bench %s", benchFunc)
		}
		cmd.Reset(testCMD).Execute()
		return
	}
	if gt {
		var testFunc string
		fmt.Print("test func:")
		fmt.Scanf("%s", &testFunc)
		testCMD := "go test -v"
		if testFunc != "" {
			testCMD += fmt.Sprintf(" -test.run %s", testFunc)
		}
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
