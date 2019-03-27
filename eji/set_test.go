package eji_test

import (
	"fmt"
	"testing"

	"github.com/gellel/emojipedia/eji"
)

func TestSet(t *testing.T) {

	emoji := &eji.Emoji{
		Category: "test",
		Name:     "test"}

	set := eji.Set{}

	fmt.Println(set.Add(emoji.Category, emoji.Name).Has(emoji.Category))

	fmt.Println(set.Get(emoji.Category))

}
