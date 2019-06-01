package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gellel/emojipedia/arguments"
	"github.com/gellel/emojipedia/category"
)

func categoryMain(arguments *arguments.Arguments) {
	category, err := category.Open(arguments.Get(0))
	if err != nil {
		fmt.Println(fmt.Sprintf(errorCannotOpen, CATEGORY, err))
		os.Exit(1)
	}
	switch strings.ToUpper(arguments.Next().Get(0)) {
	case "-A", ANCHOR:
		fmt.Println(category.Anchor)
	case "-E", EMOJI:
		category.Emoji.Sort().Each(func(_ int, i interface{}) {
			fmt.Println(i.(string))
		})
	case "-H", HREF:
		fmt.Println(category.Href)
	case "-N", NUMBER:
		fmt.Println(category.Number)
	case "-S", SUBCATEGORIES:
		category.Subcategories.Sort().Each(func(_ int, i interface{}) {
			fmt.Println(i.(string))
		})
	}
}
