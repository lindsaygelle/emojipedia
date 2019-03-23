package document

import (
	"bytes"
	"errors"
	"fmt"
	"io"
	"strings"

	"github.com/imroc/req"
	"golang.org/x/net/html"
)

var (
	f func(*html.Node)
)

// GetBody returns the HTML body.
func GetBody(document *html.Node) (body *html.Node, ok error) {
	return GetElementByTagName("body", document)
}

// GetElementByClassName returns the first element that contains the given class name.
func GetElementByClassName(class string, document *html.Node) (element *html.Node, ok error) {
	var n *html.Node
	f = func(node *html.Node) {
		if node.Type == html.ElementNode {
			for _, attr := range node.Attr {
				if attr.Key == "class" {
					for _, cls := range strings.Split(attr.Val, " ") {
						if cls == class {
							n = node
							break
						}
					}
					break
				}
			}
		}
		if n == nil && node.FirstChild != nil {
			f(node.FirstChild)
		}
		if n == nil && node.NextSibling != nil {
			f(node.NextSibling)
		}
	}
	f(document)
	if n != nil {
		return n, nil
	}
	s := fmt.Sprintf("cannot find element with class '%s' at address (*%p)", class, &document)
	return nil, errors.New(s)
}

// GetElementsByClassName returns a slice of all child elements which have all of the given class names.
func GetElementsByClassName(class string, document *html.Node) []*html.Node {
	n := []*html.Node{}
	f = func(node *html.Node) {
		if node.Type == html.ElementNode {
			for _, attr := range node.Attr {
				if attr.Key == "class" {
					for _, cls := range strings.Split(attr.Val, " ") {
						if cls == class {
							n = append(n, node)
							break
						}
					}
					break
				}
			}
		}
		if n == nil && node.FirstChild != nil {
			f(node.FirstChild)
		}
		if n == nil && node.NextSibling != nil {
			f(node.NextSibling)
		}
	}
	return n
}

// GetElementByID returns an Element struct representing the element whose id property matches the specified string.
func GetElementByID(id string, document *html.Node) (element *html.Node, ok error) {
	var n *html.Node
	f = func(node *html.Node) {
		if node.Type == html.ElementNode {
			for _, attr := range node.Attr {
				if attr.Key == "id" && attr.Val == id {
					n = node
				}
				break
			}
		}
		if n == nil && node.FirstChild != nil {
			f(node.FirstChild)
		}
		if n == nil && node.NextSibling != nil {
			f(node.NextSibling)
		}
	}
	f(document)
	if n != nil {
		return n, nil
	}
	s := fmt.Sprintf("cannot find element with id '%s' at address (*%p)", id, &document)
	return nil, errors.New(s)
}

// GetElementByTagName returns the first element with the given tag name.
func GetElementByTagName(tag string, document *html.Node) (element *html.Node, ok error) {
	var n *html.Node
	f = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == tag {
			n = node
		}
		if n == nil && node.FirstChild != nil {
			f(node.FirstChild)
		}
		if n == nil && node.NextSibling != nil {
			f(node.NextSibling)
		}
	}
	f(document)
	if n != nil {
		return n, nil
	}
	s := fmt.Sprintf("cannot find element '%s' at address (*%p)", tag, &document)
	return nil, errors.New(s)
}

// GetElementsByTagName returns a slice of elements with the given tag name.
func GetElementsByTagName(tag string, document *html.Node) []*html.Node {
	n := []*html.Node{}
	f = func(node *html.Node) {
		if node.Type == html.ElementNode && node.Data == tag {
			n = append(n, node)
		}
		if n == nil && node.FirstChild != nil {
			f(node.FirstChild)
		}
		if n == nil && node.NextSibling != nil {
			f(node.NextSibling)
		}
	}
	f(document)
	return n
}

// GetElementByTextContent returns a text node element with the matching string.
func GetElementByTextContent(text string, document *html.Node) (element *html.Node, ok error) {
	var n *html.Node
	f = func(node *html.Node) {
		if node.Type == html.TextNode && node.Data == text {
			n = node
		}
		if n == nil && node.FirstChild != nil {
			f(node.FirstChild)
		}
		if n == nil && node.NextSibling != nil {
			f(node.NextSibling)
		}
	}
	f(document)
	if n != nil {
		return n, nil
	}
	s := fmt.Sprintf("cannot find element with text '%s' at address (*%p)", text, &document)
	return nil, errors.New(s)
}

// GetTextNodes returns a slice of text nodes.
func GetTextNodes(document *html.Node) []*html.Node {
	n := []*html.Node{}
	f = func(node *html.Node) {
		if node.Type == html.TextNode {
			n = append(n, node)
		}
		if node.FirstChild != nil {
			f(node.FirstChild)
		}
		if node.NextSibling != nil {
			f(node.NextSibling)
		}
	}
	f(document)
	return n
}

// Parse parses HTML GET response document to HTML.
func Parse(response *req.Resp) (element *html.Node, ok error) {
	return html.Parse(strings.NewReader(response.String()))
}

// Render generates a plain-text string from the argument html struct.
func Render(document *html.Node) string {
	var buffer bytes.Buffer
	writer := io.Writer(&buffer)
	html.Render(writer, document)
	return buffer.String()
}
