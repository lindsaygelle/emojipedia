package program

import "reflect"

type Deconstruct struct {
	Parameters []string
	Pointer    uintptr
	Name       string
	Type       reflect.Type
	Variadic   bool
}
