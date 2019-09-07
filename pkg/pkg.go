package pkg

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"os"
	"path/filepath"
	"runtime"

	"github.com/PuerkitoBio/goquery"
)

const (
	URL = "http://www.unicode.org/emoji/charts/emoji-list.html"
)

const (
	dir    string = "emojipedia"
	folder string = "unicode"
)

var (
	_, b, _, _  = runtime.Caller(0)
	root        = filepath.Dir(filepath.Dir(b))
	storagepath = filepath.Join(root, fmt.Sprintf(".%s", dir), folder)
)

func HTTP() (*http.Response, error) {
	resp, err := http.Get(URL)
	if err != nil {
		return nil, err
	}
	if resp.StatusCode != 200 {
		return nil, fmt.Errorf(resp.Status)
	}
	return resp, nil
}

// Open attempts to open the unicode-org HTTP response from the emojipedia/unicode folder.
func Open() (*goquery.Document, error) {
	filepath := filepath.Join(storagepath, "unicode.html")
	reader, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	document, err := goquery.NewDocumentFromReader(reader)
	if err != nil {
		return nil, err
	}
	defer reader.Close()
	return document, nil
}

// Write stores and unicode-org HTTP response to the dependencies folder.
func Write(resp *http.Response) error {
	err := os.MkdirAll(storagepath,  os.ModePerm)
	if err != nil {
		return err
	}
	dump, err := httputil.DumpResponse(resp, true)
	if err != nil {
		return err
	}
	filepath := filepath.Join(storagepath, "unicode.html")
	return ioutil.WriteFile(filepath, dump,  os.ModePerm)
}

// Remove deletes the unicode-org data stored in the dependencies folder.
func Remove() error {
	return os.Remove(filepath.Join(storagepath, "unicode.html"))
}
