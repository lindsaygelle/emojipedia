package all

import (
	"fmt"
	"strconv"
	"strings"

	"github.com/gellel/emojipedia/emojipedia"
	"github.com/gellel/emojipedia/manifest"

	api "github.com/gellel/emojipedia/emojipedia-api"
)

var Export = All

var Key = "ALL"

var Options = []interface{}{Alphabetic, Characters, Numeric}

var emojidex map[string]*emojipedia.Emoji

var empty = map[string](func()){
	"ALPHABETIC": Alphabetic,
	"CHARACTERS": Characters,
	"NUMERIC":    Numeric}

func All(options ...string) {}

func Alphabetic() {
	for i, name := range emojipedia.SortByName(&emojidex) {
		fmt.Println(i, name)
	}
}

func Characters() {
	for _, e := range emojidex {
		r, _ := strconv.ParseInt(strings.TrimPrefix(e.Unicode, "\\U"), 16, 32)
		fmt.Println(e.Name, fmt.Sprintf("%s", string(r)))
	}
}

func Numeric() {
	for i, name := range emojipedia.SortByID(&emojidex) {
		fmt.Println((i + 1), name)
	}
}

func Main(m *manifest.Manifest, previous, options []string) {
	if len(options) != 0 {
		key := strings.ToUpper(strings.Replace(options[0], "-", "", -1))
		if f, ok := empty[key]; ok {
			emojidex, _ = api.GetEmojis()
			f()
		}
	}
}
