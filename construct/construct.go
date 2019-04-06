package construct

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

type Arg struct {
	Kind      reflect.Kind
	Parameter string
	Pointer   uintptr
	Position  int
	Name      string
	Value     string
	Varadict  bool
}

type Func struct {
	Args    []*Arg
	Line    int
	Path    string
	Pointer uintptr
	Name    string
}

type Prog struct {
	Description string
	Funcs       []*Func
	Name        string
}

func NewArg(i int, pointer uintptr, parameter string, t reflect.Type) *Arg {
	substrings := strings.Split(strings.TrimSpace(parameter), " ")
	arg := Arg{
		Kind:      t.In(i).Kind(),
		Parameter: parameter,
		Pointer:   pointer,
		Position:  i,
		Name:      substrings[0],
		Value:     t.In(i).Kind().String(),
		Varadict:  t.IsVariadic()}
	return &arg
}

func NewFunc(fn interface{}) *Func {
	args := []*Arg{}
	t := reflect.TypeOf(fn)
	value := reflect.ValueOf(fn)
	pointer := value.Pointer()
	functionPointer := runtime.FuncForPC(pointer)
	name := filepath.Base(functionPointer.Name())
	name = name[(strings.Index(name, ".") + 1):]
	file, line := functionPointer.FileLine(pointer)
	b, err := ioutil.ReadFile(file)
	if err != nil {
		panic(err)
	}
	contents := bytes.Split(b, []byte("\n"))
	substring := string(contents[line-1])
	re := regexp.MustCompile(`\(([^()]+)\)`)
	matches := re.FindAllStringSubmatch(substring, 1)
	if len(matches) != 1 {
		panic(len(matches))
	}
	if len(matches[0]) != 2 {
		panic(len(matches[0]))
	}
	parameters := strings.Split(matches[0][1], ",")
	for i := 0; i < reflect.TypeOf(fn).NumIn(); i++ {
		args = append(args, NewArg(i, pointer, parameters[i], t))
	}
	f := Func{
		Args:    args,
		Line:    line,
		Path:    file,
		Pointer: pointer,
		Name:    name}
	return &f
}

func NewProg(name string, description string, functions []interface{}) *Prog {
	f := []*Func{}
	for _, fn := range functions {
		f = append(f, NewFunc(fn))
	}
	return &Prog{
		Description: description,
		Funcs:       f,
		Name:        name}
}

func NewVaradictString(arg *Arg) string {
	return fmt.Sprintf("[%s [...args]", arg.Name)
}

func NewFuncString(arg *Arg) string {
	return fmt.Sprintf("[%s=%s]", arg.Name, arg.Value)
}
