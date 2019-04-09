package cli_test

import (
	"fmt"
	"testing"

	"github.com/gellel/emojipedia/cli"
)

func A(a ...string)                  {}
func B(b int, c float32, d []string) {}

func Test(t *testing.T) {

	a := cli.NewFunction(A)

	b := cli.NewFunction(B)

	if a.Name != "A" {
		t.Error(fmt.Sprintf("cli.go a.Name != A"))
	}
	if len(a.Arguments) != 1 {
		t.Error(fmt.Sprintf("cli.go len(a.Arguments) != 1"))
	}
	for _, arg := range a.Arguments {
		if arg.Varadict != true {
			t.Errorf(fmt.Sprintf("cli.go %s.Varadict != true", arg.Name))
		}
	}
	if b.Name != "B" {
		t.Error(fmt.Sprintf("cli.go a.Name != B"))
	}
	if len(b.Arguments) != 3 {
		t.Error(fmt.Sprintf("cli.go len(b.Arguments) != 2"))
	}
	for i, arg := range b.Arguments {
		if i != arg.Position {
			t.Errorf(fmt.Sprintf("cli.go %s.Position != %v", arg.Name, arg.Position))
		}
		switch i {
		case 0:
			if arg.Value != "int" {
				t.Errorf(fmt.Sprintf("cli.go %s.Value != int", arg.Name))
			}
		case 1:
			if arg.Value != "float32" {
				t.Errorf(fmt.Sprintf("cli.go %s.Value != float32", arg.Name))
			}
		}
	}
}
