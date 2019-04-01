package web

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gellel/emojipedia/cache"
)

func GetHTML() (*goquery.Document, error) {
	name := "emoji-list.html"
	exists := cache.Exists(name)
	if exists {
		return cache.GetHTML(name)
	}
	url := strings.Join([]string{"https://unicode.org/emoji/charts", name}, "/")
	response, err := http.Get(url)
	if err != nil {
		return nil, err
	}
	if response.StatusCode != 200 {
		return nil, fmt.Errorf("%s %s", url, response.Status)
	}
	err = cache.WriteHTML(name, response)
	defer response.Body.Close()
	return cache.GetHTML(name)
}
