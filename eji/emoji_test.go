package eji_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gellel/emojipedia/eji"
)

func Test(t *testing.T) {
	emoji := eji.Emoji{
		Category:    "Animals & Nature",
		Code:        "U+1f417",
		Keywords:    []string{"boar", "pig"},
		Name:        "boar",
		Number:      494,
		Sample:      "🐗",
		SubCategory: "animal-mammal",
		Unicode:     "\\U0001f417"}

	fmt.Println(strings.Replace(emoji.Unicode, "\\", "", 1))
}
