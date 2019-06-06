package main

import (
	"fmt"
	"strings"

	"github.com/gellel/emojipedia/arguments"
	"github.com/gellel/emojipedia/keywords"
)

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

func keywordsMain(arguments *arguments.Arguments) {
	switch strings.ToUpper(arguments.Get(0)) {
	case B, BUILD:
		build(KEYWORDS, keywords.Make)
	case K, KEYS:
		keywordsKeys(arguments.Next())
	}
}
