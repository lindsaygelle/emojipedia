package main

import (
	"fmt"
	"strings"

	"github.com/gellel/emojipedia/slice"

	"github.com/gellel/emojipedia/arguments"
	"github.com/gellel/emojipedia/emojipedia"
	"github.com/gellel/emojipedia/text"
)

func emojipediaGet(arguments *arguments.Arguments) {
	var (
		emojipedia = emojipedia.Get()
	)
	fmt.Fprintln(writer, "\t|Name\t|Number\t|Category\t|Subcategory\t|Keywords")
	arguments.Each(func(_ int, argument string) {
		if emoji, ok := emojipedia.Get(argument); ok {
			var (
				character   = text.Emojize(emoji.Unicode)
				name        = emoji.Name
				number      = emoji.Number
				category    = emoji.Category
				subcategory = emoji.Subcategory
				keywords    = emoji.Keywords.Sort().Join(" ")
				output      = fmt.Sprintf("%v\t|%v\t|%v\t|%v\t|%v\t|%v", character, name, number, category, subcategory, keywords)
			)
			fmt.Fprintln(writer, output)
		}
	})
	writer.Flush()
}

func emojipediaKeys(arguments *arguments.Arguments) {
	var (
		emojipedia = emojipedia.Get()
	)
	fmt.Fprintln(writer, "N\t|Name")
	emojipedia.Keys().Sort().Each(func(i int, x interface{}) {
		fmt.Fprintln(writer, fmt.Sprintf("%v\t|%v", i, x.(string)))
	})
	writer.Flush()
}

func emojipediaList(arguments *arguments.Arguments) {
	var (
		emojipedia = emojipedia.Get()
	)
	fmt.Fprintln(writer, "Name\t|Number\t|Category\t|Subcategory\t|Keywords")
	emojipedia.Keys().Sort().Each(func(_ int, i interface{}) {
		var (
			emoji       = emojipedia.Fetch(i.(string))
			name        = emoji.Name
			number      = emoji.Number
			category    = emoji.Category
			subcategory = emoji.Subcategory
			keywords    = emoji.Keywords.Len()
			output      = fmt.Sprintf("%v\t|%v\t|%v\t|%v\t|%v", name, number, category, subcategory, keywords)
		)
		fmt.Fprintln(writer, output)
	})
	writer.Flush()
}

func emojipediaMain(arguments *arguments.Arguments) {
	switch strings.ToUpper(arguments.Get(0)) {
	case B, BUILD:
		build(EMOJIPEDIA, emojipedia.Make)
	case G, GET:
		emojipediaGet(arguments.Next())
	case K, KEYS:
		emojipediaKeys(arguments.Next())
	case L, LIST:
		emojipediaList(arguments.Next())
	case R, REMOVE:
		remove(EMOJIPEDIA, emojipedia.Remove)
	default:
		var (
			t = "\t\t%s,%s\t%s"
			b = fmt.Sprintf(t, B, BUILD, "build the emojipedia")
			g = fmt.Sprintf(t, G, GET, "get an emoji or emoji set from the encyclopedia")
			k = fmt.Sprintf(t, K, KEYS, "print out the available keys")
			l = fmt.Sprintf(t, L, LIST, "iterate over the stored emoji sets")
		)
		slice.New(b, g, k, l).Each(func(_ int, i interface{}) {
			fmt.Fprintln(writer, i.(string))
		})
		writer.Flush()
	}
}
