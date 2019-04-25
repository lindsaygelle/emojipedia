package program

import "reflect"

type Argument struct {
	Address   uintptr
	Kind      reflect.Kind
	Parameter string
	Pointer   bool
	Position  int
	Name      string
	Slice     bool
	Value     string
	Variadic  bool
}
