package u

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	"github.com/gellel/emojipedia/manifest"
)

const Filename string = "unicode-emoji.html"

const root string = "emojipedia"

const url string = "https://unicode.org/emoji/charts/emoji-list.html"

var Export = Unicode

var Key = "UNICODE"

var Options = []interface{}{Cached, Get, Remove}

var empty = map[string]func(){
	"CACHED": Cached,
	"GET":    Get,
	"REMOVE": Remove}

func Unicode(options ...string) {}

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
		fmt.Println(fmt.Sprintf("%s is not currently cached on disk. without this file, emojipedia dependencies cannot be built", Filename))
	} else {
		fmt.Println(fmt.Sprintf("%s is currently stored at %s. file size %v mb. if the associated dependencies have been built, this file can be removed.", Filename, dir, (int)(info.Size()/1024)/1024))
	}
}

func Get() {
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
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
	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(filepath.Join(dir, Filename), dump, 0644)
	if err != nil {
		panic(err)
	}
}

func Remove() {

}

func Main(m *manifest.Manifest, previous, options []string) {
	if len(options) != 0 {
		key := strings.ToUpper(strings.Replace(options[0], "-", "", -1))
		if f, ok := empty[key]; ok {
			f()
		}
	}
}
