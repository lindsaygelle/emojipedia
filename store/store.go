package store

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"
)

func directory() (path string, err error) {
	current, err := os.Getwd()
	if err != nil {
		return path, err
	}
	path = filepath.Dir(current)
	return path, nil
}

func Exists(name string) bool {
	dir, err := directory()
	if err != nil {
		return false
	}
	if _, err := os.Stat(filepath.Join(dir, name)); os.IsNotExist(err) {
		return false
	}
	return true
}

func Store(name string, i interface{}) error {
	c, err := json.Marshal(i)
	if err != nil {
		return err
	}
	dir, err := directory()
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(filepath.Join(dir, name), c, 0644)
	if err != nil {
		return err
	}
	return nil
}
