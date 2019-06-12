package main

import (
	"fmt"
	"strings"

	"github.com/gellel/emojipedia/arguments"
	"github.com/gellel/emojipedia/categories"
	"github.com/gellel/emojipedia/slice"
	"github.com/gellel/emojipedia/stdin"
)

func categoriesGet(arguments *arguments.Arguments) {
	var (
		categories = categories.Get()
	)
	fmt.Fprintln(writer, "Name\t|Number\t|Subcategories")
	if arguments.Get(0) != "" {
		arguments.Each(func(_ int, argument string) {
			if category, ok := categories.Get(argument); ok {
				var (
					name          = category.Name
					number        = category.Number
					subcategories = category.Subcategories.Sort().Join(" ")
					output        = fmt.Sprintf("%v\t|%v\t|%v", name, number, subcategories)
				)
				fmt.Fprintln(writer, output)
			}
		})
		writer.Flush()
	}
}

func categoriesKeys(arguments *arguments.Arguments) {
	var (
		categories = categories.Get()
	)
	fmt.Fprintln(writer, "N\t|Name")
	categories.Keys().Sort().Each(func(i int, x interface{}) {
		fmt.Fprintln(writer, fmt.Sprintf("%v\t|%v", i, x.(string)))
	})
	writer.Flush()
}

func categoriesList(arguments *arguments.Arguments) {
	var (
		categories = categories.Get()
	)
	fmt.Fprintln(writer, "Name\t|Number\t|Emoji\t|Subcategories")
	categories.Keys().Sort().Each(func(_ int, i interface{}) {
		var (
			category      = categories.Fetch(i.(string))
			name          = category.Name
			number        = category.Number
			emoji         = category.Emoji.Len()
			subcategories = category.Subcategories.Len()
			output        = fmt.Sprintf("%v\t|%v\t|%v\t|%v", name, number, emoji, subcategories)
		)
		fmt.Fprintln(writer, output)
	})
	writer.Flush()
}

func categoriesNumber(arguments *arguments.Arguments) {
	var (
		categories = categories.Get()
	)
	fmt.Fprintln(writer, "Categories\t|Number")
	fmt.Fprintln(writer, fmt.Sprintf("\t|%v", categories.Len()))
	writer.Flush()
}

func categoriesMain(arguments *arguments.Arguments) {
	switch strings.ToUpper(arguments.Get(0)) {
	case B, BUILD:
		build(CATEGORIES, categories.Make)
	case G, GET:
		categoriesGet(arguments.Next())
	case K, KEYS:
		categoriesKeys(arguments.Next())
	case L, LIST:
		categoriesList(arguments.Next())
	case N, NUMBER:
		categoriesNumber(arguments.Next())
	case R, REMOVE:
		remove(CATEGORIES, categories.Remove)
	default:
		var (
			b = stdin.Arg{
				About:   "create the categories",
				Short:   B,
				Verbose: BUILD}
			g = stdin.Arg{
				About:   "get one or more categories",
				Short:   G,
				Verbose: GET}
			k = stdin.Arg{
				About:   "show available category choices",
				Short:   K,
				Verbose: KEYS}
			l = stdin.Arg{
				About:   "iterate and show the available categories information",
				Short:   L,
				Verbose: LIST}
			n = stdin.Arg{
				About:   "number of categories in package",
				Short:   N,
				Verbose: NUMBER}
			r = stdin.Arg{
				About:   "remove the categories (all)",
				Short:   R,
				Verbose: REMOVE}
		)
		fmt.Fprintln(writer, "usage: emojipedia [-c categories] [<option>] [--flags]")
		fmt.Fprintln(writer)
		fmt.Fprintln(writer, "installing categories")
		fmt.Fprintln(writer, b)
		fmt.Fprintln(writer)
		fmt.Fprintln(writer, "removing categories")
		fmt.Fprintln(writer, r)
		fmt.Fprintln(writer)
		fmt.Fprintln(writer, "options that support flags")
		slice.New(g, k, l, n).Each(func(_ int, i interface{}) {
			fmt.Fprintln(writer, i.(stdin.Arg))
		})
		fmt.Fprintln(writer)
		writer.Flush()
	}
}
