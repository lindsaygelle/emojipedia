package emojipedia

import (
	"fmt"
	"os"
	"text/tabwriter"
)

type Feature struct {
	About    string    `json:"about"`
	Sections []Section `json:"sections"`
}

func (feature Feature) Describe(sections ...Section) {
	writer := new(tabwriter.Writer).Init(
		os.Stdout, 0, 8, 0, '\t', 0)
	for _, section := range sections {
		fmt.Println(section.About)
		for _, argument := range section.Arguments {
			fmt.Fprintln(writer, argument.String())
		}
		fmt.Fprintln(writer)
	}
	writer.Flush()
}
