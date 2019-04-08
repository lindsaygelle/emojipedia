package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/gellel/emojipedia/cli"
	"github.com/gellel/emojipedia/emoji"
)

type Manifest struct {
	Description string `json:"description"`
	Programs    map[string]Program
	Name        string `json:"name"`
}

type Program struct {
	Description string `json:"description"`
	Name        string `json:"name"`
}

var (
	filename      = "manifest.json"
	_, file, _, _ = runtime.Caller(0)
	dir           = filepath.Dir(file)
	manifest      = &Manifest{}
)

func call(options []string, functions []*cli.Function) {
	for _, function := range functions {
		if strings.ToLower(options[0]) == strings.ToLower(function.Name) {
			function.F.(func(...string))(options[1:]...)
		}
	}
}

func get(options ...string) {
	content := manifest.Programs["get"]
	description := content.Description
	program := cli.NewProgram(manifest.Name+" "+"get", description, emoji.Options)
	if len(options) == 0 {
		fmt.Println(program.Use)
	} else {
		call(options[1:], program.Functions)
	}
}

func main() {
	directory := path.Dir(file)
	fullpath := filepath.Join(directory, filename)
	content, err := ioutil.ReadFile(fullpath)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(content, manifest)
	if err != nil {
		panic(err)
	}
	name := manifest.Name
	description := manifest.Description
	functions := []interface{}{get}
	program := cli.NewProgram(name, description, functions)
	if len(os.Args) == 1 {
		fmt.Println(program.Use)
	} else {
		call(os.Args[1:], program.Functions)
	}
}
