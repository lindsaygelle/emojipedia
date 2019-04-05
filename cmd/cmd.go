package cmd

import (
	"fmt"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
)

const columnLength int = 79

type Manifest struct {
	Description string `json:"description"`
	Name        string `json:"name"`
	Version     int    `json:"version"`
}

type Program struct {
	Description string
	Usage       string
}

func NewProgram(name string, description string, fn []interface{}) *Program {
	return &Program{
		Description: makeDescription(description),
		Usage:       makeUsage(name, fn)}
}

func NewFromManifest(manifest *Manifest) *Program {
	return &Program{}
}

func makeDescription(paragraph string) string {
	var description string
	current := 0
	for _, p := range strings.Split(paragraph, " ") {
		current = (current + len(p) + 1)
		description = fmt.Sprintf("%s%s ", description, p)
		if current >= columnLength {
			current = 0
			description = (description + "\n")
		}
	}
	return description
}

func makeUsage(program string, fn []interface{}) string {
	paragraphs := [][]string{[]string{}}
	prefix := strings.Join([]string{(program + ":"), "usage"}, " ")
	offset := len(prefix)
	current := 0
	for _, f := range fn {
		name := getNameOf(f)
		option := makeOption(name, reflect.TypeOf(f))
		current = (offset + current + len(option))
		if current >= columnLength {
			current = 0
			paragraphs = append(paragraphs, []string{})
		}
		paragraphs[len(paragraphs)-1] = append(paragraphs[len(paragraphs)-1], option)
	}
	first, paragraphs := paragraphs[0], paragraphs[1:]
	template := fmt.Sprintf("%s [%s", prefix, strings.Join(first, ","))
	for _, p := range paragraphs {
		var padding string
		substring := strings.Join(p, ",")
		j := 0
		for j < offset {
			padding = (padding + " ")
			j = (j + 1)
		}
		substring = fmt.Sprintf("\n %s%s", padding, substring)
		template = (template + substring)
	}
	template = (template + "]")
	return template
}

func makeOption(name string, t reflect.Type) string {
	var substring string
	switch t.IsVariadic() {
	case true:
		substring = fmt.Sprintf("[%s [...args]]", name)
	default:
		arguments := []string{}
		for i := 0; i < t.NumIn(); i = (i + 1) {
			option := (t.In(i).Kind().String())
			arguments = append(arguments, option)
		}
		options := strings.Join(arguments, ",")
		if len(options) != 0 {
			options = fmt.Sprintf("=[%s]", options)
		}
		substring = fmt.Sprintf("[--%s%s]", name, options)
	}
	return substring
}

func getNameOf(i interface{}) string {
	pointer := (reflect.ValueOf(i).Pointer())
	namespace := (runtime.FuncForPC(pointer).Name())
	name := filepath.Base(namespace)
	name = name[(strings.Index(filepath.Base(namespace), ".") + 1):]
	return strings.ToLower(name)
}
