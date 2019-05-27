package main

import (
	"os"
	"strings"

	"github.com/gellel/emojipedia/arguments"
)

func main() {
	arguments := arguments.NewArguments(os.Args[1:])
	switch strings.ToUpper(arguments.Get(0)) {
	case CATEGORIES:
		categoriesMain(arguments.Next())
	case EMOJI:
		emojiMain(arguments.Next())
	case EMOJIPEDIA:
		emojipediaMain(arguments.Next())
	case KEYWORDS:
		keywordsMain(arguments.Next())
	case SUBCATEGORIES:
		subcategoriesMain(arguments.Next())
	case UNICODE:
		unicodeorgMain(arguments.Next())
	default:
		os.Exit(2)
	}
}
