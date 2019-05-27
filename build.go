package main

import (
	"fmt"
	"os"

	"github.com/PuerkitoBio/goquery"
	"github.com/gellel/emojipedia/directory"
	"github.com/gellel/emojipedia/pkg"
)

func build(name string, f func(document *goquery.Document)) {
	fmt.Println(fmt.Sprintf(statusBuildPackage, name))
	if _, err := os.Stat(directory.Unicode); os.IsNotExist(err) {
		fmt.Println(fmt.Sprintf(errorCannotFind, "unicode"))
		os.Exit(2)
	}
	document, err := pkg.Open()
	if err != nil {
		fmt.Println(fmt.Sprintf(errorCannotOpen, "unicode", err))
		os.Exit(1)
	}
	f(document)
	fmt.Println(fmt.Sprintf("successfully built %s", name))
	os.Exit(0)
}
