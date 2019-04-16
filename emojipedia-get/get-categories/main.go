package categories

import (
	"fmt"
	"strings"

	api "github.com/gellel/emojipedia/emojipedia-api"
	"github.com/gellel/emojipedia/manifest"
)

const root string = "emojipedia"

var Export = Categories

var Key = "CATEGORIES"

var Options = []interface{}{
	All}

var empty = map[string](func(options ...interface{})){
	"ALL": All}

func All(options ...interface{}) {
	switch t := options[0].(type) {
	case map[int]string:
		names := []string{}
		for _, k := range t {
			names = append(names, k)
		}
		fmt.Println(fmt.Sprintf("emoji categories: %s.", strings.Join(names, ", ")))
	}
}

func Categories(options ...string) {}

func Main(m *manifest.Manifest, previous, options []string) {
	if len(options) != 0 {
		key := strings.ToUpper(strings.Replace(options[0], "-", "", -1))
		if f, ok := empty[key]; ok {
			m, _ := api.GetCategories()
			f(m)
		}
	}
}
