package cache

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"net/http"
	"testing"
)

const document string = "cache.go"

func TestExists(t *testing.T) {
	outcome := exists(document)
	expects := true
	if outcome != expects {
		message := fmt.Sprintf("%s %t != %t", document, outcome, expects)
		t.Error(message)
	}
}

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
	outcome := writeHTML("test.html", response)
	if outcome != nil {
		message := fmt.Sprintf("%s error != nil", document)
		t.Error(message)
	}
}

func TestWriteJSON(t *testing.T) {
	content := map[string]interface{}{
		"a": "test",
		"b": []string{"a", "b"}}
	outcome := writeJSON("test.json", content)
	if outcome != nil {
		message := fmt.Sprintf("%s error != nil", document)
		t.Error(message)
	}
}

func TestGetHTML(t *testing.T) {
	_, outcome := getHTML("test.html")
	if outcome != nil {
		message := fmt.Sprintf("%s error != nil: reason %s", document, outcome)
		t.Error(message)
	}
}
func TestGetJSON(t *testing.T) {
	anonymous := struct {
		a string
		b []string
	}{}
	_, outcome := getJSON("test.json", &anonymous)
	if outcome != nil {
		message := fmt.Sprintf("%s error != nil: reason %s", document, outcome)
		t.Error(message)
	}
}

func TestRemoveHTML(t *testing.T) {
	outcome := remove("test.html")
	if outcome != nil {
		message := fmt.Sprintf("%s error != nil: reason %s", document, outcome)
		t.Error(message)
	}
}
func TestRemoveJSON(t *testing.T) {
	outcome := remove("test.json")
	if outcome != nil {
		message := fmt.Sprintf("%s error != nil: reason %s", document, outcome)
		t.Error(message)
	}
}
