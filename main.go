package main

import (
	"fmt"
	"os"

	"github.com/gellel/emojipedia/emojipedia"

	unicodes "github.com/gellel/emojipedia/unicode-org"
)

func main() {
	var (
		args     = (&emojipedia.Strings{}).New(os.Args[1:]...)
		argument = args.Peek(0)
	)
	switch argument {
	case "", "-h", "--help":
	case "-v", "--version":
		fmt.Println(emojipedia.VersionString)
	case "-a", "--about":
		fmt.Println("about")
	case emojipedia.CategorizationKey:
		//categories.Go(args.Drop(0))
	case emojipedia.EmojiKey:
		//emoji.Go(args.Drop(0))
	case emojipedia.KeywordsKey:
		//keywords.Go(args.Drop(0))
	case emojipedia.SubcategorizationKey:
		//subcategories.Go(args.Drop(0))
	case emojipedia.UnicodeKey:
		unicodes.Go(args.Drop(0))
	default:
		fmt.Println(fmt.Sprintf(emojipedia.ErrorArgumentTemplate, argument))
	}
}
