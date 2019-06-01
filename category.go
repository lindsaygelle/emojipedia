package main

import (
	"fmt"
	"strings"

	"github.com/gellel/emojipedia/arguments"
	"github.com/gellel/emojipedia/category"
)

func categoryMain(arguments *arguments.Arguments) {
	c, err := category.Open(arguments.Get(0))
	switch err == nil {
	case true:
		switch strings.ToUpper(arguments.Next().Get(0)) {
		case "-A", ANCHOR:
			fmt.Println(c.Anchor)
		case "-E", EMOJI:
			c.Emoji.Sort().Each(func(_ int, i interface{}) {
				fmt.Println(i.(string))
			})
		case "-H", HREF:
			fmt.Println(c.Href)
		case "-P", POSITION:
			fmt.Println(c.Position)
		case "-N", NUMBER:
			fmt.Println(c.Number)
		case "-S", SUBCATEGORIES:
			c.Subcategories.Sort().Each(func(_ int, i interface{}) {
				fmt.Println(i.(string))
			})
		}
	default:
		fmt.Println(fmt.Sprintf("%s \"%s\" not found", CATEGORY, arguments.Get(0)))
	}
}
