package main

import (
	"fmt"
	"os"

	"github.com/gellel/emojipedia/arguments"
	"github.com/gellel/emojipedia/directory"
	"github.com/gellel/emojipedia/pkg"
	"github.com/gellel/emojipedia/subcategories"
)

func subcategoriesMain(arguments *arguments.Arguments) {
	switch arguments.Get(0) {
	case "-b", "build":
		fmt.Println("attempting to build subcategories package")
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
		subcategories.Make(document)
		fmt.Println("successfully built subcategories")
	case "-g", "get":
		subcategories, err := subcategories.Open()
		if err != nil {
			fmt.Println("something went wrong when openining the subcategories package")
		}
		subcategory, ok := subcategories.Get(arguments.Next().Get(0))
		if ok != true {
			fmt.Println("subcategory does not exist")
			os.Exit(2)
		}
		subcategory.TabWriter()
	case "-l", "list":
		subcategories, err := subcategories.Open()
		if err != nil {
			fmt.Println("something went wrong when opening the subcategories package")
			fmt.Println(err)
			os.Exit(1)
		}
		//w := new(tabwriter.Writer)

		// Format in tab-separated columns with a tab stop of 8.
		//w.Init(os.Stdout, 0, 3, 0, '\t', 0)
		subcategories.List()
		//w.Flush()

	}

}
