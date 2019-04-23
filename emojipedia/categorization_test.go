package emojipedia_test

import (
	"testing"

	"github.com/gellel/emojipedia/emojipedia"
)

func TestCategorization(t *testing.T) {
	doc, ok := emojipedia.OpenUnicodesFromFile()
	if ok != true {
		t.Fatalf("unicode.html does not exist")
	}
	categorization := emojipedia.NewCategorizationFromDocument(doc)
	ok = (categorization.Len() != 0)
	if ok != true {
		t.Fatalf("emojipedia.Categorization is empty")
	}
}
