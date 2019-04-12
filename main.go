package main

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"

	f "github.com/gellel/emojipedia/emojipedia-files"
	g "github.com/gellel/emojipedia/emojipedia-get"
	w "github.com/gellel/emojipedia/emojipedia-web"

	"github.com/gellel/emojipedia/manifest"
)

var set = map[string](func(m *manifest.Manifest, previous, options []string)){
	g.Key: g.Main,
	f.Key: f.Main,
	w.Key: w.Main}

func main() {
	filename := "manifest.json"
	_, file, _, _ := runtime.Caller(0)
	dir := filepath.Dir(file)
	m := manifest.NewManifest(filepath.Join(dir, filename))
	options := os.Args[1:]
	if len(options) != 0 {
		if f, ok := set[strings.ToUpper(options[0])]; ok != false {
			f(m, []string{m.Name, options[0]}, options[1:])
		}
	}
}
