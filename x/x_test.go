package x

import "testing"

func X(a int, b int) {}
func Test(t *testing.T) {

	runner(X).Call([]string{"X"})
}
