package categories

import (
	"fmt"
	"strings"

	c "github.com/gellel/emojipedia/emojipedia-files/files-categories"
	"github.com/gellel/emojipedia/manifest"
)

const root string = "emojipedia"

var Export = Categories

var Key = "CATEGORIES"

var Options = []interface{}{
	All}

var empty = map[string](func()){
	"ALL": All}

func All() {
	m, err := c.Open()
	if err != nil {
		fmt.Println(fmt.Sprintln("cannot open categories. has not been built or is missing."))
	} else {
		names := []string{}
		for _, k := range m {
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
			f()
		}
	}
}
