package emojipedia_test

import (
	"testing"

	"github.com/gellel/emojipedia/emojipedia"
)

func TestEmoji(t *testing.T) {
	emoji := emojipedia.Emoji{
		Category:    "TEST",
		Code:        "U+000",
		Keywords:    []string{"A", "B", "C"},
		Name:        "TEST",
		Number:      -1,
		Sample:      "NIL",
		Subcategory: "TEST",
		Unicode:     "X"}

	emoji.Println()
}
