package describe_test

import (
	"fmt"
	"testing"

	"github.com/gellel/emojipedia/describe"
)

func Test(t *testing.T) {

	a := describe.String("cache")
	b := describe.String("information")

	c := describe.Key{Literal: a}
	d := describe.Key{Literal: b}

	fmt.Println(describe.Keys{}.Append(c, d).Longest())

}
