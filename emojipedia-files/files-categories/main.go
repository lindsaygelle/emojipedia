package categories

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gellel/emojipedia/emojipedia"
	"github.com/gellel/emojipedia/manifest"

	unicodeorg "github.com/gellel/emojipedia/emojipedia-web/web-unicode"
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
		fmt.Println(fmt.Sprintf("%s is not stored on disk. checked directory %s.", Filename, dir))
	} else {
		fmt.Println(fmt.Sprintf("%s is stored at %s. file size %v kb.", Filename, dir, (int)((info.Size() / 1024))))
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
	path := filepath.Join(dir, unicodeorg.Filename)
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
	contents, err := json.Marshal(emojipedia.NewCategories(doc))
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
