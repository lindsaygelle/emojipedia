package pkg

import (
	"encoding/json"
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gellel/emojipedia/emojipedia"
)

const (
	url               string = "https://www.unicode.org/emoji/charts/emoji-list.html"
	categoriesName    string = "emoji-categories"
	subcategoriesName string = "emoji-subcategories"
	listName          string = "emoji-list"
	wordsName         string = "emoji-keywords"
)

var files []string = []string{
	categoriesName,
	subcategoriesName,
	listName,
	wordsName}

var replacements []string = []string{
	".", "",
	":", "",
	",", "",
	"⊛", "",
	"“", "",
	"”", ""}

var replacer *strings.Replacer = strings.NewReplacer(replacements...)

var pkg *emojipedia.Package = &emojipedia.Package{
	Categories:    &emojipedia.Set{},
	Subcategories: &emojipedia.Set{},
	Keywords:      &emojipedia.Set{},
	Names:         &emojipedia.Map{}}

func collect(document *goquery.Document) error {
	var category, subcategory string
	document.Find("tr").Each(func(i int, selection *goquery.Selection) {
		var unicodes string
		var columns []string
		selection.Find("th.bighead").Each(func(j int, s *goquery.Selection) {
			category = s.Text()
		})
		selection.Find("th.mediumhead").Each(func(j int, s *goquery.Selection) {
			subcategory = s.Text()
		})
		selection.Find("td").Each(func(j int, s *goquery.Selection) {
			columns = append(columns, s.Text())
		})
		if len(columns) != 5 {
			return
		}
		no, columns := columns[0], columns[1:]
		number, err := strconv.Atoi(no)
		if err != nil {
			return
		}
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
		if len(columns) != 0 {
			return
		}
		emoji := &emojipedia.Emoji{
			Category:    category,
			Code:        codes,
			Keywords:    keys,
			Name:        name,
			Number:      number,
			Sample:      sample,
			SubCategory: subcategory,
			Unicode:     unicodes}
		for _, v := range emoji.Keywords {
			pkg.Keywords.Add(v, emoji.Name)
		}
		pkg.Categories.Add(category, emoji.Name)
		pkg.Subcategories.Add(subcategory, emoji.Name)
		pkg.Names.Add(emoji.Name, emoji)
	})
	if len(*pkg.Categories) == 0 || len(*pkg.Subcategories) == 0 || len(*pkg.Names) == 0 {
		return errors.New("unable to parse content")
	}
	return nil
}

func fetch() error {
	res, err := http.Get(url)
	if err != nil {
		return err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return err
	}
	document, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		return err
	}
	return collect(document)
}

func open(m *map[string]interface{}) error {
	for key, value := range *m {
		bytes, err := store.Open(key)
		if err != nil {
			return err
		}
		if err = json.Unmarshal(bytes, value); err != nil {
			return err
		}
	}
	return nil
}

func save(m *map[string]interface{}) error {
	for key, value := range *m {
		if _, err := store.Save(key, value); err != nil {
			return err
		}
	}
	return nil
}

// Get parses content from unicode.org and converts the HTML data into a set of json data.
func Get() (*emojipedia.Package, error) {

	m := make(map[string]interface{})

	m[categoriesName] = pkg.Categories
	m[subcategoriesName] = pkg.Subcategories
	m[wordsName] = pkg.Keywords
	m[listName] = pkg.Names

	if err := store.Has(files...); err == nil {
		if err := open(&m); err == nil {
			return pkg, nil
		}
	}
	if err := fetch(); err != nil {
		return nil, err
	}
	if err := save(&m); err != nil {
		return nil, err
	}
	return pkg, nil
}
