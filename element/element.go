package element

import (
	"strings"

	"golang.org/x/net/html"
)

type Node struct {
	Node *html.Node
}

func (node *Node) GetElementByID(ID string) *Node {
	var e *Node
	var f func(*html.Node)

	f = func(n *html.Node) {
		for _, attr := range n.Attr {
			if strings.ToUpper(attr.Key) == "ID" {
				if attr.Val == ID {
					e = &Node{n}
				}
				break
			}
		}
		if e == nil {
			if n.FirstChild != nil {
				f(n.FirstChild)
			}
			if n.NextSibling != nil {
				f(n.NextSibling)
			}
		}
	}
	f(node.Node)
	return e
}

func (node *Node) GetElementsByClassName(CLS string) []*Node {
	var e []*Node
	var f func(*html.Node)

	f = func(n *html.Node) {
		for _, attr := range n.Attr {
			if strings.ToUpper(attr.Key) == "CLASS" {
				classes := strings.Split(attr.Val, " ")
				for _, cls := range classes {
					if strings.ToUpper(cls) == CLS {
						e = append(e, &Node{n})
						break
					}
				}
				break
			}
		}
		if n.FirstChild != nil {
			f(n.FirstChild)
		}
		if n.NextSibling != nil {
			f(n.NextSibling)
		}
	}
	f(node.Node)
	return e
}

func (node *Node) GetElementsByTagName(TAG string) []*Node {
	var e []*Node
	var f func(*html.Node)
	f = func(n *html.Node) {
		if n.Type == html.ElementNode {
			if strings.ToUpper(n.Data) == strings.ToUpper(TAG) {
				e = append(e, &Node{n})
			}
		}
		if n.FirstChild != nil {
			f(n.FirstChild)
		}
		if n.NextSibling != nil {
			f(n.NextSibling)
		}
	}
	f(node.Node)
	return e
}

func (node *Node) GetClassList() []string {
	var CLS []string
	for _, attr := range node.Node.Attr {
		if strings.ToUpper(attr.Key) == "CLASS" {
			CLS = strings.Split(attr.Val, " ")
			break
		}
	}
	return CLS
}

func (node *Node) GetID() string {
	var ID string
	for _, attr := range node.Node.Attr {
		if strings.ToUpper(attr.Key) == "ID" {
			ID = attr.Val
			break
		}
	}
	return ID
}

func (node *Node) GetParent() *Node {
	return &Node{node.Node.Parent}
}

func (node *Node) GetPreviouSibling() *Node {
	return &Node{node.Node.PrevSibling}
}