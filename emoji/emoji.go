package emoji

import (
	"fmt"
)

var Options = []interface{}{category, id, name}

func category(name string) int {
	fmt.Println(name)
	return 0
}

func id(name string) int {
	fmt.Println(name)
	return 0
}

func name(ID int) int {
	fmt.Println(ID)
	return 0
}
