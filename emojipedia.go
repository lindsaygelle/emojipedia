package main

import (
	"fmt"
	"strings"

	"github.com/gellel/emojipedia/arguments"
	"github.com/gellel/emojipedia/emojipedia"
	"github.com/gellel/emojipedia/slice"
	"github.com/gellel/emojipedia/stdin"
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

func emojipediaNumber(arguments *arguments.Arguments) {
	var (
		emojipedia = emojipedia.Get()
	)
	fmt.Fprintln(writer, "Emojipedia\t|Number")
	fmt.Fprintln(writer, fmt.Sprintf("\t|%v", emojipedia.Len()))
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
	case N, NUMBER:
		emojipediaNumber(arguments.Next())
	case R, REMOVE:
		remove(EMOJIPEDIA, emojipedia.Remove)
	default:
		var (
			b = stdin.Arg{
				About:   "create the emojipedia",
				Short:   B,
				Verbose: BUILD}
			g = stdin.Arg{
				About:   "get one or more emoji",
				Short:   G,
				Verbose: GET}
			k = stdin.Arg{
				About:   "show available emoji choices",
				Short:   K,
				Verbose: KEYS}
			l = stdin.Arg{
				About:   "iterate and show the available emoji information",
				Short:   L,
				Verbose: LIST}
			n = stdin.Arg{
				About:   "number of emoji",
				Short:   N,
				Verbose: NUMBER}
			r = stdin.Arg{
				About:   "remove the emojipedia (all)",
				Short:   R,
				Verbose: REMOVE}
		)
		fmt.Fprintln(writer, "usage: emojipedia [-e emojipedia] [<option>] [--flags]")
		fmt.Fprintln(writer)
		fmt.Fprintln(writer, "installing emojipedia")
		fmt.Fprintln(writer, b)
		fmt.Fprintln(writer)
		fmt.Fprintln(writer, "removing emojipedia")
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
