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
		case T, TABLE:
			var (
				anchor        = c.Anchor
				emoji         = fmt.Sprintf("%v", c.Emoji.Len())
				href          = c.Href
				name          = c.Name
				number        = fmt.Sprintf("%v", c.Number)
				position      = fmt.Sprintf("%v", c.Position)
				subcategories = fmt.Sprintf("%v", c.Subcategories.Len())
				template      = []string{
					anchor,
					emoji,
					href,
					name,
					number,
					position,
					subcategories}
			)
			fmt.Fprintln(writer, "anchor\t|emoji\t|href\t|name\t|number\t|position\t|subcategories")
			fmt.Fprintln(writer, strings.Join(template, "\t|"))
			writer.Flush()
		}
	default:
		fmt.Println(fmt.Sprintf(errorChoiceNotFound, arguments.Get(0), "-cc", strings.ToLower(CATEGORY)))
	}
}
