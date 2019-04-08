package cli

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"path"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"strings"
)

const lineLength int = 79

// Argument is a snapshot of a func parameter.
type Argument struct {
	Kind      reflect.Kind
	Parameter string
	Pointer   uintptr
	Position  int
	Name      string
	Slice     bool
	Value     string
	Varadict  bool
}

// Function is a snapshot of a Go function.
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

// Manifest is a collection of JSON data that explains a CLI function.
type Manifest struct {
	Description string `json:"description"`
	Name        string `json:"name"`
}

type Program struct {
	Description string
	Functions   []*Function
	Name        string
	Use         string
}

func NewArgument(i int, pointer uintptr, parameter string, t reflect.Type) *Argument {
	properties := t.In(i)
	value := strings.NewReplacer("[", "", "]", "").Replace(properties.String())
	slice := false
	if t.In(i).Kind().String() == "slice" {
		slice = true
	}
	substrings := strings.Split(strings.TrimSpace(parameter), " ")
	return &Argument{
		Kind:      properties.Kind(),
		Parameter: parameter,
		Pointer:   pointer,
		Position:  i,
		Name:      substrings[0],
		Slice:     slice,
		Value:     value,
		Varadict:  t.IsVariadic()}
}

func NewFunction(fn interface{}) *Function {
	arguments := []*Argument{}
	t := reflect.TypeOf(fn)
	value := reflect.ValueOf(fn)
	pointer := value.Pointer()
	functionPointer := runtime.FuncForPC(pointer)
	name := filepath.Base(functionPointer.Name())
	name = name[(strings.Index(name, ".") + 1):]
	substrings := regexp.MustCompile(`[A-Z]+[^A-Z]*|[^A-Z]+`).FindAllString(name, -1)
	if len(substrings) != 0 {
		name = strings.Join(substrings, "-")
	}
	file, line := functionPointer.FileLine(pointer)
	b, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	contents := bytes.Split(b, []byte("\n"))
	substring := string(contents[line-1])
	re := regexp.MustCompile(`\(([^()]+)\)`)
	matches := re.FindAllStringSubmatch(substring, 1)
	if len(matches) != 0 {
		parameters := strings.Split(matches[0][1], ",")
		for i := 0; i < reflect.TypeOf(fn).NumIn(); i++ {
			arguments = append(arguments, NewArgument(i, pointer, parameters[i], t))
		}
	}
	return &Function{
		Arguments: arguments,
		F:         fn,
		Line:      line,
		Path:      file,
		Pointer:   pointer,
		Name:      strings.ToLower(name),
		Varadict:  t.IsVariadic()}
}

func NewManifest(pc uintptr, file string, line int, ok bool) *Manifest {
	if ok != true {
		panic(fmt.Errorf("%s: line %v", file, line))
	}
	directory := path.Dir(file)
	fullpath := filepath.Join(directory, "manifest.json")
	content, err := ioutil.ReadFile(fullpath)
	if err != nil {
		panic(err)
	}
	manifest := &Manifest{}
	err = json.Unmarshal(content, manifest)
	if err != nil {
		panic(err)
	}
	return manifest
}

func NewProgram(name string, description string, functions []interface{}) *Program {
	f := []*Function{}
	for _, function := range functions {
		f = append(f, NewFunction(function))
	}
	return &Program{
		Description: description,
		Functions:   f,
		Name:        name,
		Use:         wrapUse(name, description, f)}
}

func NewProgramFromManifest(manifest *Manifest, functions []interface{}) *Program {
	return NewProgram(manifest.Name, manifest.Description, functions)
}

func getArgumentString(argument *Argument) string {
	if argument.Varadict {
		return fmt.Sprintf("%s [...%s]", argument.Name, argument.Value)
	}
	if argument.Slice {
		return fmt.Sprintf("%s=[...%s]", argument.Name, argument.Value)
	}
	return fmt.Sprintf("%s=<%s>", argument.Name, argument.Value)
}

func getFunctionString(function *Function) string {
	substrings := []string{}
	for _, argument := range function.Arguments {
		substrings = append(substrings, getArgumentString(argument))
	}
	usage := strings.Join(substrings, ", ")
	if len(usage) != 0 {
		return fmt.Sprintf("%s [%s]", function.Name, usage)
	}
	return fmt.Sprintf("--%s", function.Name)
}

func wrapDescription(paragraph string) string {
	var about string
	delimiter := " "
	cursor := 0
	for _, word := range strings.Split(paragraph, delimiter) {
		cursor = (cursor + len(word) + 1)
		about = fmt.Sprintf("%s%s%s", about, word, delimiter)
		if cursor >= lineLength {
			cursor = 0
			about = fmt.Sprintf("%s\n", about)
		}
	}
	return about
}

func wrapFunction(name string, functions []*Function) string {
	delimiter := " | "
	paragraphs := [][]string{[]string{}}
	prefix := fmt.Sprintf("usage: %s", name)
	offset := len(prefix)
	cursor := 0
	for _, function := range functions {
		i := len(paragraphs) - 1
		option := fmt.Sprintf("[%s]", getFunctionString(function))
		cursor = (len(strings.Join(paragraphs[i], delimiter)) + offset + len(option))
		if cursor >= lineLength {
			i = i + 1
			cursor = 0
			paragraphs = append(paragraphs, []string{})
		}
		paragraphs[i] = append(paragraphs[i], option)
	}
	first, paragraphs := paragraphs[0], paragraphs[1:]
	template := fmt.Sprintf("%s [%s", prefix, strings.Join(first, delimiter))
	for _, sentence := range paragraphs {
		var padding string
		substring := strings.Join(sentence, delimiter)
		j := 0
		for j < offset {
			padding = (padding + " ")
			j = (j + 1)
		}
		substring = fmt.Sprintf("\n %s%s", padding, substring)
		template = (template + substring)
	}
	template = fmt.Sprintf("%s]", template)
	return template
}

func wrapUse(name string, description string, functions []*Function) string {
	return fmt.Sprintf("%s\n\n%s", wrapDescription(description), wrapFunction(name, functions))
}
