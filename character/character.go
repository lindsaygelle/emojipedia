package character

import (
	"fmt"
	"strings"
)

var Options = []interface{}{byId, byName, id}

func byId(name string) {
	fmt.Println(name)
}

func byName(ID int) {
	fmt.Println(ID)
}

func id(name string) {
	fmt.Println(name)
}

func Main(options []string) {
	switch strings.ToLower(options[0]) {
	case "by-id":
		byId("hello")
	}
}
