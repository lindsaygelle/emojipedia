package eji_test

import (
	"fmt"
	"testing"

	"github.com/gellel/emojipedia/eji"
)

func TestEmoji(t *testing.T) {
	emoji := eji.Emoji{
		Category:    "Animals & Nature",
		Code:        "U+1f417",
		Keywords:    []string{"boar", "pig"},
		Name:        "boar",
		Number:      494,
		Sample:      "🐗",
		SubCategory: "animal-mammal",
		Unicode:     "\\U0001f417"}

	fmt.Println(emoji.Sample)
}
