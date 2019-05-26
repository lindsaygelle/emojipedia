package arguments

import (
	"fmt"

	"github.com/gellel/emojipedia/slice"
)

// New instantiates a new Arguments pointer.
func New() *Arguments {
	return &Arguments{slice: &slice.Slice{}}
}

// NewArguments builds an argument iterator from a collection of strings.
// Assumes that os.Args[1:] is to be passed into the constructor function.
func NewArguments(args []string) *Arguments {
	arguments := New()
	for _, arg := range args {
		arguments.slice.Append(arg)
	}
	return arguments
}

// Arguments provides an iterator to move through the arguments fetched in from os.Args
type Arguments struct {
	slice *slice.Slice
}

// Get safely accesses the argument at the iteration index.
func (pointer *Arguments) Get(i int) (argument string) {
	if property, ok := pointer.slice.Get(i); ok {
		argument = property.(string)
	}
	return argument
}

// Next unshifts the first element of the Arguments struct and returns the modified struct.
func (pointer *Arguments) Next() *Arguments {
	pointer.slice.Splice(1, pointer.slice.Len())
	return pointer
}

func (pointer *Arguments) String() string {
	return fmt.Sprintf("%s", *pointer.slice)
}
