package files

import (
	"strings"

	categories "github.com/gellel/emojipedia/emojipedia-files/files-categories"
	keywords "github.com/gellel/emojipedia/emojipedia-files/files-keywords"
	subcategories "github.com/gellel/emojipedia/emojipedia-files/files-subcategories"
	"github.com/gellel/emojipedia/manifest"
)

var Exports = func(options ...string) {}

var Key = "FILES"

var Options = []interface{}{
	categories.Export,
	keywords.Export,
	subcategories.Export}

var set = map[string](func(m *manifest.Manifest, previous, options []string)){
	categories.Key:    categories.Main,
	keywords.Key:      keywords.Main,
	subcategories.Key: subcategories.Main}

func Main(m *manifest.Manifest, previous, options []string) {
	if len(options) != 0 {
		key := strings.ToUpper(strings.Replace(options[0], "-", "", -1))
		if f, ok := set[key]; ok {
			f(m, append(previous, key), options[1:])
		}
	}
}
