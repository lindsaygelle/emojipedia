package categories

import (
	"fmt"
	"strings"

	f "github.com/gellel/emojipedia/emojipedia-files/files-keywords"
	"github.com/gellel/emojipedia/manifest"
)

const root string = "emojipedia"

var Export = Keywords

var Key = "KEYWORDS"

var Options = []interface{}{
	All}

var empty = map[string](func()){
	"ALL": All}

func All() {
	m, err := f.Open()
	if err != nil {
		fmt.Println(fmt.Sprintln("cannot open keywords. has not been built or is missing."))
	} else {
		for key := range m {
			fmt.Println(key)
		}
	}
}

func Keywords(options ...string) {}

func Main(m *manifest.Manifest, previous, options []string) {
	if len(options) != 0 {
		key := strings.ToUpper(strings.Replace(options[0], "-", "", -1))
		if f, ok := empty[key]; ok {
			f()
		}
	}
}
