package x

import (
	"bytes"
	"fmt"
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
	Arguments *Args
	F         interface{}
	Pointer   uintptr
	Name      string
	Variadic  bool
}

type Runner map[string]*Function

func (args Args) Bounds(i int) (ok bool) {
	ok = ((i > -1) && (i < len(args)))
	return ok
}

func (args Args) Empty() (empty bool) {
	empty = (args.Length() == 0)
	return empty
}

func (args Args) Length() (length int) {
	length = (len(args))
	return length
}

func (args *Args) Peek(i int) (argument *Argument, ok bool) {
	if ok = args.Bounds(i); ok != false {
		argument = (*args)[i]
	}
	return argument, ok
}

func (args Args) Same() bool {
	var previous string
	for _, argument := range args {
		if len(previous) != 0 && previous != argument.Value {
			return false
		}
		previous = argument.Value
	}
	return true
}

func (args *Args) Push(argument *Argument) {
	*args = append(*args, argument)
}

func (runner Runner) Next(arguments []string) (caller func(arguments []string), ok bool) {
	var argument string
	if len(arguments) != 0 {
		argument = strings.ToUpper(arguments[0])
	}
	fmt.Println(argument)
	/*if function, ok := runner[argument]; ok {
		if function.Arguments.Length() {

		}
	}*/
	return caller, ok
}

func (runner Runner) Keys() (keys []string) {
	for key := range runner {
		keys = append(keys, key)
	}
	return keys
}

func args(reflection reflect.Type, pointer uintptr, variadic bool, parameters []string) *Args {
	args := &Args{}
	for i, parameter := range parameters {
		in := reflection.In(i)
		substrings := strings.Split(parameter, " ")
		argument := argument(substrings[0], in.String(), i, pointer, variadic, in.Kind())
		*args = append(*args, argument)
	}
	return args
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
	variadic := reflection.IsVariadic()
	name := filepath.Base(reference.Name())
	i := strings.Index(name, ".")
	for i > -1 {
		name = name[(i + 1):]
		i = strings.Index(name, ".")
	}
	parameters := parameters(reference.FileLine(pointer))
	arguments := args(reflection, pointer, variadic, parameters)
	return &Function{
		Arguments: arguments,
		F:         f,
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
