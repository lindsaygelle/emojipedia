package emojipedia

import (
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

var _ codex = (*Codex)(nil)

var _ catalogue = (*Catalogue)(nil)

var _ names = (*Names)(nil)

var replacements = []string{
	".", "",
	":", "",
	",", "",
	"⊛", "",
	"“", "",
	"”", ""}

var replacer = strings.NewReplacer(replacements...)

type catalogue interface {
	Add(key string, name string) string
	Get(key string) *Names
	Has(key string) bool
}

type codex interface {
	Add(key int, name string) string
	Get(key int) string
	Has(key int) bool
}

type names interface {
	Add(name string) string
	Get(i int) string
	Has(name string) bool
}

type Catalogue map[string]*Names

type Codex map[int]string

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
	Categories    *Catalogue
	Subcategories *Catalogue
	Keywords      *Catalogue
	Numeric       *Codex
}

type Emojipedia struct {
	Emojidex     *Emojidex
	Encyclopedia *Encyclopedia
}

type Names []string

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
			category = encyclopedia.Categories.Set(strings.TrimSpace(s.Text()))
		})
		selection.Find("th.mediumhead").Each(func(j int, s *goquery.Selection) {
			subcategory = encyclopedia.Subcategories.Set(strings.TrimSpace(s.Text()))
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
		Categories:    &Catalogue{},
		Subcategories: &Catalogue{},
		Keywords:      &Catalogue{},
		Numeric:       &Codex{}}
}

func (catalogue *Catalogue) Add(key string, name string) string {
	if catalogue.Has(key) != true {
		catalogue.Set(key)
	}
	return catalogue.Get(key).Add(name)
}

func (catalogue *Catalogue) Has(key string) bool {
	_, ok := (*catalogue)[key]
	return ok
}

func (catalogue *Catalogue) Get(key string) *Names {
	names, _ := (*catalogue)[key]
	return names
}

func (catalogue *Catalogue) Set(key string) string {
	(*catalogue)[key] = &Names{}
	return key
}

func (codex *Codex) Add(key int, name string) string {
	(*codex)[key] = name
	return name
}

func (codex *Codex) Has(key int) bool {
	_, ok := (*codex)[key]
	return ok
}

func (codex *Codex) Get(key int) string {
	name, _ := (*codex)[key]
	return name
}

func (names *Names) Add(name string) string {
	*names = append(*names, name)
	return name
}

func (names *Names) Has(name string) bool {
	n := *names
	i := 0
	j := len(n) - 1
	if j < 0 {
		return false
	}
	for i <= j {
		if n[i] == name || n[j] == name {
			return true
		}
		j = j - 1
		i = i + 1
	}
	return false
}

func (names *Names) Get(i int) string {
	var s string
	n := *names
	if len(n) > 0 {
		s = n[i]
	}
	return s
}
