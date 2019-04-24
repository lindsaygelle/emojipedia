package emojipedia_test

import (
	"testing"

	"github.com/gellel/emojipedia/emojipedia"
)

func TestSubcategorization(t *testing.T) {
	doc, ok := emojipedia.OpenUnicodesFile()
	if ok != true {
		t.Fatalf("%s does not exist", emojipedia.UnicodeFile)
	}
	subcategorization := emojipedia.NewSubcategorizationFromDocument(doc)
	ok = (subcategorization.Len() != 0)
	if ok != true {
		t.Fatalf("emojipedia.Subategorization is empty")
	}
}
