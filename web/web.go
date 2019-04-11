package web

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

const UnicodeOrgURL string = "https://www.unicode.org/emoji/charts/emoji-list.html"

func Http(url string) *goquery.Document {
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
