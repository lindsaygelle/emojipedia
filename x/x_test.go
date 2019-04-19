package x

import "testing"
import "fmt"

func X(a int, b int) {}
func Test(t *testing.T) {

	fmt.Println(function(X).Arguments.Same())

	runner(X).Next([]string{"X"})
}
