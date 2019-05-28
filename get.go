package main

import "fmt"

func get(name string, reader func(string) (*[]byte, error), parser func(*[]byte)) {
	content, err := reader(name)
	if err != nil {
		fmt.Println(fmt.Sprintf(errorCannotOpen, name, err))
	} else {
		parser(content)
	}
}
