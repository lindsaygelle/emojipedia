package main

import (
	"strings"

	"github.com/gellel/emojipedia/arguments"
	"github.com/gellel/emojipedia/emoji"
	"github.com/gellel/emojipedia/emojipedia"
)

func emojipediaMain(arguments *arguments.Arguments) {
	switch strings.ToUpper(arguments.Get(0)) {
	case B, BUILD:
		build(EMOJIPEDIA, emojipedia.Make)
	case G, GET:
	case K, KEYS:
	case L, LIST:
		list(EMOJIPEDIA, emojipedia.List, emoji.List)
	}
}
