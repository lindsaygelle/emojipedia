package emojipedia_test

import (
	"testing"

	"github.com/gellel/emojipedia/emojipedia"
	"github.com/gellel/emojipedia/files"
)

func TestEncyclopedia(t *testing.T) {
	doc, exists := files.HTML(files.Unicode)
	if exists != true {
		t.Fatalf("required document does not exist")
	}
	encyclopedia := emojipedia.NewEncyclopediaFromDocument(doc)
	ok := (encyclopedia.Len() != 0)
	if ok != true {
		t.Fatalf("emojipedia.Encyclopedia is empty")
	}
	ok = emojipedia.StoreEncyclopediaAsJSON(encyclopedia)
	if ok != true {
		t.Fatalf("emojipedia.Encyclopedia unable to deconstruct and store.")
	}
}
