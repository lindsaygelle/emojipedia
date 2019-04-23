package emojipedia_test

import (
	"testing"

	"github.com/gellel/emojipedia/emojipedia"
)

func TestSubcategorization(t *testing.T) {
	doc, ok := emojipedia.OpenUnicodesFromFile()
	if ok != true {
		t.Fatalf("unicode.html does not exist")
	}
	subcategorization := emojipedia.NewSubcategorizationFromDocument(doc)
	ok = (subcategorization.Len() != 0)
	if ok != true {
		t.Fatalf("emojipedia.Subategorization is empty")
	}
}
