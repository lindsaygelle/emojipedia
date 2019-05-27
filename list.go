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
		lexicon.Each(func(key string, i interface{}) {
			printer(writer, i)
		})
		writer.Flush()
	default:
		fmt.Println(err)
	}
}
