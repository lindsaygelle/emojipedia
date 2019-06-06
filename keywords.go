package main

import (
	"fmt"
	"strings"

	"github.com/gellel/emojipedia/arguments"
	"github.com/gellel/emojipedia/keywords"
)

func keywordsGet(arguments *arguments.Arguments) {
	var (
		keywords = keywords.Get()
	)
	fmt.Fprintln(writer, "N\t|Name\t|Emoji")
	arguments.Each(func(i int, argument string) {
		if slice, ok := keywords.Get(argument); ok {
			fmt.Fprintln(writer, fmt.Sprintf("%v\t|%v\t|%vi", i, argument, slice.Join(" ")))
		}
	})
	writer.Flush()
}

func keywordsKeys(arguments *arguments.Arguments) {
	var (
		keywords = keywords.Get()
	)
	fmt.Fprintln(writer, "N\t|Name")
	keywords.Keys().Sort().Each(func(i int, x interface{}) {
		fmt.Fprintln(writer, fmt.Sprintf("%v\t|%v", i, x.(string)))
	})
	writer.Flush()
}

func keywordsList(arguments *arguments.Arguments) {
	var (
		keywords = keywords.Get()
	)
	fmt.Fprintln(writer, "N\t|Name\t|Emoji")
	keywords.Keys().Sort().Each(func(i int, x interface{}) {
		key := x.(string)
		slice := keywords.Fetch(key)
		fmt.Fprintln(writer, fmt.Sprintf("%v\t|%v\t|%v", i, key, slice.Len()))
	})
	writer.Flush()
}

func keywordsMain(arguments *arguments.Arguments) {
	switch strings.ToUpper(arguments.Get(0)) {
	case B, BUILD:
		build(KEYWORDS, keywords.Make)
	case K, KEYS:
		keywordsKeys(arguments.Next())
	case L, LIST:
		keywordsList(arguments.Next())
	}
}
