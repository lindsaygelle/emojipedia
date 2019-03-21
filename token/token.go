package token

import (
	"bytes"
	"fmt"
	"io"

	"golang.org/x/net/html"
)

func Render(n *html.Node) string {
	var buf bytes.Buffer
	w := io.Writer(&buf)
	html.Render(w, n)
	return buf.String()
}

func Get(HTMLElement string, root *html.Node) *html.Node {
	fmt.Println(Render(root))
	return root
}

func GetAll(HTMLElement string, root *html.Node) []*html.Node {
	return []*html.Node{root}
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
