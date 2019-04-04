package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"reflect"
	"runtime"
	"strings"
)

func getProgramUsage(description string, choices []interface{}) (string, string) {
	_, file, _, _ := runtime.Caller(0)
	programFunctions := []string{}
	for _, program := range choices {
		memoryPointer := reflect.ValueOf(program).Pointer()
		functionNamespace := runtime.FuncForPC(memoryPointer).Name()
		programName := strings.Replace(functionNamespace, file, "", 1)
		templateString := fmt.Sprintf("[%s..args]", strings.ToUpper(programName))
		programFunctions = append(programFunctions, templateString)
	}
	templateString := strings.Join(programFunctions, ",")
	commandChoices := fmt.Sprintf("[%s]", templateString)
	return format(description), commandChoices
}

func about(manifest *Manifest) {
	message := fmt.Sprintf("%s: %s\n", manifest.Name, manifest.Description)
	fmt.Println(fmt.Sprintf(format(message), "GET"))
	options := namespaces([]interface{}{get, search})
	functions := []string{}
	for _, opt := range options {
		name := "[" + strings.ToUpper(opt) + "...args" + "]"
		functions = append(functions, name)
	}
	choices := strings.Join(functions, ",")
	usage := fmt.Sprintf("usage: %s [%s]\n", manifest.Name, choices)
	fmt.Println(usage)
	os.Exit(0)
}

func get(manifest *Manifest) {
}

func search(manfiest *Manifest) {}

func main() {
	_, file, _, _ := runtime.Caller(0)
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
	switch len(os.Args) {
	case 1:
		d, t := getProgramUsage(manifest.Description, []interface{}{get, search})
		fmt.Println(d, t)
		//about(manifest)
	default:
		switch strings.ToUpper(os.Args[1]) {
		case "GET":
			get(manifest)
		case "SEARCH":
			search(manifest)
		default:
			os.Exit(2)
		}
	}
}

func namespaces(fn []interface{}) []string {
	names := []string{}
	for _, f := range fn {
		value := reflect.ValueOf(f)
		address := value.Pointer()
		name := runtime.FuncForPC(address).Name()
		name = strings.Replace(name, "main.", "", 1)
		names = append(names, name)
	}
	return names
}

func format(text string) (paragraph string) {
	maxLength := 79
	currentLength := 0
	substrings := strings.Split(text, " ")
	for _, substring := range substrings {
		paragraph = paragraph + substring + " "
		currentLength = currentLength + len(substring)
		if currentLength > maxLength {
			paragraph = paragraph + "\n"
			currentLength = 0
		}
	}
	return paragraph
}
