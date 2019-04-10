package emojipedia

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

var _ emoji = (*Emoji)(nil)

type emoji interface {
	Keys() []string
	Println()
	Values() []string
}

type Emoji struct {
	Category    string   `json:"Category"`
	Code        string   `json:"Code"`
	Keywords    []string `json:"Keywords"`
	Name        string   `json:"Name"`
	Number      int      `json:"Number"`
	Sample      string   `json:"Sample"`
	Subcategory string   `json:"Subcategory"`
	Unicode     string   `json:"Unicode"`
}

func (emoji *Emoji) Keys() (keys []string) {
	keys = []string{
		"category",
		"code",
		"keywords",
		"name",
		"number",
		"sample",
		"subcategory",
		"unicode"}
	return keys
}

func (emoji *Emoji) Values() (values []string) {
	values = []string{
		emoji.Category,
		emoji.Code,
		strings.Join(emoji.Keywords, ","),
		emoji.Name,
		string(emoji.Number),
		emoji.Sample,
		emoji.Subcategory,
		emoji.Unicode}
	return values
}

func (emoji *Emoji) Println() {
	writer := new(tabwriter.Writer)
	writer.Init(os.Stdout, 0, 0, 0, ' ', tabwriter.Debug|tabwriter.AlignRight)
	fmt.Fprintln(writer, strings.Join(emoji.Keys(), "\t")+"\t")
	fmt.Fprintln(writer, strings.Join(emoji.Values(), "\t")+"\t")
	fmt.Fprintln(writer)
	writer.Flush()
}
