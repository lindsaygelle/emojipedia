package store

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
)

func Dirs(name string) (path string, err error) {
	current, err := os.Getwd()
	if err != nil {
		return path, err
	}
	parent := filepath.Dir(current)
	_, err = os.Stat(parent)
	if os.IsNotExist(err) {
		return path, err
	}
	name = strings.Replace(name, ".json", "", -1)
	path = filepath.Join(parent, name)
	return path, nil
}

func Has(files ...string) error {
	for _, file := range files {
		path, err := Dirs(file)
		if err != nil {
			return err
		}
		path = strings.Join([]string{path, "json"}, ".")
		if _, err = os.Stat(path); os.IsNotExist(err) {
			return err
		}
	}
	return nil
}

func Open(name string) ([]byte, error) {
	path, err := Dirs(name)
	if err != nil {
		return nil, err
	}
	path = strings.Join([]string{path, "json"}, ".")
	file, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer file.Close()
	return ioutil.ReadAll(file)
}

func Save(name string, i interface{}) (path string, err error) {
	path, err = Dirs(name)
	if err != nil {
		return path, err
	}
	content, err := json.Marshal(i)
	if err != nil {
		return path, err
	}
	path = strings.Join([]string{path, "json"}, ".")
	return path, ioutil.WriteFile(path, content, 0644)
}
