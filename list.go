package main

import (
	"fmt"
	"os"
	"text/tabwriter"

	"github.com/gellel/emojipedia/lexicon"
)

func list(name string, getter func() (*lexicon.Lexicon, error), printer func(w *tabwriter.Writer, i interface{})) {
	lexicon, err := getter()
	switch err == nil {
	case true:
		writer := new(tabwriter.Writer)
		writer.Init(os.Stdout, 0, 8, 0, '\t', 0)
		lexicon.Keys().Sort().Each(func(_ int, i interface{}) {
			printer(writer, lexicon.Fetch(i.(string)))
		})
		writer.Flush()
	default:
		fmt.Println(err)
	}
}
