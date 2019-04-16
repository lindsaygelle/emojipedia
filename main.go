package main

import (
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/gellel/emojipedia/manifest"

	files "github.com/gellel/emojipedia/emojipedia-files"
	get "github.com/gellel/emojipedia/emojipedia-get"
	web "github.com/gellel/emojipedia/emojipedia-web"
)

const filename = "manifest.json"

var _, file, _, _ = runtime.Caller(0)

var dir = filepath.Dir(file)

var m = manifest.NewManifest(filepath.Join(dir, filename))

var programs = map[string](func(m *manifest.Manifest, previous, options []string)){
	get.Key:   get.Main,
	files.Key: files.Main,
	web.Key:   web.Main}

func main() {
	var argument string
	if len(os.Args[1:]) != 0 {
		argument = strings.ToUpper(os.Args[1])
	}
	if f, ok := programs[argument]; ok {
		f(m, ([]string{m.Name, argument}), os.Args[2:])
	}
}
