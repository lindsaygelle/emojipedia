package main

import (
	"fmt"
	"os"

	"github.com/gellel/emojipedia/arguments"
)

func main() {
	arguments := arguments.NewArguments(os.Args[1:])
	switch arguments.Get(0) {
	case "-c", "categories":
		categoriesMain(arguments.Next())
	case "-e", "emoji":
		emojiMain(arguments.Next())
	case "-k", "keywords":
		keywordsMain(arguments.Next())
	case "-s", "subcategories":
		subcategoriesMain(arguments.Next())
	case "-u", "unicode", "unicode-org":
		unicodeorgMain(arguments.Next())
	default:
		fmt.Println("command not recognised. please try again")
		os.Exit(2)
	}
}
