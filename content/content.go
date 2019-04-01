package content

import (
	"strconv"
	"strings"

	"github.com/gellel/emojipedia/emoji"

	"github.com/PuerkitoBio/goquery"
)

func Create(document *goquery.Document) map[string]*emoji.Emoji {
	var category, subcategory string
	replacements := []string{".", "", ":", "", ",", "", "⊛", "", "“", "", "”", ""}
	replacer := strings.NewReplacer(replacements...)
	emojis := map[string]*emoji.Emoji{}
	document.Find("tr").Each(func(i int, selection *goquery.Selection) {
		var unicodes string
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
		e := &emoji.Emoji{
			Category:    category,
			Code:        codes,
			Keywords:    keys,
			Name:        name,
			Number:      number,
			Sample:      sample,
			SubCategory: subcategory,
			Unicode:     unicodes}
		emojis[name] = e
	})
	return emojis
}
