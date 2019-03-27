package store_test

import (
	"fmt"
	"testing"

	"github.com/gellel/emojipedia/eji"

	"github.com/gellel/emojipedia/store"
)

func Test(t *testing.T) {
	fmt.Println(store.Set("emoji-categories_test.json", &eji.Set{"test": []string{"test"}}))
}
