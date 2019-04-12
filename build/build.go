package build

import (
	"fmt"

	"github.com/gellel/emojipedia/emojipedia"
	"github.com/gellel/emojipedia/web"
)

var Options = []interface{}{all}

func all() {
	fmt.Println(fmt.Sprintf("getting data from %s. this could take awhile.", web.UnicodeOrgURL))
	document := web.Http(web.UnicodeOrgURL)
	e := emojipedia.NewEmojipediaFromDocument(document)
	dir := emojipedia.MarshallEmojidex("emoji", e.Emojidex)
	fmt.Println(fmt.Sprintf("dependencies successfully stored at %s", dir))
}

func categories() {

}

func keywords() {

}

func subcategories() {

}
