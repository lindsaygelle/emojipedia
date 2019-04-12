package subcategories

import "github.com/gellel/emojipedia/manifest"

var Call = Main

var Key = "SUBCATEGORIES"

var Export = Subcategories

func Subcategories(options ...string) {}

func Main(m *manifest.Manifest, previous, options []string) {}
