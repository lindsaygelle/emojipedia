package x

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"strings"
)

var re, _ = regexp.Compile(`\(([^()]+)\)`)

var replacements = []string{"[", "", "]", ""}

var replacer = strings.NewReplacer(replacements...)

type Args []*Argument

type Argument struct {
	Kind      reflect.Kind
	Parameter string
	Pointer   uintptr
	Position  int
	Name      string
	Slice     bool
	Value     string
	Variadic  bool
}

type Function struct {
	Arguments []*Argument
	Empty     bool
	Function  interface{}
	Length    int
	Pointer   uintptr
	Name      string
	Variadic  bool
}

type Runner map[string]*Function

func (runner Runner) Add(f *Function) {
	runner[f.Name] = f
}

func (runner Runner) Keys() (keys []string) {
	for key, _ := range runner {
		keys = append(keys, key)
	}
	return keys
}

func argument(name, value string, position int, pointer uintptr, variadic bool, kind reflect.Kind) (argument *Argument) {
	return &Argument{
		Kind:     kind,
		Pointer:  pointer,
		Position: position,
		Slice:    (kind.String() == "slice"),
		Name:     name,
		Value:    replacer.Replace(value),
		Variadic: variadic}
}

func function(f interface{}) (function *Function) {
	reflection := reflect.TypeOf(f)
	pointer := reflect.ValueOf(f).Pointer()
	reference := runtime.FuncForPC(pointer)
	name := filepath.Base(reference.Name())
	i := strings.Index(name, ".")
	for i > -1 {
		name = name[(i + 1):]
		i = strings.Index(name, ".")
	}
	parameters := parameters(reference.FileLine(pointer))
	length := len(parameters)
	arguments := make([]*Argument, length)
	variadic := reflection.IsVariadic()
	if len(parameters) != 0 {
		for i, parameter := range parameters {
			in := reflection.In(i)
			substrings := strings.Split(parameter, " ")
			arguments[i] = argument(substrings[0], in.String(), i, pointer, variadic, in.Kind())
		}
	}
	return &Function{
		Arguments: arguments,
		Empty:     (length == 0),
		Function:  f,
		Length:    length,
		Pointer:   pointer,
		Name:      name,
		Variadic:  variadic}
}

func parameters(file string, line int) (arguments []string) {
	b, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	contents := bytes.Split(b, []byte("\n"))
	substring := string(contents[line-1])
	matches := re.FindAllStringSubmatch(substring, 1)
	if len(matches) != 0 {
		arguments = strings.Split(matches[0][1], ",")
	}
	return arguments
}

func runner(f ...interface{}) (runner *Runner) {
	runner = &Runner{}
	for _, f := range f {
		runner.Add(function(f))
	}
	return runner
}
