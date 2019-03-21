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

	fmt.Println(token.GetElementsByTagName("div", body))

	fmt.Println(len(token.GetElementsByClassName("container", body)))
}
