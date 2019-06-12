package main

import (
	"fmt"
	"strings"

	"github.com/gellel/emojipedia/slice"
	"github.com/gellel/emojipedia/stdin"

	"github.com/gellel/emojipedia/arguments"
	"github.com/gellel/emojipedia/category"
)

func categoryMain(arguments *arguments.Arguments) {
	c, err := category.Open(arguments.Get(0))
	switch err == nil {
	case true:
		switch strings.ToUpper(arguments.Next().Get(0)) {
		case A, ANCHOR:
			fmt.Println(c.Anchor)
		case E, EMOJI:
			c.Emoji.Sort().Each(func(_ int, i interface{}) {
				fmt.Println(i.(string))
			})
		case H, HREF:
			fmt.Println(c.Href)
		case N, NUMBER:
			fmt.Println(c.Number)
		case P, POSITION:
			fmt.Println(c.Position)
		case S, SUBCATEGORIES:
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
		default:
			var (
				a = stdin.Arg{"get the category href", A, ANCHOR}
				e = stdin.Arg{"show all emoji (list)", E, EMOJI}
				h = stdin.Arg{"get the full emoji category URL", H, HREF}
				n = stdin.Arg{"get the categorical number", N, NUMBER}
				p = stdin.Arg{"show the position the category was parsed", P, POSITION}
				s = stdin.Arg{"show all subcategories for category (list)", S, SUBCATEGORIES}
				t = stdin.Arg{"table the category", T, TABLE}
			)
			fmt.Fprintln(writer, fmt.Sprintf("usage: emojipedia [-cc category] %s [<option>] [--flags]", c.Name))
			fmt.Fprintln(writer)
			slice.New(a, e, h, n, p, s, t).Each(func(_ int, i interface{}) {
				fmt.Fprintln(writer, i.(stdin.Arg))
			})
			fmt.Fprintln(writer)
			writer.Flush()
		}
	default:
		fmt.Println(fmt.Sprintf(errorChoiceNotFound, arguments.Get(0), "-cc", strings.ToLower(CATEGORY)))
	}
}
