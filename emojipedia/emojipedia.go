package emojipedia

import (
	"strconv"
	"strings"
)

var replacements = []string{
	".", "",
	":", "",
	",", "",
	"⊛", "",
	"“", "",
	"”", ""}

var replacer = strings.NewReplacer(replacements...)

type Emoji struct {
	Category    string   `json:"Category"`
	Code        string   `json:"Code"`
	Keywords    []string `json:"Keywords"`
	Name        string   `json:"Name"`
	Number      int      `json:"Number"`
	Sample      string   `json:"Sample"`
	SubCategory string   `json:"SubCategory"`
	Unicode     string   `json:"Unicode"`
}

type Emojidex map[string]*Emoji

type Encyclopedia struct {
	Categories    map[string][]string
	Subcategories map[string][]string
	Keywords      map[string][]string
	Numeric       map[int]string
}

func NewEmoji(columns []string) *Emoji {
	var unicodes string
	no, columns := columns[0], columns[1:]
	number, _ := strconv.Atoi(no)
	sample, columns := columns[0], columns[1:]
	codes, columns := columns[0], columns[1:]
	for _, code := range strings.Fields(codes) {
		replacement := "000"
		if len(code) == 6 {
			replacement = "0000"
		}
		unicodes = unicodes + strings.Replace(code, "+", replacement, 1)
	}
	unicodes = unicodes + strings.Replace(unicodes, "U", "\\U", 1)
	name, columns := columns[0], columns[1:]
	name = strings.Replace(strings.TrimSpace(replacer.Replace(name)), " ", "_", -1)
	name = strings.ToLower(name)
	keywords, columns := columns[0], columns[1:]
	keywords = strings.Replace(keywords, "|", "", -1)
	keys := strings.Fields(keywords)
	category, columns := columns[0], columns[1:]
	subcategory, columns := columns[0], columns[1:]
	return &Emoji{
		Category:    category,
		Code:        codes,
		Keywords:    keys,
		Name:        name,
		Number:      number,
		Sample:      sample,
		SubCategory: subcategory,
		Unicode:     unicodes}
}

func NewEmojidex() *Emojidex {
	return &Emojidex{}
}

func NewEncyclopedia() *Encyclopedia {
	return &Encyclopedia{
		Categories:    map[string][]string{},
		Subcategories: map[string][]string{},
		Keywords:      map[string][]string{},
		Numeric:       map[int]string{}}
}
