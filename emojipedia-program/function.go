package program

type Function struct {
	Arguments *Arguments
	Empty     bool
	F         interface{}
	Length    int
	Pointer   uintptr
	Name      string
	Variadic  bool
}

func (function *Function) Set(f interface{}) *Function {
	*function = *NewFunction(f)
	return function
}
