package emojipedia

import (
	"fmt"
	"os"
	"reflect"
	"strconv"
	"strings"
	"text/tabwriter"
)

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

func (emoji *Emoji) Fields() (fields []string) {
	r := reflect.ValueOf(emoji).Elem()
	fields = []string{}
	for i := 0; i < r.NumField(); i++ {
		fields = append(fields, r.Type().Field(i).Name)
	}
	return fields
}

func (emoji *Emoji) Values() (values []string) {
	r := reflect.ValueOf(emoji).Elem()
	values = []string{}
	for i := 0; i < r.NumField(); i++ {
		in := r.Field(i).Interface()
		switch in.(type) {
		case []string:
			values = append(values, strings.Join((in.([]string)), ","))
		case int:
			values = append(values, strconv.Itoa((in.(int))))
		default:
			values = append(values, (in.(string)))
		}
	}
	return values
}

func (emoji *Emoji) Println() {
	fields := emoji.Fields()
	values := emoji.Values()
	writer := new(tabwriter.Writer)
	writer.Init(os.Stdout, 0, 0, 0, ' ', tabwriter.Debug|tabwriter.AlignRight)
	fmt.Fprintln(writer, strings.Join(fields, "\t")+"\t")
	fmt.Fprintln(writer, strings.Join(values, "\t")+"\t")
	writer.Flush()
}
