package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gellel/emojipedia/arguments"
	"github.com/gellel/emojipedia/subcategory"
)

func subcategoryMain(arguments *arguments.Arguments) {
	subcategory, err := subcategory.Open(arguments.Get(0))
	if err != nil {
		fmt.Println(fmt.Sprintf(errorCannotOpen, SUBCATEGORY, err))
		os.Exit(1)
	}
	switch strings.ToUpper(arguments.Next().Get(0)) {
	case ANCHOR:
		fmt.Println(subcategory.Anchor)
	case CATEGORY:
		fmt.Println(subcategory.Category)
	case EMOJI:
		subcategory.Emoji.Sort().Each(func(_ int, i interface{}) {
			fmt.Println(i.(string))
		})
	case HREF:
		fmt.Println(subcategory.Href)
	case NUMBER:
		fmt.Println(subcategory.Number)
	}
}
