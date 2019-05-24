package arguments

import (
	"fmt"

	"github.com/gellel/emojipedia/slice"
)

func New() *Arguments {
	return &Arguments{slice: &slice.Slice{}}
}

func NewArguments(args []string) *Arguments {
	arguments := New()
	for _, arg := range args {
		arguments.slice.Append(arg)
	}
	return arguments
}

type Arguments struct {
	slice *slice.Slice
}

func (pointer *Arguments) Get(i int) (argument string) {
	if property, ok := pointer.slice.Get(i); ok {
		argument = property.(string)
	}
	return argument
}

func (pointer *Arguments) Next() *Arguments {
	pointer.slice.Splice(1, pointer.slice.Len())
	return pointer
}

func (pointer *Arguments) String() string {
	return fmt.Sprintf("%s", *pointer.slice)
}
