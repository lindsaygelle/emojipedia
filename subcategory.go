package main

import (
	"fmt"
	"strings"

	"github.com/gellel/emojipedia/arguments"
	"github.com/gellel/emojipedia/subcategory"
)

func subcategoryMain(arguments *arguments.Arguments) {
	s, err := subcategory.Open(arguments.Get(0))
	switch err == nil {
	case true:
		switch strings.ToUpper(arguments.Next().Get(0)) {
		case "-A", ANCHOR:
			fmt.Println(s.Anchor)
		case "-C", CATEGORY:
			fmt.Println(s.Category)
		case "-E", EMOJI:
			s.Emoji.Sort().Each(func(_ int, i interface{}) {
				fmt.Println(i.(string))
			})
		case "-H", HREF:
			fmt.Println(s.Href)
		case "-P", POSITION:
			fmt.Println(s.Position)
		case "-N", NUMBER:
			fmt.Println(s.Number)
		}
	}
}
