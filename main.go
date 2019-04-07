package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/gellel/emojipedia/cli"
	get "github.com/gellel/emojipedia/emojipedia-get"
)

var programs = []interface{}{get.Get}

func main() {
	manifest := cli.NewManifest(runtime.Caller(0))
	program := cli.NewProgramFromManifest(manifest, programs)
	switch len(os.Args) {
	case 0:
		panic(fmt.Errorf("%s", strings.Join(os.Args, ",")))
	case 1:
		fmt.Println(program.Use)
	default:
		switch strings.ToUpper(os.Args[1]) {
		case "GET":
			get.Main(os.Args[2:])
		}
	}
}
