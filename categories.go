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
		)
		fmt.Fprintln(writer, "usage: emojipedia categories [<option>] [--flags]")
		fmt.Fprintln(writer)
		slice.New(b, g, k, l).Each(func(_ int, i interface{}) {
			fmt.Fprintln(writer, i.(stdin.Arg))
		})
		fmt.Fprintln(writer)
	}
}
