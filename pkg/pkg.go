package pkg

import (
	"errors"
	"net/http"
	"strconv"
	"strings"

	"github.com/gellel/emojipedia/store"

	"github.com/PuerkitoBio/goquery"
	"github.com/gellel/emojipedia/eji"
)

const (
	url               string = "https://www.unicode.org/emoji/charts/emoji-list.html"
	categoriesName    string = "emoji-categories"
	subcategoriesName string = "emoji-subcategories"
	listName          string = "emoji-list"
	wordsName         string = "emoji-keywords"
)

var (
	replacements  []string          = []string{".", "", ":", "", ",", "", "⊛", "", "“", "", "”", ""}
	replacer      *strings.Replacer = strings.NewReplacer(replacements...)
	categories    *eji.Set          = &eji.Set{}
	subcategories *eji.Set          = &eji.Set{}
	emojis        *eji.Map          = &eji.Map{}
	words         *eji.Set          = &eji.Set{}
	files         []string          = []string{categoriesName, subcategoriesName, listName, wordsName}
)

func collect(document *goquery.Document) (*eji.Pkg, error) {
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
		emoji := &eji.Emoji{
			Category:    category,
			Code:        codes,
			Keywords:    keys,
			Name:        name,
			Number:      number,
			Sample:      sample,
			SubCategory: subcategory,
			Unicode:     unicodes}
		for _, v := range emoji.Keywords {
			words.Add(v, emoji.Name)
		}
		categories.Add(category, emoji.Name)
		subcategories.Add(subcategory, emoji.Name)
		emojis.Add(emoji.Name, emoji)
	})
	if len(*categories) != 0 && len(*subcategories) != 0 {
		pkg := &eji.Pkg{
			Categories:    categories,
			Subcategories: subcategories,
			Keywords:      words,
			Names:         emojis}
		return pkg, nil
	}
	return nil, errors.New("unable to parse content")
}

func fetch() (*eji.Pkg, error) {
	res, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		return nil, err
	}
	document, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		return nil, err
	}
	return collect(document)
}

func Get() error {
	err := store.Has(files...)
	if err == nil {
		return nil
	}
	pkg, err := fetch()
	if err != nil {
		return err
	}
	if _, err := store.Save(categoriesName, pkg.Categories); err != nil {
		return err
	}
	if _, err := store.Save(subcategoriesName, pkg.Subcategories); err != nil {
		return err
	}
	if _, err := store.Save(wordsName, pkg.Keywords); err != nil {
		return err
	}
	if _, err := store.Save(listName, pkg.Names); err != nil {
		return err
	}
	return nil
}
