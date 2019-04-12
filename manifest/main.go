package manifest

import (
	"encoding/json"
	"io/ioutil"
)

func NewManifest(filepath string) *Manifest {
	content, err := ioutil.ReadFile(filepath)
	if err != nil {
		panic(err)
	}
	manifest := &Manifest{}
	err = json.Unmarshal(content, manifest)
	if err != nil {
		panic(err)
	}
	return manifest
}
