package emojipedia

import (
	"regexp"
	"strconv"
	"strings"
	"unicode"

	"github.com/PuerkitoBio/goquery"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

type Emoji struct {
	Category    string   `json:"Category"`
	Code        string   `json:"Code"`
	Description string   `json:"Description"`
	Keywords    []string `json:"Keywords"`
	Name        string   `json:"Name"`
	Number      int      `json:"Number"`
	Sample      string   `json:"Sample"`
	Subcategory string   `json:"Subcategory"`
	Unicode     string   `json:"Unicode"`
}

var replacements = []string{" ", "-", "&", "and"}

var replacer = strings.NewReplacer(replacements...)

var reg, _ = regexp.Compile(`[^a-zA-Z0-9\-\?\!:]+`)

func NewCategories(doc *goquery.Document) map[int]string {
	categories := []string{}
	doc.Find("tr").Each(func(_ int, selection *goquery.Selection) {
		selection.Find("th.bighead").Each(func(_ int, selection *goquery.Selection) {
			categories = append(categories, strings.TrimSpace(selection.Text()))
		})
	})
	m := map[int]string{}
	for i, c := range categories {
		m[i] = Normalize(c)
	}
	return m
}

func NewKeywords(doc *goquery.Document) map[string][]string {
	m := map[string][]string{}
	doc.Find("tr").Each(func(_ int, selection *goquery.Selection) {
		fields := []string{}
		selection.Find("td").Each(func(i int, selection *goquery.Selection) {
			fields = append(fields, strings.TrimSpace(selection.Text()))

		})
		if len(fields) != 5 {
			return
		}
		fields = fields[3:]
		name := Normalize(fields[0])
		for _, key := range strings.Split(fields[1], "|") {
			key = Normalize(key)
			if _, ok := m[key]; ok != true {
				m[key] = []string{}
			}
			m[key] = append(m[key], name)
		}
	})
	return m
}

func NewSubcategories(doc *goquery.Document) map[int]string {
	subcategories := []string{}
	doc.Find("tr").Each(func(_ int, selection *goquery.Selection) {
		selection.Find("th.mediumhead").Each(func(_ int, selection *goquery.Selection) {
			subcategories = append(subcategories, strings.TrimSpace(selection.Text()))
		})
	})
	m := map[int]string{}
	for i, c := range subcategories {
		m[i] = Normalize(c)
	}
	return m
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
	name = Normalize(name)
	keywords, columns := columns[0], columns[1:]
	keywords = strings.Replace(keywords, "|", "", -1)
	keys := strings.Fields(keywords)
	category, columns := Normalize(columns[0]), columns[1:]
	subcategory, columns := Normalize(columns[0]), columns[1:]
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

func NewEmojis(doc *goquery.Document) map[string]*Emoji {
	e := map[string]*Emoji{}
	var category, subcategory string
	doc.Find("tr").Each(func(i int, selection *goquery.Selection) {
		var columns []string
		selection.Find("th.bighead").Each(func(j int, s *goquery.Selection) {
			category = Normalize(strings.TrimSpace(s.Text()))
		})
		selection.Find("th.mediumhead").Each(func(j int, s *goquery.Selection) {
			subcategory = Normalize(strings.TrimSpace(s.Text()))
		})
		selection.Find("td").Each(func(j int, s *goquery.Selection) {
			columns = append(columns, strings.TrimSpace(s.Text()))
		})
		if len(columns) != 5 {
			return
		}
		emoji := NewEmoji(append(columns, category, subcategory))
		e[emoji.Name] = emoji
	})
	return e
}

func Normalize(value string) string {
	f := func(r rune) bool {
		return unicode.Is(unicode.Mn, r)
	}
	t := transform.Chain(norm.NFD, transform.RemoveFunc(f), norm.NFC)
	result, _, _ := transform.String(t, value)
	result = replacer.Replace(strings.TrimSpace(result))
	result = reg.ReplaceAllString(strings.ToLower(result), "")
	if strings.HasPrefix(result, "-") {
		result = strings.TrimPrefix(result, "-")
	}
	if strings.HasSuffix(result, "-") {
		result = strings.TrimSuffix(result, "-")
	}
	return result
}
