package emojipedia_test

import (
	"testing"

	"github.com/gellel/emojipedia/emojipedia"
)

func TestSet(t *testing.T) {
	set := emojipedia.Set{}
	if value := set.Add(1, "a"); value != "a" {
		t.Fatalf("%s != a", value)
	}
	if ok := set.Has(1); ok != true {
		t.Fatalf("ok != true")
	}
	if value, ok := set.Get(1); value != "a" || ok != true {
		t.Fatalf("ok != true || %s != a", value)
	}
}
