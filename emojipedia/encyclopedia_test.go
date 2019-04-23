package emojipedia_test

import (
	"testing"

	"github.com/gellel/emojipedia/emojipedia"
)

func TestEncyclopedia(t *testing.T) {
	doc, ok := emojipedia.OpenUnicodesFromFile()
	if ok != true {
		t.Fatalf("unicode.html does not exist")
	}
	encyclopedia := emojipedia.NewEncyclopediaFromDocument(doc)
	ok = (encyclopedia.Len() != 0)
	if ok != true {
		t.Fatalf("emojipedia.Encyclopedia is empty")
	}
}
