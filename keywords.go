package main

import (
	"strings"

	"github.com/gellel/emojipedia/arguments"
	"github.com/gellel/emojipedia/keywords"
)

func keywordsMain(arguments *arguments.Arguments) {
	switch strings.ToUpper(arguments.Get(0)) {
	case B, BUILD:
		build(KEYWORDS, keywords.Make)
	}
}
