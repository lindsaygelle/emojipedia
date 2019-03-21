package token

import (
	"bytes"
	"errors"
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

func GetAll(HTMLElement string, document *html.Node) []*html.Node {
	return []*html.Node{document}
}

func GetBody(document *html.Node) (body *html.Node, ok error) {
	var b *html.Node
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode && n.Data == "body" {
			b = n
		}
		for c := n.FirstChild; c != nil; c = c.NextSibling {
			f(c)
		}
	}
	f(document)
	if b != nil {
		return b, nil
	}
	return nil, errors.New("Missing <body> in the node tree")
}

func GetElementsByClassName(class string, root *html.Node) (elements []*html.Node, ok bool) {
	return []*html.Node{}, false
}

func GetElementById(id string, root *html.Node) (element *html.Node, ok bool) {
	for _, a := range root.Attr {
		if a.Key == "id" && a.Val == id {
			return root, true
		}
	}
	for c := root.FirstChild; c != nil; c = c.NextSibling {
		if element, ok = GetElementById(id, c); ok {
			return
		}
	}
	return
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

func Parse(response *req.Resp) (element *html.Node, ok error) {
	return html.Parse(strings.NewReader(response.String()))
}
