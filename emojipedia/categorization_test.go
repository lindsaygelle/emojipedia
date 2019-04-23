package emojipedia_test

import (
	"testing"

	"github.com/gellel/emojipedia/emojipedia"
	"github.com/gellel/emojipedia/files"
)

func TestCategorization(t *testing.T) {
	doc, exists := files.HTML(files.Unicode)
	if exists != true {
		t.Fatalf("required document does not exist")
	}
	categorization := emojipedia.NewCategorizationFromDocument(doc)
	ok := (categorization.Len() != 0)
	if ok != true {
		t.Fatalf("emojipedia.Categorization is empty")
	}
}
