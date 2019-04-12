package c

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/PuerkitoBio/goquery"
	u "github.com/gellel/emojipedia/emojipedia-web/web-unicode"
	"github.com/gellel/emojipedia/manifest"
)

const Filename string = "categories.json"

const root string = "emojipedia"

var Export = Categories

var Key = "CATEGORIES"

var Options = []interface{}{Cached, Make, Remove}

var empty = map[string](func()){
	"CACHED": Cached,
	"MAKE":   Make,
	"REMOVE": Remove}

var replacements = []string{" ", "-", "&", "and"}

var replacer = strings.NewReplacer(replacements...)

func Categories(options ...string) {}

func Cached() {
	_, file, _, _ := runtime.Caller(0)
	dir := filepath.Dir(file)
	for {
		base := filepath.Base(dir)
		if base == root {
			break
		}
		dir = filepath.Dir(dir)
	}
	info, err := os.Stat(filepath.Join(dir, Filename))
	if err != nil {
		fmt.Printf("%s is not stored on disk. checked directory %s.", Filename, dir)
	} else {
		fmt.Printf("%s is stored at %s. file size %v mb.", Filename, dir, (int)((info.Size()/1024)/1024))
	}
}

func Make() {
	_, file, _, _ := runtime.Caller(0)
	dir := filepath.Dir(file)
	for {
		base := filepath.Base(dir)
		if base == root {
			break
		}
		dir = filepath.Dir(dir)
	}
	path := filepath.Join(dir, u.Filename)
	_, err := os.Stat(path)
	if err != nil {
		panic(err)
	}
	f, err := os.Open(path)
	if err != nil {
		panic(err)
	}
	defer f.Close()
	doc, err := goquery.NewDocumentFromReader(f)
	if err != nil {
		panic(err)
	}
	categories := []string{}
	doc.Find("tr").Each(func(_ int, selection *goquery.Selection) {
		selection.Find("th.bighead").Each(func(_ int, selection *goquery.Selection) {
			categories = append(categories, strings.TrimSpace(selection.Text()))
		})
	})
	m := make(map[int]string, len(categories))
	for i, c := range categories {
		m[i] = strings.ToLower(replacer.Replace(c))
	}
	contents, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(Filename, contents, 0644)
	if err != nil {
		panic(err)
	}
}

func Open() (map[int]string, error) {
	_, file, _, _ := runtime.Caller(0)
	dir := filepath.Dir(file)
	for {
		base := filepath.Base(dir)
		if base == root {
			break
		}
		dir = filepath.Dir(dir)
	}
	path := filepath.Join(dir, Filename)
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	b, err := ioutil.ReadAll(f)
	if err != nil {
		return nil, err
	}
	m := make(map[int]string)
	err = json.Unmarshal(b, &m)
	if err != nil {
		return nil, err
	}
	return m, nil
}

func Remove() {
	_, file, _, _ := runtime.Caller(0)
	dir := filepath.Dir(file)
	for {
		base := filepath.Base(dir)
		if base == root {
			break
		}
		dir = filepath.Dir(dir)
	}
	err := os.Remove(Filename)
	if err != nil {
		fmt.Println(fmt.Sprintf("cannot remove %s. file is not on disk.", Filename))
	} else {
		fmt.Println(fmt.Sprintf("successfully removed %s.", Filename))
	}
}

func Main(m *manifest.Manifest, previous, options []string) {
	if len(options) != 0 {
		key := strings.ToUpper(strings.Replace(options[0], "-", "", -1))
		if f, ok := empty[key]; ok {
			f()
		}
	}
}
