package e

import (
	"net/http"

	"github.com/gellel/emojipedia/manifest"
)

var Export = Emojipedia

var Key = "EMOJIPEDIA"

func Emojipedia(options ...string) {}

func Get(name string) {
	url := "https://emojipedia.org/" + name
	resp, err := http.Get(url)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	if resp.StatusCode != 200 {
		panic(err)
	}
}

func Main(m *manifest.Manifest, previous, options []string) {}
