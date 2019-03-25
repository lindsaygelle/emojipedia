package pkg

import (
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
)

type Emoji struct {
	Category    string
	Code        string
	Keywords    []string
	Name        string
	Number      int
	Sample      string
	SubCategory string
	Unicode     string
}

const url string = "https://www.unicode.org/emoji/charts/emoji-list.html"

var replacer *strings.Replacer = strings.NewReplacer(":", "", ",", "", "⊛", "", "“", "", "”", "")

func collect(document *goquery.Document) (content map[string]string, err error) {
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
		name = strings.Replace(strings.TrimSpace(name), " ", "_", -1)
		keywords, columns := columns[0], columns[1:]
		keywords = strings.Replace(keywords, "|", "", -1)
		keys := strings.Fields(keywords)
		if len(columns) != 0 {
			return
		}
		emoji := &Emoji{
			Category:    category,
			Code:        codes,
			Keywords:    keys,
			Name:        name,
			Number:      number,
			Sample:      sample,
			SubCategory: subcategory,
			Unicode:     unicodes}

		fmt.Println("name:", emoji.Name, "unicode:", emoji.Unicode)
		fmt.Println("keywords:", emoji.Keywords)
		fmt.Println()
	})
	return content, err
}

func fetch() (content map[string]string, err error) {
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

func Get() (map[string]string, error) {
	return fetch()
}
