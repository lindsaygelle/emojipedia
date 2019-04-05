package cmd

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
)

const veridict string = "args"

type program struct {
	Description string
	Name        string
	Routines    []routines
}

type routines struct {
	Args        []string
	Description string
	Program     interface{}
	Name        string
}

func getUsageOptionStrings(prefix string, fn []interface{}) {
	paragraphs := [][]string{[]string{}}
	offset := len(prefix)
	current := 0
	max := 79
	for _, f := range fn {
		name := getFunctionName(f)
		name = getFunctionOptions(name, reflect.TypeOf(f))
		current = offset + current + len(name) + 1
		if current >= max {
			current = 0
			paragraphs = append(paragraphs, []string{})
		}
		paragraphs[len(paragraphs)-1] = append(paragraphs[len(paragraphs)-1], name)
	}
	first, paragraphs := paragraphs[0], paragraphs[1:]
	template := fmt.Sprintf("%s [%s ", prefix, strings.Join(first, ","))
	for _, p := range paragraphs {
		padding := ""
		substring := strings.Join(p, ",")
		j := 0
		for j < offset {
			padding = padding + " "
			j = j + 1
		}
		substring = fmt.Sprintf("\n%s%s", padding, substring)
		template = template + substring
	}
	template = template + "]"
	fmt.Println(template)
}

func getFunction(name string) string {
	return fmt.Sprintf("[%s [...%s]]", name, veridict)
}

func getFunctionOptions(name string, t reflect.Type) string {
	arguments := []string{}
	for i := 0; i < t.NumIn(); i = (i + 1) {
		option := (t.In(i).Kind().String())
		arguments = append(arguments, option)
	}
	options := strings.Join(arguments, ",")
	if len(options) != 0 {
		options = fmt.Sprintf("=[%s]", options)
	}
	return fmt.Sprintf("[--%s%s]", name, options)
}

func getFunctionName(fn interface{}) string {
	pointer := (reflect.ValueOf(fn).Pointer())
	namespace := (runtime.FuncForPC(pointer).Name())
	name := filepath.Base(namespace)
	return name[(strings.Index(filepath.Base(namespace), ".") + 1):]
}

func get(property string) {}

func search(category string) {}

func version() {}

func help() {}

func Get() {
	getUsageOptionStrings("usage: emojipedia ", []interface{}{version, help, get, search})
}
