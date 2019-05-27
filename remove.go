package main

import "fmt"

func remove(name string, remover func() error) {
	fmt.Println(fmt.Sprintf(statusRemovePackage, name))
	err := remover()
	if err != nil {
		fmt.Println(fmt.Sprintf(errorRemovePackage, name, err))
	} else {
		fmt.Println(fmt.Sprintf(successRemovePackage, name))
	}
}
