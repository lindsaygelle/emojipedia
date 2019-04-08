package character

import (
	"fmt"
	"strings"
)

var Options = []interface{}{byId, byName}

func byId(name string) int {
	fmt.Println(name)
	return 0
}

func byName(ID int) int {
	fmt.Println(ID)
	return 0
}

func Main(options []string) {
	switch strings.ToLower(options[0]) {
	case "by-id":
		byId("hello")
	}
}
