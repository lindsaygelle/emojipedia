package web

import (
	"net/http"

	"github.com/PuerkitoBio/goquery"
)

func Http() *goquery.Document {
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
