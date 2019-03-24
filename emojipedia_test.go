package emojipedia_test

import (
	"testing"

	"github.com/gellel/emojipedia"
)

func Test(t *testing.T) {
	emojipedia.Get("https://emojipedia.org")
}
