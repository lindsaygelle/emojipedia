package token

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/imroc/req"

	"golang.org/x/net/html"
)

func Render(n *html.Node) string {
	var buf bytes.Buffer
	w := io.Writer(&buf)
	html.Render(w, n)
	return buf.String()
}

func GetBody(document *html.Node) (body *html.Node, ok error) {
	var f func(*html.Node)
	var n *html.Node
	f = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == "body" {
			n = node
		}
		for c := node.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(document)
	if n != nil {
		return n, nil
	}
	return nil, errors.New("Missing <body> in the node tree")
}

func GetElementsByClassName(class string, document *html.Node) []*html.Node {
	elements := []*html.Node{}
	var f func(*html.Node)
	f = func(node *html.Node) {
		if node != nil {
			if node.Type == html.ElementNode {
				c := true
				for _, a := range node.Attr {
					if a.Key == "class" {
						classes := strings.Split(a.Val, " ")
						for _, cls := range classes {
							if cls == class {
								elements = append(elements, node)
								break
							}
						}
					}
					if c == false {
						break
					}
				}
			}
			f(node.FirstChild)
			f(node.NextSibling)
		}
	}
	f(document)
	return elements
}

func GetElementById(id string, document *html.Node) (element *html.Node, ok bool) {
	for _, a := range document.Attr {
		if a.Key == "id" && a.Val == id {
			return document, true
		}
	}
	for c := document.FirstChild; c != nil; c = c.NextSibling {
		if element, ok = GetElementById(id, c); ok {
			return
		}
	}
	return
}

func GetElementByTextContent(s string, document *html.Node) (element *html.Node, ok error) {
	var f func(*html.Node)
	var n *html.Node
	f = func(node *html.Node) {
		if node != nil {
			if node.Type == html.TextNode {
				fmt.Println(node.Data)
			}
			if node.Type == html.TextNode && s == node.Data {
				n = node.Parent
			}
			if n == nil {
				f(node.FirstChild)
				f(node.NextSibling)
			}
		}
	}
	f(document)
	if n != nil {
		return n, nil
	}
	return nil, errors.New("Content not found")
}

func GetElementsByTagName(tag string, document *html.Node) []*html.Node {
	elements := []*html.Node{}
	var f func(*html.Node)
	f = func(node *html.Node) {
		if node != nil {
			if node.Type == html.ElementNode && node.Data == tag {
				elements = append(elements, node)
			}
			f(node.FirstChild)
			f(node.NextSibling)
		}
	}
	f(document)
	return elements
}

func GetTextNodes(document *html.Node) []*html.Node {
	text := []*html.Node{}
	var f func(*html.Node)
	f = func(node *html.Node) {
		if node != nil {
			if node.Type == html.TextNode {
				text = append(text, node)
			}
			f(node.FirstChild)
			f(node.NextSibling)
		}
	}
	f(document)
	return text
}

func Parse(response *req.Resp) (element *html.Node, ok error) {
	return html.Parse(strings.NewReader(response.String()))
}
