package store

import (
	"encoding/json"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/gellel/emojipedia/eji"
)

func directory() (path string, err error) {
	current, err := os.Getwd()
	if err != nil {
		return path, err
	}
	path = filepath.Dir(current)
	return path, nil
}

func Set(name string, set *eji.Set) error {
	c, err := json.Marshal(set)
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
