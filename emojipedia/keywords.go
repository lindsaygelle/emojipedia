package emojipedia

import (
	text "github.com/gellel/emojipedia/emojipedia-text"
)

// Keywords is a map of string slices.
type Keywords map[string]*Strings

// Append assigns or makes and assigns a new string.
func (keywords *Keywords) Append(key string, value string) *Keywords {
	strings, ok := (*keywords)[key]
	if ok != true {
		(*keywords)[key] = &Strings{value}
	} else {
		strings.Push(value)
	}
	return keywords
}

// Assign a new strings.
func (keywords *Keywords) Assign(key string, strings *Strings) string {
	(*keywords)[key] = strings
	return key
}

// Each key iterates.
func (keywords *Keywords) Each(f func(key string, strings *Strings)) *Keywords {
	for key, strings := range *keywords {
		f(key, strings)
	}
	return keywords
}

// Fetch strings.
func (keywords *Keywords) Fetch(key string) (strings *Strings) {
	strings, _ = keywords.Get(key)
	return strings
}

// Get strings.
func (keywords *Keywords) Get(key string) (strings *Strings, ok bool) {
	strings, ok = (*keywords)[key]
	return strings, ok
}

// Has key in map.
func (keywords *Keywords) Has(key string) (ok bool) {
	_, ok = (*keywords)[key]
	return ok
}

// Keys in map.
func (keywords *Keywords) Keys() (keys *Strings) {
	keys = &Strings{}
	for key := range *keywords {
		keys.Push(key)
	}
	return keys
}

// Len of map.
func (keywords *Keywords) Len() (number int) {
	number = len(*keywords)
	return number
}

// New strings.
func (keywords *Keywords) New(value string) (key string) {
	key = text.Normalize(value)
	(*keywords)[key] = &Strings{}
	return key
}

// Remove strings.
func (keywords *Keywords) Remove(key string) (ok bool) {
	delete(*keywords, key)
	ok = (keywords.Has(key) == false)
	return ok
}
