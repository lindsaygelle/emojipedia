package cache_test

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"

	"github.com/gellel/emojipedia/cache"
)

func TestWriteHTML(t *testing.T) {
	body := "Hello world"
	response := &http.Response{
		Status:        "200 OK",
		StatusCode:    200,
		Proto:         "HTTP/1.1",
		ProtoMajor:    1,
		ProtoMinor:    1,
		Body:          ioutil.NopCloser(bytes.NewBufferString(body)),
		ContentLength: int64(len(body)),
		Header:        make(http.Header, 0),
	}
	err := cache.WriteHTML("test.html", response)
	if err != nil {
		message := fmt.Sprintf("cache.go error != nil")
		t.Error(message)
	}
}

func TestWriteJSON(t *testing.T) {
	content := map[string]interface{}{
		"a": "test",
		"b": []string{"a", "b"}}
	err := cache.WriteJSON("test.json", content)
	if err != nil {
		message := fmt.Sprintf("cache.go error != nil: reason %s", err)
		t.Error(message)
	}
}

func TestGetHTML(t *testing.T) {
	_, err := cache.GetHTML("test.html")
	if err != nil {
		message := fmt.Sprintf("cache.go error != nil: reason %s", err)
		t.Error(message)
	}
}
func TestGetJSON(t *testing.T) {
	anonymous := struct {
		a string
		b []string
	}{}
	_, err := cache.GetJSON("test.json", &anonymous)
	if err != nil {
		message := fmt.Sprintf("cache.go error != nil: reason %s", err)
		t.Error(message)
	}
}
func TestRemoveHTML(t *testing.T) {
	err := cache.Remove("test.html")
	if err != nil {
		message := fmt.Sprintf("cache.go error != nil: reason %s", err)
		t.Error(message)
	}
}
func TestRemoveJSON(t *testing.T) {
	err := cache.Remove("test.json")
	if err != nil {
		message := fmt.Sprintf("cache.go error != nil: reason %s", err)
		t.Error(message)
	}
}
