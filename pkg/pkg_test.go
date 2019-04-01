package pkg_test

import (
	"fmt"
	"testing"

	"github.com/gellel/emojipedia/pkg"
)

func Test(t *testing.T) {

	p, err := pkg.Get()
	if err != nil {
		panic(err)
	}
	fmt.Println(p)
}
