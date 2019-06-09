package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gellel/emojipedia/arguments"
	"github.com/gellel/emojipedia/slice"
)

func main() {
	arguments := arguments.NewArguments(os.Args[1:])
	switch strings.ToUpper(arguments.Get(0)) {
	case C, CATEGORIES:
		categoriesMain(arguments.Next())
	case CC, CATEGORY:
		categoryMain(arguments.Next())
	case EE, EMOJI:
		emojiMain(arguments.Next())
	case E, EMOJIPEDIA:
		emojipediaMain(arguments.Next())
	case K, KEYWORDS:
		keywordsMain(arguments.Next())
	case S, SUBCATEGORIES:
		subcategoriesMain(arguments.Next())
	case SS, SUBCATEGORY:
		subcategoryMain(arguments.Next())
	case U, UNICODE:
		unicodeorgMain(arguments.Next())
	default:
		fmt.Fprintln(writer, "usage: emojipedia [-abbreviation|verbose] <command> [args [...<args>]]")
		fmt.Fprintln(writer)
		fmt.Fprintln(writer, "Small program that scrapes unicode.org for emoji content. Parses out HTML into categorically ordered data subsets.")
		fmt.Fprintln(writer)
		fmt.Fprintln(writer, "browsing programs collection of contents")
		slice.New(copt, kopt, eopt, sopt).Each(func(_ int, i interface{}) {
			fmt.Fprintln(writer, i.(string))
		})
		fmt.Fprintln(writer)
		fmt.Fprintln(writer, "browsing specific content")
		slice.New(ccopt, eeopt, ssopt).Each(func(_ int, i interface{}) {
			fmt.Fprintln(writer, i.(string))
		})
		writer.Flush()
	}
}
