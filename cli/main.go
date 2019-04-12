package cli

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
	"os/exec"
	"path/filepath"
	"reflect"
	"regexp"
	"runtime"
	"strconv"
	"strings"
)

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

func NewProgram(name string, description string, functions []interface{}) *Program {
	f := []*Function{}
	for _, function := range functions {
		f = append(f, NewFunction(function))
	}
	return &Program{
		Description: description,
		Functions:   f,
		Name:        name,
		Use:         WrapUse(name, description, f)}
}

func NewWindow() *Window {
	cmd := exec.Command("stty", "size")
	cmd.Stdin = os.Stdin
	out, err := cmd.Output()
	input, err := string(out), err
	parts := strings.Split(input, " ")
	y, err := strconv.Atoi(parts[0])
	if err != nil {
		panic(err)
	}
	x, err := strconv.Atoi(strings.Replace(parts[1], "\n", "", 1))
	if err != nil {
		panic(err)
	}
	return &Window{
		Height: int(y),
		Width:  int(x)}
}

func GetArgumentString(argument *Argument) string {
	if argument.Varadict {
		return fmt.Sprintf("%s [...%s]", argument.Name, argument.Value)
	}
	if argument.Slice {
		return fmt.Sprintf("%s=[...%s]", argument.Name, argument.Value)
	}
	return fmt.Sprintf("%s=<%s>", argument.Name, argument.Value)
}

func GetFunctionString(function *Function) string {
	substrings := []string{}
	for _, argument := range function.Arguments {
		substrings = append(substrings, strings.ToLower(GetArgumentString(argument)))
	}
	usage := strings.Join(substrings, ", ")
	if len(usage) != 0 {
		return fmt.Sprintf("%s [%s]", function.Name, usage)
	}
	return fmt.Sprintf("--%s | -%s", function.Name, string(function.Name[0]))
}

func GetLineLength() int {
	lineLength := int((NewWindow().Width * 75) / 100)
	if lineLength > 79 {
		return 79
	}
	return lineLength
}

func WrapDescription(paragraph string) string {
	description := ""
	delimiter := " "
	cursor := 0
	lineLength := GetLineLength()
	for _, word := range strings.Split(paragraph, delimiter) {
		cursor = (cursor + len(word) + 1)
		description = fmt.Sprintf("%s%s%s", description, word, delimiter)
		if cursor >= lineLength {
			cursor = 0
			description = fmt.Sprintf("%s\n", description)
		}
	}
	return strings.TrimSuffix(description, "\n")
}

func WrapFunction(name string, functions []*Function) string {
	delimiter := " "
	paragraphs := [][]string{[]string{}}
	prefix := fmt.Sprintf("usage: %s", name)
	offset := len(prefix)
	cursor := 0
	lineLength := GetLineLength()
	for _, function := range functions {
		i := len(paragraphs) - 1
		option := fmt.Sprintf("[%s]", GetFunctionString(function))
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

func WrapUse(name string, description string, functions []*Function) string {
	return fmt.Sprintf("%s\n\n%s", WrapDescription(description), WrapFunction(name, functions))
}
