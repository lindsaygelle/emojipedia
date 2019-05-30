package main

import (
	"fmt"
	"strings"

	"github.com/gellel/emojipedia/arguments"
	"github.com/gellel/emojipedia/emoji"
	"github.com/gellel/emojipedia/emojipedia"
)

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
		get(arguments.Next().Get(0), emoji.Read, emoji.Detail)
	case K, KEYS:
		emojipediaKeys(arguments.Next())
	case L, LIST:
		emojipediaList(arguments.Next())
	case R, REMOVE:
		remove(EMOJIPEDIA, emojipedia.Remove)
	}
}
