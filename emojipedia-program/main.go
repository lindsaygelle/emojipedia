package program

import (
	"bytes"
	"io/ioutil"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"strings"
)

var (
	re, _    = regexp.Compile(`\(([^()]+)\)`)
	replacer = strings.NewReplacer([]string{"[", "", "]", "", " ", ""}...)
)

func NewArgument(name, value string, position int, pointer uintptr, variadic bool, kind reflect.Kind) (argument *Argument) {
	return &Argument{
		Address:  pointer,
		Kind:     kind,
		Pointer:  strings.Index(value, "*") != -1,
		Position: position,
		Slice:    (kind.String() == "slice"),
		Name:     name,
		Value:    replacer.Replace(value),
		Variadic: variadic}
}

func NewArguments(reflection reflect.Type, pointer uintptr, variadic bool, parameters []string) (arguments *Arguments) {
	arguments = &Arguments{}
	for i, parameter := range parameters {
		in := reflection.In(i)
		substrings := strings.Split(parameter, " ")
		argument := NewArgument(substrings[0], in.String(), i, pointer, variadic, in.Kind())
		*arguments = append(*arguments, argument)
	}
	return arguments
}

func NewDeconstruct(f interface{}) *Deconstruct {
	reflection := reflect.TypeOf(f)
	pointer := reflect.ValueOf(f).Pointer()
	reference := runtime.FuncForPC(pointer)
	variadic := reflection.IsVariadic()
	name := filepath.Base(reference.Name())
	i := strings.Index(name, ".")
	for i > -1 {
		name = name[(i + 1):]
		i = strings.Index(name, ".")
	}
	parameters := NewParameters(reference.FileLine(pointer))
	return &Deconstruct{
		Parameters: parameters,
		Pointer:    pointer,
		Name:       name,
		Type:       reflection,
		Variadic:   variadic}
}

func NewFunction(f interface{}) (function *Function) {
	deconstruct := NewDeconstruct(f)
	arguments := NewArguments(deconstruct.Type, deconstruct.Pointer, deconstruct.Variadic, deconstruct.Parameters)
	length := len(*arguments)
	return &Function{
		Arguments: arguments,
		Empty:     length == 0,
		F:         f,
		Length:    length,
		Pointer:   deconstruct.Pointer,
		Name:      deconstruct.Name,
		Variadic:  deconstruct.Variadic}
}

func NewParameters(file string, line int) (arguments []string) {
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

func NewProgram(key string, f interface{}) *Program {
	return &Program{
		Function: NewFunction(f),
		Key:      key}
}