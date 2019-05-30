package main

import (
	"fmt"
	"strings"

	"github.com/gellel/emojipedia/arguments"
	"github.com/gellel/emojipedia/subcategories"
)

func subcategoriesGet(arguments *arguments.Arguments) {
	var (
		subcategories = subcategories.Get()
	)
	fmt.Fprintln(writer, "Name\t|Number\t|Category")
	arguments.Each(func(_ int, argument string) {
		if subcategory, ok := subcategories.Get(argument); ok {
			var (
				name     = subcategory.Name
				number   = subcategory.Number
				category = subcategory.Category
				output   = fmt.Sprintf("%v\t|%v\t|%v", name, number, category)
			)
			fmt.Fprintln(writer, output)
		}
	})
	writer.Flush()
}

func subcategoriesKeys(arguments *arguments.Arguments) {
	var (
		subcategories = subcategories.Get()
	)
	fmt.Fprintln(writer, "N\t|Name")
	subcategories.Keys().Sort().Each(func(i int, x interface{}) {
		fmt.Fprintln(writer, fmt.Sprintf("%v\t|%v", i, x.(string)))
	})
	writer.Flush()
}

func subcategoriesList(arguments *arguments.Arguments) {
	var (
		subcategories = subcategories.Get()
	)
	fmt.Fprintln(writer, "Name\t|Number\t|Category\t|Emoji")
	subcategories.Keys().Sort().Each(func(_ int, i interface{}) {
		var (
			subcategory = subcategories.Fetch(i.(string))
			name        = subcategory.Name
			number      = subcategory.Number
			category    = subcategory.Category
			emoji       = subcategory.Emoji.Len()
			output      = fmt.Sprintf("%v\t|%v\t|%v\t|%v", name, number, category, emoji)
		)
		fmt.Fprintln(writer, output)
	})
	writer.Flush()
}

func subcategoriesMain(arguments *arguments.Arguments) {
	switch strings.ToUpper(arguments.Get(0)) {
	case B, BUILD:
		build(SUBCATEGORIES, subcategories.Make)
	case G, GET:
		subcategoriesGet(arguments.Next())
	case K, KEYS:
		subcategoriesKeys(arguments.Next())
	case L, LIST:
		subcategoriesList(arguments.Next())
	case R, REMOVE:
		remove(SUBCATEGORIES, subcategories.Remove)
	}
}
