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

	fmt.Println(document)

	e, _ := token.GetElementByTagName("div", document)

	fmt.Println(token.Render(e))
}
