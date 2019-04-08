package main

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/gellel/emojipedia/cli"
	files "github.com/gellel/emojipedia/emojipedia-files"
	get "github.com/gellel/emojipedia/emojipedia-get"
)

func main() {
	functions := []interface{}{files.Files, get.Get}
	manifest := cli.NewManifest(runtime.Caller(0))
	program := cli.NewProgramFromManifest(manifest, functions)
	switch len(os.Args) {
	case 0:
		panic(fmt.Errorf("%s", strings.Join(os.Args, ",")))
	case 1:
		fmt.Println(program.Use)
	default:
		switch strings.ToUpper(os.Args[1]) {
		case "FILES":
			files.Main(os.Args[1:])
		case "GET":
			get.Main(os.Args[1:])
		}
	}
}
