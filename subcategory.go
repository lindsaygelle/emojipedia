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
		case "-N", NUMBER:
			fmt.Println(s.Number)
		case "-P", POSITION:
			fmt.Println(s.Position)
		case T, TABLE:
			var (
				anchor   = s.Anchor
				category = s.Category
				emoji    = fmt.Sprintf("%v", s.Emoji.Len())
				href     = s.Href
				name     = s.Name
				number   = fmt.Sprintf("%v", s.Number)
				position = fmt.Sprintf("%v", s.Position)
				template = []string{
					anchor,
					category,
					emoji,
					href,
					name,
					number,
					position}
			)
			fmt.Fprintln(writer, "anchor\t|category\t|emoji\t|href\t|name\t|number\t|position")
			fmt.Fprintln(writer, strings.Join(template, "\t|"))
			writer.Flush()
		}
	default:
		fmt.Println(fmt.Sprintf(errorChoiceNotFound, arguments.Get(0), "-ss", strings.ToLower(SUBCATEGORY)))
	}
}
