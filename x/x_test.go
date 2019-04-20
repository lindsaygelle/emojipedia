package x_test

import (
	"fmt"
	"testing"

	"github.com/gellel/emojipedia/x"
)

func X(a int, b int)     {}
func Y(i ...interface{}) {}
func Test(t *testing.T) {

	fmt.Println((&x.Runner{}).Set(Y).Get("Y"))
}
