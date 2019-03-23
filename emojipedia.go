package emojipedia

import (
	"fmt"

	"github.com/gellel/emojipedia/document"
	"github.com/gellel/emojipedia/element"
	"github.com/imroc/req"
)

func Get() {

	header := req.Header{
		"Accept": "application/json"}

	response, _ := req.Get("https://emojipedia.org", header)

	root, _ := document.Parse(response)

	body, _ := document.GetBody(root)

	t, _ := document.GetElementByTextContent("Categories", body)

	fmt.Println(document.Render(t.Parent))

	node := element.Node{body}

	fmt.Println(node.GetID())
}
