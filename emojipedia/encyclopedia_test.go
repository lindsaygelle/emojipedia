package emojipedia_test

import (
	"testing"

	"github.com/gellel/emojipedia/emojipedia"
)

func TestEncyclopedia(t *testing.T) {
	doc, ok := emojipedia.OpenUnicodesFile()
	if ok != true {
		t.Fatalf("%s does not exist", emojipedia.UnicodeFile)
	}
	encyclopedia := emojipedia.NewEncyclopediaFromDocument(doc)
	ok = (encyclopedia.Len() != 0)
	if ok != true {
		t.Fatalf("emojipedia.Encyclopedia is empty")
	}
}
