package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gellel/emojipedia/arguments"
	"github.com/gellel/emojipedia/emoji"
	"github.com/gellel/emojipedia/text"
)

func emojiMain(arguments *arguments.Arguments) {
	emoji, err := emoji.Open(arguments.Get(0))
	if err != nil {
		fmt.Println(fmt.Sprintf(errorCannotOpen, EMOJI, err))
		os.Exit(1)
	}
	switch strings.ToUpper(arguments.Next().Get(0)) {
	case ANCHOR:
		fmt.Println(emoji.Anchor)
	case CATEGORY:
		fmt.Println(emoji.Category)
	case CODES:
		emoji.Codes.Each(func(_ int, i interface{}) {
			fmt.Println(i.(string))
		})
	case DESCRIPTION:
		if len(emoji.Description) == 0 {
			fmt.Println("emoji needs to be updated for description to work")
		}
	case EMOJI:
		fmt.Println(text.Emojize(emoji.Unicode))
	case HREF:
		fmt.Println(emoji.Href)
	case IMAGE:
		fmt.Println(emoji.Image)
	case KEYWORDS:
		emoji.Keywords.Sort().Each(func(_ int, i interface{}) {
			fmt.Println(i.(string))
		})
	case NUMBER:
		fmt.Println(emoji.Number)
	case SUBCATEGORY:
		fmt.Println(emoji.Subcategory)
	case UNICODE:
		fmt.Println(emoji.Unicode)
	}
}
