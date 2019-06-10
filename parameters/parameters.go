package parameters

import (
	"fmt"
	"strings"
	"text/tabwriter"
)

const (
	template string = "[%s|%s]\t%s"
)

type Parameter struct {
	About   string
	Short   string
	Verbose string
}

func NewParameter(short, verbose, about string) *Parameter {
	return &Parameter{
		About:   strings.ToLower(about),
		Short:   strings.ToLower(short),
		Verbose: verbose}
}

func (pointer *Parameter) Fprintln(writer *tabwriter.Writer) {
	fmt.Fprintln(writer, fmt.Sprintf(template, pointer.Short, pointer.Verbose, pointer.About))
}
