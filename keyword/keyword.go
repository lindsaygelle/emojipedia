package keyword

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/gellel/emojipedia/directory"
	"github.com/gellel/emojipedia/slice"
)

// Open attempts to open a Keyword slice from the emojipedia/keywords folder.
func Open(name string) (*slice.Slice, error) {
	filepath := filepath.Join(directory.Keywords, fmt.Sprintf("%s.json", name))
	reader, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadAll(reader)
	defer reader.Close()
	if err != nil {
		return nil, err
	}
	slice := &slice.Slice{}
	err = json.Unmarshal(content, slice)
	if err != nil {
		return nil, err
	}
	return slice, nil
}

func Parse(content *[]byte) (*slice.Slice, error) {
	keywords := &slice.Slice{}
	err := json.Unmarshal(*content, keywords)
	if err != nil {
		return nil, err
	}
	return keywords, nil
}

func Read(name string) (*[]byte, error) {
	filepath := filepath.Join(directory.Keywords, fmt.Sprintf("%s.json", name))
	reader, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadAll(reader)
	defer reader.Close()
	if err != nil {
		return nil, err
	}
	return &content, nil
}

// Remove deletes the Keyword data stored in the dependencies folder.
func Remove(name string) error {
	return os.Remove(filepath.Join(directory.Keywords, fmt.Sprintf("%s.json", name)))
}

// Write stores and Keyword entry to the dependencies folder.
func Write(key string, keywords *slice.Slice) error {
	err := os.MkdirAll(directory.Keywords, 0644)
	if err != nil {
		return err
	}
	content, err := json.Marshal(keywords)
	if err != nil {
		return err
	}
	filepath := filepath.Join(directory.Keywords, fmt.Sprintf("%s.json", key))
	return ioutil.WriteFile(filepath, content, 0644)
}
