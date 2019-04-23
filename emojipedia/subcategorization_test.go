package emojipedia_test

import (
	"testing"

	"github.com/gellel/emojipedia/emojipedia"
	"github.com/gellel/emojipedia/files"
)

func TestSubcategorization(t *testing.T) {
	doc, exists := files.HTML(files.Unicode)
	if exists != true {
		t.Fatalf("required document does not exist")
	}
	subcategorization := emojipedia.NewSubcategorizationFromDocument(doc)
	ok := (subcategorization.Len() != 0)
	if ok != true {
		t.Fatalf("emojipedia.Subategorization is empty")
	}
}
