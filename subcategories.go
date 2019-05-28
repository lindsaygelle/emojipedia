package main

import (
	"strings"

	"github.com/gellel/emojipedia/arguments"
	"github.com/gellel/emojipedia/subcategories"
	"github.com/gellel/emojipedia/subcategory"
)

func subcategoriesMain(arguments *arguments.Arguments) {
	switch strings.ToUpper(arguments.Get(0)) {
	case B, BUILD:
		build(SUBCATEGORIES, subcategories.Make)
	case G, GET:
	case K, KEYS:
		keys(SUBCATEGORIES, subcategories.List)
	case L, LIST:
		list(SUBCATEGORIES, subcategories.List, subcategory.List)
	case R, REMOVE:
		remove(SUBCATEGORIES, subcategories.Remove)
	}
}
