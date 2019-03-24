package emojipedia

import (
	"errors"
	"fmt"

	"github.com/gellel/emojipedia/document"
	"golang.org/x/net/html"

	"github.com/imroc/req"
)

func Categories(root *html.Node) (categories []string, ok error) {

	heading, ok := document.GetElementByTextContent("Categories", root)

	if ok != nil {
		err := fmt.Sprintf("unable to find reliable anchor point to parse page. does content still have 'categories' title?")
		return nil, errors.New(err)
	}

	ul, _ := document.GetElementByTagName("ul", heading.Parent)

	anchors := document.GetElementsByTagName("a", ul)

	if len(anchors) == 0 {
		err := fmt.Errorf("cannot find emoji categories")
		fmt.Println(err)
	}

	for i, a := range anchors {
		fmt.Println(a, i)
	}
	return categories, nil
}

func Get(url string) (node *html.Node, ok error) {

	header := req.Header{
		"Accept": "application/json"}

	response, ok := req.Get(url, header)

	if ok != nil {
		r := response.Response()
		err := fmt.Sprintf("unable to fetch content from %s. status: %s. status code: %v", url, r.Status, r.StatusCode)
		return nil, errors.New(err)
	}

	html, ok := document.Parse(response)

	if ok != nil {
		err := fmt.Sprintf("unable to parse content returned from emojipedia.org. possible malformed or incomplete")
		return nil, errors.New(err)
	}

	body, ok := document.GetElementByTagName("body", html)

	if ok != nil {
		err := fmt.Sprintf("cannot find body node in document content. possibly missing or mislabelled")
		return nil, errors.New(err)
	}

	return body, ok
}
