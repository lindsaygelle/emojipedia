package get

import (
	"strings"

	categories "github.com/gellel/emojipedia/emojipedia-get/get-categories"
	emoji "github.com/gellel/emojipedia/emojipedia-get/get-emoji"
	keywords "github.com/gellel/emojipedia/emojipedia-get/get-keywords"
	subcategories "github.com/gellel/emojipedia/emojipedia-get/get-subcategories"
	"github.com/gellel/emojipedia/manifest"
)

var Exports = Get

var Key = "GET"

var Options = []interface{}{
	categories.Export,
	emoji.Export,
	keywords.Export,
	subcategories.Export}

var set = map[string](func(m *manifest.Manifest, previous, options []string)){
	categories.Key:    categories.Main,
	emoji.Key:         emoji.Main,
	keywords.Key:      keywords.Main,
	subcategories.Key: subcategories.Main}

var m *manifest.Manifest

var s *manifest.Program

func Get(options ...string) {}

func Main(m *manifest.Manifest, previous []string, options []string) {
	if len(options) != 0 {
		key := strings.ToUpper(strings.Replace(options[0], "-", "", -1))
		if f, ok := set[key]; ok {
			f(m, append(previous, key), options[1:])
		}
	}
}
