package construct_test

import (
	"fmt"
	"testing"

	"github.com/gellel/emojipedia/construct"
)

func A(a ...string)                  {}
func B(b int, c float32, d []string) {}

func Test(t *testing.T) {

	a := construct.NewFunc(A)

	b := construct.NewFunc(B)

	if a.Name != "A" {
		t.Error(fmt.Sprintf("construct.go a.Name != A"))
	}
	if len(a.Args) != 1 {
		t.Error(fmt.Sprintf("construct.go len(a.Args) != 1"))
	}
	for _, arg := range a.Args {
		if arg.Varadict != true {
			t.Errorf(fmt.Sprintf("construct.go %s.Varadict != true", arg.Name))
		}
	}
	if b.Name != "B" {
		t.Error(fmt.Sprintf("construct.go a.Name != B"))
	}
	if len(b.Args) != 3 {
		t.Error(fmt.Sprintf("construct.go len(b.Args) != 2"))
	}
	for i, arg := range b.Args {
		fmt.Println(arg.Usage())
		if i != arg.Position {
			t.Errorf(fmt.Sprintf("construct.go %s.Position != %v", arg.Name, arg.Position))
		}
		switch i {
		case 0:
			if arg.Value != "int" {
				t.Errorf(fmt.Sprintf("construct.go %s.Value != int", arg.Name))
			}
		case 1:
			if arg.Value != "float32" {
				t.Errorf(fmt.Sprintf("construct.go %s.Value != float32", arg.Name))
			}
		}
	}

	//construct.NewHelpString(construct.NewProg("name", "desc", []interface{}{A, B}))
}
