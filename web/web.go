package web

import (
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gellel/emojipedia/emojipedia"
)

func HttptUnicodeOrg() *goquery.Document {
	url := "https://www.unicode.org/emoji/charts/emoji-list.html"
	res, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer res.Body.Close()
	if res.StatusCode != 200 {
		panic(err)
	}
	document, err := goquery.NewDocumentFromResponse(res)
	if err != nil {
		panic(err)
	}
	return document
}

func HttpToEmojiPackage(document *goquery.Document) (*emojipedia.Emojidex, *emojipedia.Encyclopedia) {
	var category, subcategory string
	encyclopedia := emojipedia.NewEncyclopedia()
	emojidex := emojipedia.NewEmojidex()
	document.Find("tr").Each(func(i int, selection *goquery.Selection) {
		var columns []string
		selection.Find("th.bighead").Each(func(j int, s *goquery.Selection) {
			category = strings.TrimSpace(s.Text())
		})
		selection.Find("th.mediumhead").Each(func(j int, s *goquery.Selection) {
			subcategory = strings.TrimSpace(s.Text())
		})
		selection.Find("td").Each(func(j int, s *goquery.Selection) {
			columns = append(columns, strings.TrimSpace(s.Text()))
		})
		if len(columns) != 5 {
			return
		}
		columns = append(columns, category, subcategory)
		emoji := emojipedia.NewEmoji(columns)
		for _, keyword := range emoji.Keywords {
			if _, ok := encyclopedia.Keywords[keyword]; !ok {
				encyclopedia.Keywords[keyword] = []string{}
			}
			encyclopedia.Keywords[keyword] = append(encyclopedia.Keywords[keyword], emoji.Name)
		}
		if _, ok := encyclopedia.Categories[category]; !ok {
			encyclopedia.Categories[category] = []string{}
		}
		if _, ok := encyclopedia.Subcategories[subcategory]; !ok {
			encyclopedia.Categories[subcategory] = []string{}
		}
		encyclopedia.Categories[category] = append(encyclopedia.Categories[category], emoji.Name)
		encyclopedia.Subcategories[subcategory] = append(encyclopedia.Categories[subcategory], emoji.Name)
		encyclopedia.Numeric[emoji.Number] = emoji.Name

	})
	return emojidex, encyclopedia
}
