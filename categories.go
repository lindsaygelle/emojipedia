package main

import (
	"strings"

	"github.com/gellel/emojipedia/arguments"
	"github.com/gellel/emojipedia/categories"
	"github.com/gellel/emojipedia/category"
)

func categoriesMain(arguments *arguments.Arguments) {
	switch strings.ToUpper(arguments.Get(0)) {
	case B, BUILD:
		build(CATEGORIES, categories.Make)
	case G, GET:
		get(arguments.Next().Get(0), category.Read, category.Detail)
	case K, KEYS:
		keys(CATEGORIES, categories.List)
	case L, LIST:
		list(CATEGORIES, categories.List, category.List)
	case R, REMOVE:
		remove(CATEGORIES, categories.Remove)
	}
}
