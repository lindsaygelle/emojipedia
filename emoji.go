package main

import (
	"fmt"
	"net/http"
	"regexp"
	"strings"

	"github.com/PuerkitoBio/goquery"

	"github.com/gellel/emojipedia/arguments"
	"github.com/gellel/emojipedia/emoji"
	"github.com/gellel/emojipedia/slice"
	"github.com/gellel/emojipedia/text"
)

func emojiMain(arguments *arguments.Arguments) {
	e, err := emoji.Open(arguments.Get(0))
	switch err == nil {
	case true:
		switch strings.ToUpper(arguments.Next().Get(0)) {
		case "-A", ANCHOR:
			fmt.Println(e.Anchor)
		case "-C", CATEGORY:
			fmt.Println(e.Category)
		case "-CC", CODES:
			e.Codes.Each(func(_ int, i interface{}) {
				fmt.Println(i.(string))
			})
		case "-D", DESCRIPTION:
			if len(e.Description) == 0 {
				var (
					resp, _     = http.Get("https://emojipedia.org/" + e.Name + "/")
					document, _ = goquery.NewDocumentFromResponse(resp)
					re          = regexp.MustCompile(`\r?\n`)
					paragraphs  = &slice.Slice{}
				)
				document.Find("section.description > p").Each(func(_ int, selection *goquery.Selection) {
					paragraphs.Append(re.ReplaceAllString(strings.TrimSpace(selection.Text()), " "))
				})
				e.Description = paragraphs.Join(" ")
				emoji.Write(e)
			}
			fmt.Println(e.Description)
		case "-E", EMOJI:
			fmt.Println(text.Emojize(e.Unicode))
		case "-H", HREF:
			fmt.Println(e.Href)
		case "-I", IMAGE:
			fmt.Println(e.Image)
		case "-K", KEYWORDS:
			e.Keywords.Sort().Each(func(_ int, i interface{}) {
				fmt.Println(i.(string))
			})
		case "-N", NUMBER:
			fmt.Println(e.Number)
		case "-S", SUBCATEGORY:
			fmt.Println(e.Subcategory)
		case "-T", "TABLE":
			var (
				character   = text.Emojize(e.Unicode)
				category    = e.Category
				codes       = e.Codes.Join(" ")
				href        = e.Href
				keywords    = e.Keywords.Sort().Join(" ")
				name        = e.Name
				number      = fmt.Sprintf("%v", e.Number)
				subcategory = e.Subcategory
				template    = []string{
					character,
					category,
					codes,
					href,
					keywords,
					name,
					number,
					subcategory}
			)
			fmt.Fprintln(writer, "\t|category\t|codes\t|href\t|keywords\t|name\t|number\t|subcategory")
			fmt.Fprintln(writer, strings.Join(template, "\t|"))
			writer.Flush()
		case "-U", UNICODE:
			fmt.Println(e.Unicode)
		}
	default:
		fmt.Println(fmt.Sprintf("emoji \"%s\" not found. please try again", arguments.Get(0)))
	}
}
