package web

import (
	"strings"

	e "github.com/gellel/emojipedia/emojipedia-web/web-emojipedia"
	u "github.com/gellel/emojipedia/emojipedia-web/web-unicode"
	"github.com/gellel/emojipedia/manifest"
)

var Exports = func(options ...string) {}

var Key = "WEB"

var Options = []interface{}{
	e.Export,
	u.Export}

var set = map[string](func(m *manifest.Manifest, previous, options []string)){
	e.Key: e.Main,
	u.Key: u.Main}

func Main(m *manifest.Manifest, previous, options []string) {
	if len(options) != 0 {
		key := strings.ToUpper(strings.Replace(options[0], "-", "", -1))
		if f, ok := set[key]; ok {
			f(m, append(previous, key), options[1:])
		}
	}
}
