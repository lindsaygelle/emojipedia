package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gellel/emojipedia/arguments"
)

func main() {
	arguments := arguments.NewArguments(os.Args[1:])
	switch strings.ToUpper(arguments.Get(0)) {
	case "-C", CATEGORIES:
		categoriesMain(arguments.Next())
	case "-CC", CATEGORY:
		categoryMain(arguments.Next())
	case "-EE", EMOJI:
		emojiMain(arguments.Next())
	case "-E", EMOJIPEDIA:
		emojipediaMain(arguments.Next())
	case "-K", KEYWORDS:
		keywordsMain(arguments.Next())
	case "-S", SUBCATEGORIES:
		subcategoriesMain(arguments.Next())
	case "-SS", SUBCATEGORY:
		subcategoryMain(arguments.Next())
	case UNICODE:
		unicodeorgMain(arguments.Next())
	default:
		fmt.Println("usage: emojipedia <command> [<args>]")
	}
}
