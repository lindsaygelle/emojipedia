package main

import (
	"fmt"
	"os"
	"strings"

	"github.com/gellel/emojipedia/directory"

	"github.com/gellel/emojipedia/arguments"
	"github.com/gellel/emojipedia/pkg"
)

func unicodeorgMain(arguments *arguments.Arguments) {
	switch strings.ToUpper(arguments.Get(0)) {
	case B, BUILD:
		fmt.Println("attempting to build unicode-org package.")
		if _, err := os.Stat(directory.Unicode); os.IsExist(err) {
			fmt.Println("already built. nothing to do.")
			os.Exit(0)
		}
		fmt.Println("must collect package. making http request. can take awhile.")
		response, err := pkg.HTTP()
		if err != nil {
			fmt.Println("cannot collect content. encountered error.")
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("http request succeeded. attempting to store.")
		err = pkg.Write(response)
		if err != nil {
			fmt.Println("unable to store content. error occurred.")
			fmt.Println(err)
			os.Exit(1)
		}
		fmt.Println("successfully stored content.")
		fmt.Println(directory.Unicode)
		os.Exit(0)
	case R, REMOVE:
		remove(UNICODE, pkg.Remove)
	}
}
