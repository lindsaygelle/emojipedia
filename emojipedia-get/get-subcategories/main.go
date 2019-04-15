package subcategories

import (
	"fmt"
	"strings"

	"github.com/gellel/emojipedia/manifest"

	filesubcategories "github.com/gellel/emojipedia/emojipedia-files/files-subcategories"
)

const root string = "emojipedia"

var Export = Subcategories

var Key = "SUBCATEGORIES"

var Options = []interface{}{
	All}

var empty = map[string](func()){
	"ALL": All}

func All() {
	m, err := filesubcategories.Open()
	if err != nil {
		fmt.Println(fmt.Sprintln("cannot open subcategories. has not been built or is missing."))
	} else {
		names := []string{}
		for _, k := range m {
			names = append(names, k)
		}
		fmt.Println(fmt.Sprintf("emoji subcategories: %s.", strings.Join(names, ", ")))
	}
}

func Subcategories(options ...string) {}

func Main(m *manifest.Manifest, previous, options []string) {
	if len(options) != 0 {
		key := strings.ToUpper(strings.Replace(options[0], "-", "", -1))
		if f, ok := empty[key]; ok {
			f()
		}
	}
}
