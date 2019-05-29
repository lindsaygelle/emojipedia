package main

import (
	"fmt"
	"strings"

	"github.com/gellel/emojipedia/arguments"
	"github.com/gellel/emojipedia/category"
)

func categoryMain(arguments *arguments.Arguments) {
	category, err := category.Open(arguments.Get(0))
	if err != nil {
		fmt.Println(fmt.Sprintf(errorCannotOpen, CATEGORY, err))
	} else {
		switch strings.ToUpper(arguments.Next().Get(0)) {
		case ANCHOR:
			fmt.Println(category.Anchor)
		case EMOJI:
			category.Emoji.Sort().Each(func(_ int, i interface{}) {
				fmt.Println(i.(string))
			})
		case SUBCATEGORIES:
			category.Subcategories.Sort().Each(func(_ int, i interface{}) {
				fmt.Println(i.(string))
			})
		}
	}
}
