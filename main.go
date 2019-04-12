package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	b "github.com/gellel/emojipedia/build"
	"github.com/gellel/emojipedia/cli"
	e "github.com/gellel/emojipedia/emoji"
	f "github.com/gellel/emojipedia/files"
)

type Manifest struct {
	Author      string `json:"author"`
	Description string `json:"description"`
	Programs    map[string]Program
	Name        string  `json:"name"`
	Version     float64 `json:"version"`
}

type Program struct {
	Description string `json:"description"`
	Programs    map[string]Program
	Name        string `json:"name"`
}

var (
	filename      = "manifest.json"
	_, file, _, _ = runtime.Caller(0)
	dir           = filepath.Dir(file)
	manifest      = &Manifest{}
	subcalls      = []string{}
)

var subprogram Program

func call(options []string, functions []*cli.Function) {
	for _, function := range functions {
		argument := strings.ToLower(options[0])
		if argument == strings.ToLower(function.Name) {
			function.F.(func(...string))(options[1:]...)
			break
		} else if argument == fmt.Sprintf("--%s", function.Name) ||
			argument == fmt.Sprintf("-%s", string(function.Name[0])) {
			function.F.(func())()
			break
		}
	}
}

func author() {
	s := "%s written by %s. data scraped from unicode.org & emojipedia.org. feel free to contribute or improve."
	fmt.Println(cli.WrapDescription(fmt.Sprintf(s, manifest.Name, manifest.Author)))
}

func version() {
	v := strconv.FormatFloat(manifest.Version, 'g', 1, 64)
	fmt.Println(fmt.Sprintf("%s version %s", manifest.Name, v))
}

func categories(options ...string) {}

func emoji(options ...string) {
	key := "emoji"
	subprogram = subprogram.Programs[key]
	subcalls = append(subcalls, key)
	name := strings.Join(subcalls, " ")
	description := subprogram.Description
	program := cli.NewProgram(name, description, e.Options)
	if len(options) == 0 {
		fmt.Println(program.Use)
	} else {
		e.Main(options[0:])
	}
}

func subcategories(options ...string) {}

func build(options ...string) {
	key := "build"
	subprogram = manifest.Programs[key]
	subcalls = append(subcalls, key)
	name := strings.Join(subcalls, " ")
	description := subprogram.Description
	functions := b.Options
	program := cli.NewProgram(name, description, functions)
	if len(options) == 0 {
		fmt.Println(program.Use)
	} else {
		f.Main(options[0:])
	}
}

func files(options ...string) {
	key := "files"
	subprogram = manifest.Programs[key]
	subcalls = append(subcalls, key)
	name := strings.Join(subcalls, " ")
	description := subprogram.Description
	functions := f.Options
	program := cli.NewProgram(name, description, functions)
	if len(options) == 0 {
		fmt.Println(program.Use)
	} else {
		f.Main(options[0:])
	}
}

func get(options ...string) {
	key := "get"
	subprogram = manifest.Programs[key]
	subcalls = append(subcalls, key)
	name := strings.Join(subcalls, " ")
	description := subprogram.Description
	functions := []interface{}{emoji}
	program := cli.NewProgram(name, description, functions)
	if len(options) == 0 {
		fmt.Println(program.Use)
	} else {
		call(options[0:], program.Functions)
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
	subcalls = append(subcalls, name)
	functions := []interface{}{author, version, build, files, get}
	program := cli.NewProgram(name, description, functions)
	if len(os.Args) == 1 {
		fmt.Println(program.Use)
	} else {
		call(os.Args[1:], program.Functions)
	}
}
