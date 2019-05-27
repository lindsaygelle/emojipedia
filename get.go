package main

import "fmt"

func get(name string, f func() (interface{}, error)) {
	i, err := f()
	if err != nil {
		fmt.Println(fmt.Sprintf(errorCannotOpen, name, err))
	} else {
		fmt.Println(i)
	}
}
