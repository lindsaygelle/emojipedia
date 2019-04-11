package emojipedia_test

import (
	"testing"

	"github.com/gellel/emojipedia/emojipedia"
)

func TestEmojidex(t *testing.T) {
	emoji := emojipedia.Emoji{
		Category:    "TEST",
		Code:        "U+000",
		Keywords:    []string{"A", "B", "C"},
		Name:        "TEST",
		Number:      -1,
		Sample:      "NIL",
		Subcategory: "TEST",
		Unicode:     "NIL"}
	emojidex := emojipedia.Emojidex{}
	emojidex.Add(emoji.Name, &emoji)
	if ok := emojidex.Has(emoji.Name); ok != true {
		t.Fatalf("ok != true")
	}
	if _, ok := emojidex.Get(emoji.Name); ok != true {
		t.Fatalf("ok != true")
	}
	if e := emojidex.GetUnsafely(emoji.Name); e != &emoji {
		t.Fatalf("(*%p) != (*%p)", &e, &emoji)
	}
	if ok := emojidex.Remove(emoji.Name); ok != true {
		t.Fatalf("ok != true")
	}
	if l := emojidex.Length(); l != 0 {
		t.Fatalf("%v != 0", l)
	}
}
