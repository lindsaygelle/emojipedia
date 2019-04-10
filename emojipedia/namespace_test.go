package emojipedia_test

import (
	"testing"

	"github.com/gellel/emojipedia/emojipedia"
)

func TestNamespace(t *testing.T) {
	namespace := emojipedia.Namespace{}
	namespace.Push("b")
	namespace.Unshift("a")
	value, ok := namespace.PeekFirst()
	if ok != true {
		t.Fatalf("cannot access start")
	}
	if value != "a" {
		t.Fatalf("%s != a", value)
	}
	value, ok = namespace.PeekLast()
	if ok != true {
		t.Fatalf("cannot access end")
	}
	if value != "b" {
		t.Fatalf("%s != b", value)
	}
}
