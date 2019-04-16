package files

import (
	"strings"

	"github.com/gellel/emojipedia/manifest"

	categories "github.com/gellel/emojipedia/emojipedia-files/files-categories"
	emojis "github.com/gellel/emojipedia/emojipedia-files/files-emojis"
	keywords "github.com/gellel/emojipedia/emojipedia-files/files-keywords"
	subcategories "github.com/gellel/emojipedia/emojipedia-files/files-subcategories"
)

var Key = "FILES"

var programs = map[string](func(m *manifest.Manifest, previous, options []string)){
	categories.Key:    categories.Main,
	emojis.Key:        emojis.Main,
	keywords.Key:      keywords.Main,
	subcategories.Key: subcategories.Main}

func Main(m *manifest.Manifest, previous, options []string) {
	var argument string
	if len(options) != 0 {
		argument = strings.ToUpper(options[0])
	}
	if f, ok := programs[argument]; ok {
		f(m, append(previous, argument), options[1:])
	}
}
