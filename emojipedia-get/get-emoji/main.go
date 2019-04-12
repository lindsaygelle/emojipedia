package emoji

import (
	"fmt"
	"strings"

	"github.com/gellel/emojipedia/emojipedia"
	all "github.com/gellel/emojipedia/emojipedia-get/get-emoji/emoji-all"
	"github.com/gellel/emojipedia/manifest"
)

var Export = Emoji

var Key = "EMOJI"

var Options = []interface{}{Category, Description, Keywords, Subcategory}

var emojidex *emojipedia.Emojidex

var name = map[string]func(name string){
	"CATEGORY":    Category,
	"CODES":       Codes,
	"DESCRIPTION": Description,
	"KEYWORDS":    Keywords,
	"NUMBER":      Number,
	"SUBCATEGORY": Subcategory}

var set = map[string](func(m *manifest.Manifest, previous, options []string)){
	all.Key: all.Main}

func Emoji(options ...string) {}

func Category(name string) {
	if e, ok := emojidex.Get(name); ok {
		fmt.Println(fmt.Sprintf("%s category: %s.", name, e.Category))
	}
}

func Codes(name string) {
	if e, ok := emojidex.Get(name); ok {
		fmt.Println(fmt.Sprintf("%s codes: %s", name, e.Code))
	}
}

func Description(name string) {
	if _, ok := emojidex.Get(name); ok {
		fmt.Println(fmt.Sprintf("%s description: %s.", name, "MISSING"))
	}
}

func Keywords(name string) {
	if e, ok := emojidex.Get(name); ok {
		fmt.Println(fmt.Sprintf("%s keywords: %s.", name, strings.Join(e.Keywords, ", ")))
	}
}

func Number(name string) {
	if e, ok := emojidex.Get(name); ok {
		fmt.Println(fmt.Sprintf("%s number: %v.", name, e.Number))
	}
}

func Subcategory(name string) {
	if e, ok := emojidex.Get(name); ok {
		fmt.Println(fmt.Sprintf("%s category: %s.", name, e.Subcategory))
	}
}

func Main(m *manifest.Manifest, previous []string, options []string) {
	if len(options) != 0 {
		key := strings.ToUpper(strings.Replace(options[0], "-", "", -1))
		if f, ok := name[key]; ok {
			emojidex = emojipedia.UnmarshallEmojidex()
			f(options[1])
		} else if f, ok := set[key]; ok {
			f(m, append(previous, key), options[1:])
		}
	}
}
