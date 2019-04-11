package emojipedia

import (
	"strconv"
	"strings"
	"unicode"

	"github.com/PuerkitoBio/goquery"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

var replacements = []string{
	".", "",
	":", "",
	",", "",
	"⊛", "",
	"“", "",
	"”", "",
	"_", "-",
	"&", "and",
	" ", "-"}

var replacer = strings.NewReplacer(replacements...)

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
	name = strings.Replace(strings.TrimSpace(replacer.Replace(name)), " ", "-", -1)
	name = strings.Replace(strings.ToLower(name), "_", "-", -1)
	name = Normalize(name)
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
		Subcategory: subcategory,
		Unicode:     unicodes}
}

func NewEmojidex() *Emojidex {
	return &Emojidex{}
}

func NewEmojipedia() *Emojipedia {
	return &Emojipedia{
		Emojidex:     NewEmojidex(),
		Encyclopedia: NewEncyclopedia()}
}

func NewEmojipediaFromDocument(document *goquery.Document) *Emojipedia {
	var category, subcategory string
	emojidex := NewEmojidex()
	encyclopedia := NewEncyclopedia()
	document.Find("tr").Each(func(i int, selection *goquery.Selection) {
		var columns []string
		selection.Find("th.bighead").Each(func(j int, s *goquery.Selection) {
			category = replacer.Replace(encyclopedia.Categories.Set(strings.TrimSpace(s.Text())))
		})
		selection.Find("th.mediumhead").Each(func(j int, s *goquery.Selection) {
			subcategory = replacer.Replace(encyclopedia.Subcategories.Set(strings.TrimSpace(s.Text())))
		})
		selection.Find("td").Each(func(j int, s *goquery.Selection) {
			columns = append(columns, strings.TrimSpace(s.Text()))
		})
		if len(columns) != 5 {
			return
		}
		columns = append(columns, category, subcategory)
		emoji := NewEmoji(columns)
		for _, keyword := range emoji.Keywords {
			encyclopedia.Keywords.Add(keyword, emoji.Name)
		}
		encyclopedia.Categories.Add(category, emoji.Name)
		encyclopedia.Subcategories.Add(subcategory, emoji.Name)
		encyclopedia.Numeric.Add(emoji.Number, emoji.Name)
	})
	return &Emojipedia{
		Emojidex:     emojidex,
		Encyclopedia: encyclopedia}
}

func NewEncyclopedia() *Encyclopedia {
	return &Encyclopedia{
		Categories:    &Associative{},
		Subcategories: &Associative{},
		Keywords:      &Associative{},
		Numeric:       &Set{}}
}

func Normalize(value string) string {
	f := func(r rune) bool {
		return unicode.Is(unicode.Mn, r)
	}
	t := transform.Chain(norm.NFD, transform.RemoveFunc(f), norm.NFC)
	result, _, _ := transform.String(t, value)
	return result
}
