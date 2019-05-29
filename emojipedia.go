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
		get(arguments.Next().Get(0), emoji.Read, emoji.Detail)
	case K, KEYS:
		keys(EMOJIPEDIA, emojipedia.List)
	case L, LIST:
		list(EMOJIPEDIA, emojipedia.List, emoji.List)
	case R, REMOVE:
		remove(EMOJIPEDIA, emojipedia.Remove)
	}
}
