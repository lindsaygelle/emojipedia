package x

import "testing"

func X(a int, b int) {}
func Test(t *testing.T) {

	var A func(a []int)

	var B func(a int)

	A = func(a []int) {}

	B = func(a int) {}

	function(X)

	function(A)

	runner(A, B)
}
