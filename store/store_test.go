package store_test

import (
	"fmt"
	"os"
	"testing"

	"github.com/gellel/emojipedia/eji"

	"github.com/gellel/emojipedia/store"
)

func Test(t *testing.T) {
	path, err := store.Save("test", &eji.Set{"test": []string{"test"}})
	if err != nil {
		panic(err)
	}
	err = store.Has("test")
	if err != nil {
		panic(err)
	}
	value, err := store.Open("test")
	if err != nil {
		panic(err)
	}
	fmt.Println(value)
	err = os.Remove(path)
	if err != nil {
		panic(err)
	}
}
