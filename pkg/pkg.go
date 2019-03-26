package pkg

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gellel/emojipedia/eji"
)

const url string = "https://www.unicode.org/emoji/charts/emoji-list.html"

var replacer *strings.Replacer = strings.NewReplacer(":", "", ",", "", "⊛", "", "“", "", "”", "")

var categories map[string][]*eji.Emoji = map[string][]*eji.Emoji{}

var subcategories map[string][]*eji.Emoji = map[string][]*eji.Emoji{}

func collect(document *goquery.Document) (content map[string]interface{}, err error) {
	var category, subcategory string
	document.Find("tr").Each(func(i int, selection *goquery.Selection) {
		var unicodes string
		var columns []string
		selection.Find("th.bighead").Each(func(j int, s *goquery.Selection) {
			category = s.Text()
			categories[category] = []*eji.Emoji{}
		})
		selection.Find("th.mediumhead").Each(func(j int, s *goquery.Selection) {
			subcategory = s.Text()
			subcategories[subcategory] = []*eji.Emoji{}
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
		if s, ok := categories[category]; ok {
			categories[category] = append(s, emoji)
		}
		if s, ok := subcategories[subcategory]; ok {
			subcategories[subcategory] = append(s, emoji)
		}
		fmt.Println(emoji.Number, emoji.Name, emoji.Sample)
	})
	return content, err
}

func fetch() (content map[string]interface{}, err error) {
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

func Get() (map[string]interface{}, error) {
	return fetch()
}
