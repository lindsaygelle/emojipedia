package main

import (
	"fmt"

	"github.com/gellel/emojipedia/emojipedia"
)

var (
	unicodeOrgHTTPRequestMessage = "emojipedia: status. program is attempting to make a request to \"%s\".\ncontent response is quite large and may take awhile."
	unicodeOrgHTTPStoreMessage   = "emojipedia: status. program storing downloaded package \"%s\" to directory \"%s\".\nthank you for being patient."
)

func unicodeorg(args *emojipedia.Strings) {
	argument := args.Peek(0)
	switch argument {
	case "", "-h", "--help":
		fmt.Println("help")
	case "-a", "--about":
		fmt.Println("about")
	case "get":
		unicodeorgBuild(args.Drop(0))
	default:
		fmt.Println(fmt.Sprintf(emojipedia.ErrorArgumentTemplate, argument))
	}
}

func unicodeorgBuild(args *emojipedia.Strings) {
	verbose := (args.Peek(0) == "-v") || (args.Peek(0) == "--verbose")
	if verbose {
		fmt.Println(unicodeOrgHTTPRequestMessage)
	}
	//dump, ok := emojipedia.NewUnicodeOrgHTMLDump()
	//if ok != true {

	//}
	if verbose {
		fmt.Println(unicodeOrgHTTPStoreMessage)
	}
	//ok = emojipedia.StoreUnicodeOrgFileAsHTML(&dump)
	if verbose {
		fmt.Println("x")
	}
}
