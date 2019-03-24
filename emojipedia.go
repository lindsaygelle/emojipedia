package emojipedia

import (
	"errors"
	"fmt"
	"strings"

	"github.com/gellel/emojipedia/document"
	"golang.org/x/net/html"

	"github.com/imroc/req"
)

const (
	URL string = "https://www.emojipedia.org"
)

func GetCategories(root *html.Node) (categories []*Category, ok error) {

	heading, ok := document.GetElementByTextContent("Categories", root)

	if ok != nil {
		err := fmt.Sprintf("unable to find reliable anchor point to parse page. does content still have 'categories' title?")
		return nil, errors.New(err)
	}

	ul, ok := document.GetElementByTagName("ul", heading.Parent)

	if ok != nil {
		err := fmt.Sprintf("cannot find collection of categories. are these still held by a UL HTML tag?")
		return nil, errors.New(err)
	}

	anchors := document.GetElementsByTagName("a", ul)

	if len(anchors) == 0 {
		err := fmt.Sprintf("cannot find a collection of categories. are there any A HTML tags held by a parent UL?")
		return nil, errors.New(err)
	}

	for i, a := range anchors {
		for _, attr := range a.Attr {
			if strings.ToLower(attr.Key) == "href" {
				category := &Category{
					Fragment: attr.Val,
					Href:     fmt.Sprintf(URL+"%s", attr.Val),
					Name:     strings.Replace(attr.Val, "/", "", -1),
					Position: i}
				categories = append(categories, category)
				break
			}
		}
	}
	return categories, nil
}

func GetCategoryPage(category *Category) (collection *Collection, ok error) {

	body, ok := GetPage(category.Href)

	if ok != nil {
		fmt.Println(ok)
	}

	heading, ok := document.GetElementByTagName("h1", body)

	if ok != nil {
		err := fmt.Sprintf("unable to find categories title for '%s'. does the page '%s' still have a h1 tag?", category.Name, category.Href)
		return nil, errors.New(err)
	}

	text := document.GetTextNodes(heading.FirstChild)

	for _, t := range text {
		if t.Data != "" {
			fmt.Println(t.Data)
		}
	}

	//fmt.Println(text)

	collection = &Collection{
		Heading: heading.FirstChild.Data}

	fmt.Println(collection.Heading)

	return collection, nil
}

func GetPage(url string) (node *html.Node, ok error) {

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
