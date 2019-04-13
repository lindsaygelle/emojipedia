package all

import (
	"fmt"

	"github.com/gellel/emojipedia/emojipedia"
	"github.com/gellel/emojipedia/manifest"
)

var Export = All

var Key = "ALL"

var Options = []interface{}{ID, Name}

var emojidex map[string]*emojipedia.Emoji

func All(options ...string) {}

func ID(ID int) {}

func Name(name string) {}

func Main(m *manifest.Manifest, previous, options []string) {
	fmt.Println(previous)
}
