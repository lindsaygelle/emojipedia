package main

import (
	"fmt"
	"os"
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
		categories, err := categories.Open()
		if err != nil {
			fmt.Println("something went wrong when openining the categories package")
		}
		category, ok := categories.Get(arguments.Next().Get(0))
		if ok != true {
			fmt.Println("category does not exist")
			os.Exit(2)
		}
		fmt.Println(category)
	case "-k", "keys":
		categories, err := categories.Open()
		if err != nil {
			fmt.Println("something went wrong opening the categories package")
			fmt.Println(err)
			os.Exit(1)
		}
		categories.Keys().Each(func(i int, x interface{}) {
			fmt.Println(fmt.Sprintf("%v\t%v", i, x.(string)))
		})
	case L, LIST:
		list(CATEGORIES, categories.List, category.List)
	case R, REMOVE:
		remove(CATEGORIES, categories.Remove)
	}
}
