package main

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"

	"github.com/gellel/emojipedia/x"

	"github.com/gellel/emojipedia/manifest"

	files "github.com/gellel/emojipedia/emojipedia-files"
	get "github.com/gellel/emojipedia/emojipedia-get"
	web "github.com/gellel/emojipedia/emojipedia-web"
)

const filename = "manifest.json"

var (
	_, file, _, _ = runtime.Caller(0)
	dir           = filepath.Dir(file)
	m             = manifest.NewManifest(filepath.Join(dir, filename))
	routines      = x.NewRoutines(files.Export, get.Export, web.Export)
	runner        = (&x.Runner{}).Use(routines)
)

func main() {

	ok := runner.Next(os.Args[1]).Call(m, []string{os.Args[1]}, os.Args[2:])
	if ok != true {
		fmt.Println("help msg.")
	}
}
