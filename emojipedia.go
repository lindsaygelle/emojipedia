package emojipedia

import (
	"fmt"

	"github.com/gellel/emojipedia/token"
	"github.com/imroc/req"
)

func Get() {

	header := req.Header{
		"Accept": "application/json"}

	response, _ := req.Get("https://emojipedia.org", header)

	document, _ := token.Parse(response)

	body, _ := token.GetBody(document)

	t, _ := token.GetElementByTextContent("Categories", body)

	fmt.Println(t)
}
