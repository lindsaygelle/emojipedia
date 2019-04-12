package emoji

import (
	"fmt"
	"strings"

	"github.com/gellel/emojipedia/emojipedia"
)

var Options = []interface{}{byId, byName, id}

func byId(name string) {
	emojidex := emojipedia.UnmarshallEmojidex()
	emoji, ok := emojidex.Get(name)
	if ok != true {
		fmt.Println(fmt.Sprintf("emoji %s does not exist", name))
	} else {
		emojipedia.PrintEmoji(emoji)
	}
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
		byId(options[1])
	}
}
