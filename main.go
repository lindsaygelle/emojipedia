package main

import (
	"os"

	"github.com/gellel/emojipedia/arguments"
)

func main() {
	arguments := arguments.NewArguments(os.Args[1:])
	switch arguments.Get(0) {
	case "-c", "categories":
		categoriesMain(arguments.Next())
	case "-s", "subcategories":
		subcategoriesMain(arguments.Next())
	case "-u", "unicode", "unicode-org":
		unicodeorgMain(arguments.Next())
	}
}
