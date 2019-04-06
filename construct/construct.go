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
	Slice     bool
	Value     string
	Varadict  bool
}

type Func struct {
	Args     []*Arg
	Line     int
	Path     string
	Pointer  uintptr
	Name     string
	Varadict bool
}

type Program struct {
	Description string
	Funcs       []*Func
	Name        string
}

func NewArg(i int, pointer uintptr, parameter string, t reflect.Type) *Arg {
	properties := t.In(i)
	value := strings.NewReplacer("[", "", "]", "").Replace(properties.String())
	slice := false
	if t.In(i).Kind().String() == "slice" {
		slice = true
	}
	substrings := strings.Split(strings.TrimSpace(parameter), " ")
	arg := Arg{
		Kind:      properties.Kind(),
		Parameter: parameter,
		Pointer:   pointer,
		Position:  i,
		Name:      substrings[0],
		Slice:     slice,
		Value:     value,
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
		Args:     args,
		Line:     line,
		Path:     file,
		Pointer:  pointer,
		Name:     name,
		Varadict: t.IsVariadic()}
	return &f
}

func NewProgram(name string, description string, functions []interface{}) *Program {
	f := []*Func{}
	for _, fn := range functions {
		f = append(f, NewFunc(fn))
	}
	return &Program{
		Description: description,
		Funcs:       f,
		Name:        name}
}

func (arg *Arg) Usage() string {
	var str string
	switch arg.Slice {
	case true:
		switch arg.Varadict {
		case true:
			str = fmt.Sprintf("%s [...%s]", arg.Name, arg.Value)
		default:
			str = fmt.Sprintf("%s=[...%s]", arg.Name, arg.Value)
		}
	default:
		str = fmt.Sprintf("%s=%s", arg.Name, arg.Value)
	}
	return str
}
