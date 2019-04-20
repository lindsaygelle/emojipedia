package x_test

import (
	"fmt"
	"testing"

	"github.com/gellel/emojipedia/x"
)

func X(a int, b int) {}
func Y()             {}
func Test(t *testing.T) {

	fmt.Println((&x.Runner{}).Set(X, Y).Functions)
}
