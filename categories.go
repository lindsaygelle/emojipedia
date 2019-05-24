package main

import (
	"fmt"
	"os"

	"github.com/gellel/emojipedia/arguments"
	"github.com/gellel/emojipedia/categories"
	"github.com/gellel/emojipedia/directory"
	"github.com/gellel/emojipedia/pkg"
)

func categoriesMain(arguments *arguments.Arguments) {
	switch arguments.Get(0) {
	case "-b", "build":
		fmt.Println("attempting to build categories package")
		if _, err := os.Stat(directory.Unicode); os.IsNotExist(err) {
			fmt.Println("cannot build without unicode package. please build and try again")
			os.Exit(2)
		}
		document, err := pkg.Open()
		if err != nil {
			fmt.Println("an error occurred when trying to open the unicode package.")
			fmt.Println(err)
			os.Exit(1)
		}
		categories.Make(document)
		fmt.Println("successfully built categories")
	case "-g", "get":
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
	case "-l", "list":
		categories, err := categories.Open()
		if err != nil {
			fmt.Println("something went wrong when opening the categories package")
			fmt.Println(err)
			os.Exit(1)
		}
		categories.List()
	case "-r", "-remove":

	}
}
