package stdin

import (
	"fmt"
	"strings"
)

type Arg struct {
	About   string
	Short   string
	Verbose string
}

func (arg Arg) String() string {
	return fmt.Sprintf("%s\t%s", strings.ToLower(fmt.Sprintf("  [%s %s]", arg.Short, arg.Verbose)), arg.About)
}
