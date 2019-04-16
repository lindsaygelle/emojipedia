package web

import (
	"strings"

	"github.com/gellel/emojipedia/manifest"

	emojipedia "github.com/gellel/emojipedia/emojipedia-web/web-emojipedia"
	unicode "github.com/gellel/emojipedia/emojipedia-web/web-unicode"
)

var Key = "WEB"

var programs = map[string](func(m *manifest.Manifest, previous, options []string)){
	emojipedia.Key: emojipedia.Main,
	unicode.Key:    unicode.Main}

func Main(m *manifest.Manifest, previous, options []string) {
	var argument string
	if len(options) != 0 {
		argument = strings.ToUpper(options[0])
	}
	if f, ok := programs[argument]; ok {
		f(m, append(previous, argument), options[1:])
	}
}
