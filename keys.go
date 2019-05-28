package main

import (
	"fmt"

	"github.com/gellel/emojipedia/lexicon"
)

func keys(name string, getter func() (*lexicon.Lexicon, error)) {
	lexicon, err := getter()
	if err != nil {
		fmt.Println(fmt.Sprintf(errorCannotOpen, name, err))
	} else {
		lexicon.Keys().Each(func(_ int, i interface{}) {
			fmt.Println(i.(string))
		})
	}
}
