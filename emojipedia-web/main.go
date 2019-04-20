package web

import (
	"strings"

	"github.com/gellel/emojipedia/manifest"

	emojipedia "github.com/gellel/emojipedia/emojipedia-web/web-emojipedia"
	unicode "github.com/gellel/emojipedia/emojipedia-web/web-unicode"
)

var Export = web

var programs = map[string](func(m *manifest.Manifest, previous, options []string)){
	emojipedia.Key: emojipedia.Main,
	unicode.Key:    unicode.Main}

func main(m *manifest.Manifest, previous, options []string) {
	var argument string
	if len(options) != 0 {
		argument = strings.ToUpper(options[0])
	}
	if f, ok := programs[argument]; ok {
		f(m, append(previous, argument), options[1:])
	}
}

func web(arguments ...interface{}) {
	main(arguments[0].(*manifest.Manifest), arguments[1].([]string), arguments[2].([]string))
}
