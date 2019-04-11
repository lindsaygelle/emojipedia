package emojipedia_test

import (
	"testing"

	"github.com/gellel/emojipedia/emojipedia"
)

func TestAssociative(t *testing.T) {
	associative := &emojipedia.Associative{}

	associative.Set("a")

	if ok := associative.Has("a"); ok != true {
		t.Fatalf("ok != true")
	}
	associative.Add("a", "a")
	if length := associative.GetUnsafely("a").Length(); length != 1 {
		t.Fatalf("%v != 1", length)
	}
	if value, position := associative.GetUnsafely("a").Search("a"); value != "a" || position != 0 {
		t.Fatalf("%v != 0 || %s != a", position, value)
	}
}
