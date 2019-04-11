package cli

// A Function is a snapshot of a Go function.
// Each function holds a collection of Argument structs.
// If a function is a varadict function, it will only contain one argument.
// Function structs should be created using the NewFunction method.
type Function struct {
	Arguments []*Argument
	F         interface{}
	Line      int
	Path      string
	Pointer   uintptr
	Name      string
	Varadict  bool
}
