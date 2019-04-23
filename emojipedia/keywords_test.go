package emojipedia_test

import (
	"testing"

	"github.com/gellel/emojipedia/emojipedia"
	"github.com/gellel/emojipedia/files"
)

func TestKeywords(t *testing.T) {
	doc, exists := files.HTML(files.Unicode)
	if exists != true {
		t.Fatalf("required document does not exist")
	}
	keywords := emojipedia.NewKeywordsFromDocument(doc)
	ok := (keywords.Len() != 0)
	if ok != true {
		t.Fatalf("emojipedia.Keywords is empty")
	}
}
