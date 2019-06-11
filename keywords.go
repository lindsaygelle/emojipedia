package main

import (
	"fmt"
	"strings"

	"github.com/gellel/emojipedia/arguments"
	"github.com/gellel/emojipedia/keywords"
	"github.com/gellel/emojipedia/slice"
	"github.com/gellel/emojipedia/stdin"
)

func keywordsGet(arguments *arguments.Arguments) {
	var (
		keywords = keywords.Get()
	)
	fmt.Fprintln(writer, "N\t|Name\t|Emoji")
	arguments.Each(func(i int, argument string) {
		if slice, ok := keywords.Get(argument); ok {
			fmt.Fprintln(writer, fmt.Sprintf("%v\t|%v\t|%v", i, argument, slice.Join(" ")))
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
	case G, GET:
		keywordsGet(arguments.Next())
	case K, KEYS:
		keywordsKeys(arguments.Next())
	case L, LIST:
		keywordsList(arguments.Next())
	default:
		var (
			b = stdin.Arg{
				About:   "create the keywords",
				Short:   B,
				Verbose: BUILD}
			g = stdin.Arg{
				About:   "get one or more keywords",
				Short:   G,
				Verbose: GET}
			k = stdin.Arg{
				About:   "show available keyword choices",
				Short:   K,
				Verbose: KEYS}
			l = stdin.Arg{
				About:   "iterate and show the available keywords information",
				Short:   L,
				Verbose: LIST}
			r = stdin.Arg{
				About:   "remove the keywords (all)",
				Short:   R,
				Verbose: REMOVE}
		)
		fmt.Fprintln(writer, "usage: emojipedia [-k keywords] [<option>] [--flags]")
		fmt.Fprintln(writer)
		fmt.Fprintln(writer, "installing keywords")
		fmt.Fprintln(writer, b)
		fmt.Fprintln(writer)
		fmt.Fprintln(writer, "removing keywords")
		fmt.Fprintln(writer, r)
		fmt.Fprintln(writer)
		fmt.Fprintln(writer, "options that support flags")
		slice.New(g, k, l).Each(func(_ int, i interface{}) {
			fmt.Fprintln(writer, i.(stdin.Arg))
		})
		fmt.Fprintln(writer)
		writer.Flush()
	}
}
