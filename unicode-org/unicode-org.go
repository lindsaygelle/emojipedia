package org

import (
	"fmt"

	"github.com/gellel/emojipedia/emojipedia"
)

func Go(args *emojipedia.Strings) {
	argument := args.Peek(0)
	switch argument {
	case "", "-h", "--help":
	case "-a", "--about":
		fmt.Println("about")
	case "save":
		fmt.Println("save!")
	default:
		fmt.Println(fmt.Sprintf(emojipedia.ErrorArgumentTemplate, argument))
	}
}
