package main

import (
	"fmt"

	"github.com/gellel/emojipedia/lexicon"
)

func list(name string, getter func() (*lexicon.Lexicon, error), printer func(*lexicon.Lexicon)) {
	lexicon, err := getter()
	switch err == nil {
	case true:
		printer(lexicon)
	default:
		fmt.Println(err)
	}
}
