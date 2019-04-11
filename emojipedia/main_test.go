package emojipedia_test

import (
	"os"
	"testing"

	"github.com/gellel/emojipedia/emojipedia"
)

func Test(t *testing.T) {

	associative := &emojipedia.Associative{}

	emojidex := &emojipedia.Emojidex{
		"test": &emojipedia.Emoji{
			Category: "test"}}

	set := &emojipedia.Set{}

	if err := os.Remove(emojipedia.MarshallAssociative("test-associative", associative)); err != nil {
		panic(err)
	}
	if err := os.Remove(emojipedia.MarshallEmojidex("test-emojidex", emojidex)); err != nil {
		panic(err)
	}
	if err := os.Remove(emojipedia.MarshallSet("test-set", set)); err != nil {
		panic(err)
	}
}
