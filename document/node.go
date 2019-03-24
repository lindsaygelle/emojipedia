package document

import (
	"strings"

	"golang.org/x/net/html"
)

type Node struct {
	Element *html.Node
}

func (node *Node) GetChildren() (n []*Node) {
	f = func(child *html.Node) {
		if child.Type == html.ElementNode {
			n = append(n, &Node{child})
		}
		if child.FirstChild != nil {
			f(child.FirstChild)
		}
		if child.NextSibling != nil {
			f(child.NextSibling)
		}
	}
	f(node.Element)
	return n
}

func (node *Node) GetClasses() (classes map[string]int) {
	for _, attr := range node.Element.Attr {
		if strings.ToLower(attr.Key) == "class" {
			for i, cls := range strings.Split(attr.Val, " ") {
				classes[cls] = i
			}
			break
		}
	}
	return classes
}

func (node *Node) GetElementByClassName(class string) (n *Node) {
	f = func(child *html.Node) {
		if child.Type == html.ElementNode {
			for _, attr := range child.Attr {
				if strings.ToLower(attr.Key) == "class" {
					for _, cls := range strings.Split(attr.Val, " ") {
						if strings.ToLower(cls) == strings.ToLower(class) {
							n = &Node{child}
						}
					}
				}
			}
		}
		if n != nil && child.FirstChild != nil {
			f(child.FirstChild)
		}
		if n != nil && child.NextSibling != nil {
			f(child.NextSibling)
		}
	}
	f(node.Element)
	return n
}

func (node *Node) GetElementByTextContent(text string) (n *Node) {
	f = func(child *html.Node) {
		if child.Type == html.TextNode {
			if child.Data == text {
				n = &Node{child.Parent}
			}
		}
		if n == nil && child.FirstChild != nil {
			f(child.FirstChild)
		}
		if n == nil && child.NextSibling != nil {
			f(child.NextSibling)
		}
	}
	f(node.Element)
	return n
}

func (node *Node) GetElementByTagName(tag string) (n *Node) {
	f = func(child *html.Node) {
		if child.Type == html.ElementNode {
			if strings.ToLower(child.Data) == strings.ToLower(tag) {
				n = &Node{child}
			}
		}
		if n == nil && child.FirstChild != nil {
			f(child.FirstChild)
		}
		if n == nil && child.NextSibling != nil {
			f(child.NextSibling)
		}
	}
	f(node.Element)
	return n
}

func (node *Node) GetElementsByClassName(class string) (n []*Node) {
	f = func(child *html.Node) {
		if child.Type == html.ElementNode {
			for _, attr := range child.Attr {
				if strings.ToLower(attr.Key) == "class" {
					for _, cls := range strings.Split(attr.Val, " ") {
						if strings.ToLower(cls) == strings.ToLower(class) {
							n = append(n, &Node{child})
							break
						}
					}
					break
				}
			}
		}
		if child.FirstChild != nil {
			f(child.FirstChild)
		}
		if child.NextSibling != nil {
			f(child.NextSibling)
		}
	}
	f(node.Element)
	return n
}

func (node *Node) GetElementsByTagName(tag string) (n []*Node) {
	f = func(child *html.Node) {
		if child.Type == html.ElementNode {
			if strings.ToLower(tag) == strings.ToLower(tag) {
				n = append(n, &Node{child})
			}
		}
		if child.FirstChild != nil {
			f(child.FirstChild)
		}
		if child.NextSibling != nil {
			f(child.NextSibling)
		}
	}
	f(node.Element)
	return n
}

func (node *Node) GetID(id string) (n *Node) {
	f = func(child *html.Node) {
		for _, attr := range node.Element.Attr {
			if strings.ToLower(attr.Key) == "id" {
				if attr.Val == id {
					n = &Node{child}
				}
			}
			break
		}
		if n == nil && child.FirstChild != nil {
			f(child.FirstChild)
		}
		if n == nil && child.NextSibling != nil {
			f(child.NextSibling)
		}
	}
	f(node.Element)
	return n
}

func (node *Node) GetTextNode() (n []*html.Node) {
	f = func(child *html.Node) {
		if child.Type == html.TextNode {
			n = append(n, child)
		}
		if child.FirstChild != nil {
			f(child.FirstChild)
		}
		if child.NextSibling != nil {
			f(child.NextSibling)
		}
	}
	f(node.Element)
	return n
}
