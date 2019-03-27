package pkg_test

import (
	"testing"

	"github.com/gellel/emojipedia/pkg"
)

func Test(t *testing.T) {

	if err := pkg.Get(); err != nil {
		panic(err)
	}
}
