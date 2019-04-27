package main

import (
	"fmt"
	"os"

	"github.com/gellel/emojipedia/emojipedia"
	"github.com/gellel/emojipedia/x"
)

func main() {
	fmt.Println(x.Help())
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
	case emojipedia.UnicodeKey:
		unicodeorg(args.Drop(0))
	default:
		fmt.Println(fmt.Sprintf(emojipedia.ErrorArgumentTemplate, argument))
	}
}
