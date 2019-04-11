package emojipedia_test

import (
	"os"
	"testing"

	"github.com/gellel/emojipedia/emojipedia"
)

func Test(t *testing.T) {

	emojidex := &emojipedia.Emojidex{}

	associative := &emojipedia.Associative{}

	if err := os.Remove(emojipedia.MarshallAssociative("test-associative", associative)); err != nil {
		panic(err)
	}
	if err := os.Remove(emojipedia.MarshallEmojidex("test-emojidex", emojidex)); err != nil {
		panic(err)
	}
}
