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
	options := []interface{}{files.Files, get.Get}
	manifest := cli.NewManifest(runtime.Caller(0))
	program := cli.NewProgramFromManifest(manifest, options)
	switch len(os.Args) {
	case 0:
		panic(fmt.Errorf("%s", strings.Join(os.Args, ",")))
	case 1:
		fmt.Println(program.Use)
	default:
		for _, function := range program.Functions {
			if strings.ToLower(os.Args[1]) == strings.ToLower(function.Name) {
				os.Exit(function.F.(func(...string) int)(os.Args[1:]...))
			}
		}
		os.Exit(2)
	}
}
