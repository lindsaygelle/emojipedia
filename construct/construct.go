package construct

import (
	"bytes"
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

func NewFunc(fn interface{}) *Func {
	args := []*Arg{}
	t := reflect.TypeOf(fn)
	value := reflect.ValueOf(fn)
	pointer := value.Pointer()
	functionPointer := runtime.FuncForPC(pointer)
	name := filepath.Base(functionPointer.Name())
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
	parameters := strings.Split(matches[0][1], " ")
	for i := 0; i < reflect.TypeOf(fn).NumIn(); i++ {
		args = append(args, NewArg(i, pointer, parameters[i], t))
	}
	name = name[(strings.Index(name, ".") + 1):]
	f := Func{
		Args:    args,
		Line:    line,
		Path:    file,
		Pointer: pointer,
		Name:    name}
	return &f
}

func NewArg(i int, pointer uintptr, parameter string, t reflect.Type) *Arg {
	value := t.In(i)
	replacements := []string{"slice", "", value.String(), "", ",", ""}
	replacer := strings.NewReplacer(replacements...)
	arg := Arg{
		Kind:      value.Kind(),
		Parameter: parameter,
		Pointer:   pointer,
		Position:  i,
		Name:      strings.TrimSpace(replacer.Replace(parameter)),
		Value:     value.Kind().String(),
		Varadict:  t.IsVariadic()}
	return &arg
}
