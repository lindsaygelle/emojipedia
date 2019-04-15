package org

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"path/filepath"
	"regexp"
	"runtime"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gellel/emojipedia/manifest"

	emojis "github.com/gellel/emojipedia/emojipedia-files/files-emojis"
)

const root string = "emojipedia"

const url string = "https://emojipedia.org/"

var Export = Emojipedia

var Key = "EMOJIPEDIA"

var Options = []interface{}{Get}

var name = map[string](func(name string)){
	"GET": Get}

func Emojipedia(options ...string) {}

func Get(name string) {
	resp, err := http.Get(url + name)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		panic(err)
	}
	doc, err := goquery.NewDocumentFromResponse(resp)
	if err != nil {
		panic(err)
	}
	description := []string{}
	doc.Find("section.description").First().Each(func(i int, selection *goquery.Selection) {
		selection.Find("p").Each(func(j int, selection *goquery.Selection) {
			description = append(description, strings.TrimSpace(selection.Text()))
		})
	})
	if len(description) == 0 {
		panic(fmt.Errorf("%s description not found", name))
	}
	m, err := emojis.Open()
	if err != nil {
		panic(err)
	}
	e, ok := m[name]
	if ok != true {
		panic(fmt.Errorf("%s is missing from json", name))
	}
	e.Description = regexp.MustCompile(`\r?\n`).ReplaceAllString(strings.Join(description, " "), " ")
	m[name] = e
	dump, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	_, file, _, _ := runtime.Caller(0)
	dir := filepath.Dir(file)
	for {
		base := filepath.Base(dir)
		if base == root {
			break
		}
		dir = filepath.Dir(dir)
	}
	err = ioutil.WriteFile(filepath.Join(dir, emojis.Filename), dump, 0644)
	if err != nil {
		panic(err)
	}

}

func Main(m *manifest.Manifest, previous, options []string) {
	if len(options) != 0 {
		key := strings.ToUpper(strings.Replace(options[0], "-", "", -1))
		if f, ok := name[key]; ok {
			f(options[1])
		}
	}
}
