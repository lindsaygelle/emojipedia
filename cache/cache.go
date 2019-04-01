package cache

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"os"
	"path/filepath"

	"github.com/PuerkitoBio/goquery"
)

func exists(name string) bool {
	_, err := os.Stat(getPath(name))
	return !os.IsNotExist(err)
}

func getHTML(name string) (*goquery.Document, error) {
	file, err := os.Open(getPath(name))
	if err != nil {
		return nil, err
	}
	doc, err := goquery.NewDocumentFromReader(file)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return doc, nil
}

func getJSON(name string, i interface{}) (interface{}, error) {
	file, err := os.Open(getPath(name))
	if err != nil {
		panic(err)
	}
	defer file.Close()
	dump, err := ioutil.ReadAll(file)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(dump, i)
	return i, err
}

func getPath(name string) string {
	directory, err := os.Getwd()
	if err != nil {
		panic(err)
	}
	return filepath.Join(directory, name)
}

func remove(name string) error {
	err := os.Remove(getPath(name))
	if err != nil {
		return err
	}
	return nil
}

func write(dump []byte, name string) error {
	return ioutil.WriteFile(getPath(name), dump, 0644)
}

func writeHTML(name string, response *http.Response) error {
	dump, err := httputil.DumpResponse(response, true)
	if err != nil {
		return err
	}
	return write(dump, name)
}

func writeJSON(name string, i interface{}) error {
	dump, err := json.Marshal(i)
	if err != nil {
		return err
	}
	return write(dump, name)
}
