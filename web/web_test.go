package web_test

import (
	"fmt"
	"testing"

	"github.com/gellel/emojipedia/web"
)

func TestGetHTML(t *testing.T) {
	_, outcome := web.GetHTML()
	if outcome != nil {
		message := fmt.Sprintf("web.go outcome != nil: reason %s", outcome)
		t.Error(message)
	}
}
