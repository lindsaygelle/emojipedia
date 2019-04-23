package emojipedia_test

import (
	"fmt"
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
	emoji, ok = emojipedia.OpenEmojiFromFile(emoji.Name)
	fmt.Println(emoji)
	if ok != true {
		t.Fatalf("emojipedia.OpenEmojiFromFile did not open NIL.json")
	}
	ok = emojipedia.RemoveEmojiJSON(emoji.Name)
	if ok != true {
		t.Fatalf("emojipedia.RemoveEmojiJSON did not remove NIL.json")
	}
}
