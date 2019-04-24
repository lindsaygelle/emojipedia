package emojipedia_test

import (
	"testing"

	"github.com/gellel/emojipedia/emojipedia"
)

func TestEmoji(t *testing.T) {
	emoji := &emojipedia.Emoji{
		Name: "NIL"}
	ok := emojipedia.StoreEmojiAsJSON(emoji)
	if ok != true {
		t.Fatalf("emojipedia.StoreEmojiAsJSON did not NIL.json")
	}
	emoji, ok = emojipedia.OpenEmojiFile(emoji.Name)
	if ok != true {
		t.Fatalf("emojipedia.OpenEmojiFromFile did not open NIL.json")
	}
	ok = emojipedia.RemoveEmojiFile(emoji.Name)
	if ok != true {
		t.Fatalf("emojipedia.RemoveEmojiFile did not remove NIL.json")
	}
}
