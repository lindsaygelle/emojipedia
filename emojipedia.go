package emojipedia

import (
	"strings"

	"github.com/gellel/emojipedia/token"
	"github.com/imroc/req"
	"golang.org/x/net/html"
)

func Get() {

	header := req.Header{
		"Accept": "application/json"}

	response, _ := req.Get("https://emojipedia.org", header)

	tokens, _ := html.Parse(strings.NewReader(response.String()))

	token.Get("html", tokens.FirstChild.NextSibling.FirstChild)

}
