package web

import (
	"strings"

	"github.com/gellel/emojipedia/manifest"

	emojipedia "github.com/gellel/emojipedia/emojipedia-web/web-emojipedia"
	unicode "github.com/gellel/emojipedia/emojipedia-web/web-unicode"
)

var Exports = func(options ...string) {}

var Key = "WEB"

var Options = []interface{}{
	emojipedia.Export,
	unicode.Export}

var set = map[string](func(m *manifest.Manifest, previous, options []string)){
	emojipedia.Key: emojipedia.Main,
	unicode.Key:    unicode.Main}

func Main(m *manifest.Manifest, previous, options []string) {
	if len(options) != 0 {
		key := strings.ToUpper(strings.Replace(options[0], "-", "", -1))
		if f, ok := set[key]; ok {
			f(m, append(previous, key), options[1:])
		}
	}
}
