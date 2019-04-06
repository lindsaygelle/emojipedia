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

type Function struct {
	Arguments []*Argument
	Line      int
	Path      string
	Pointer   uintptr
	Name      string
	Varadict  bool
}

type Program struct {
	Description string
	Functions   []*Function
	Name        string
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

func NewDescription(sentence string, functions []*Function) string {

	for _, word := range strings.Split(sentence, " ") {
		fmt.Println(word)
	}

	return ""
}

func NewFunction(fn interface{}) *Function {
	arguments := []*Argument{}
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
		arguments = append(arguments, NewArgument(i, pointer, parameters[i], t))
	}
	return &Function{
		Arguments: arguments,
		Line:      line,
		Path:      file,
		Pointer:   pointer,
		Name:      name,
		Varadict:  t.IsVariadic()}
}

func NewProgram(name string, description string, functions []interface{}) *Program {
	f := []*Function{}
	for _, function := range functions {
		f = append(f, NewFunction(function))
	}
	return &Program{
		Description: description,
		Functions:   f,
		Name:        name}
}

func (argument *Argument) Usage() string {
	if argument.Slice {
		if argument.Varadict {
			return fmt.Sprintf("%s [...%s]", argument.Name, argument.Value)
		}
		return fmt.Sprintf("%s=[...%s]", argument.Name, argument.Value)
	}
	return fmt.Sprintf("%s=%s", argument.Name, argument.Value)
}

func (function *Function) Usage() string {
	substrings := []string{}
	for _, argument := range function.Arguments {
		substrings = append(substrings, argument.Usage())
	}
	usage := strings.Join(substrings, ", ")
	return fmt.Sprintf("%s [%s]", function.Name, usage)
}

func (function *Function) Slice() []string {
	arguments := []string{}
	for _, argument := range function.Arguments {
		arguments = append(arguments, argument.Usage())
	}
	return arguments
}

func (program *Program) Usage() string {
	substrings := []string{}
	for _, function := range program.Functions {
		substrings = append(substrings, function.Usage())
	}
	return fmt.Sprintf("[%s]", strings.Join(substrings, " | "))
}
