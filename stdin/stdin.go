package stdin

import (
	"fmt"
)

type Arg struct {
	About   string
	Short   string
	Verbose string
}

func (arg Arg) String() string {
	return fmt.Sprintf("  %s %s\t%s", arg.Short, arg.Verbose, arg.About)
}
