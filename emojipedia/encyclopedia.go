package emojipedia

import (
	text "github.com/gellel/emojipedia/emojipedia-text"
)

// Encyclopedia is a map of Emoji.
type Encyclopedia map[string]*Emoji

// Assign a new emoji.
func (encyclopedia *Encyclopedia) Assign(emoji *Emoji) (key string) {
	key = emoji.Name
	(*encyclopedia)[key] = emoji
	return key
}

// Each key iterates.
func (encyclopedia *Encyclopedia) Each(f func(key string, emoji *Emoji)) *Encyclopedia {
	for key, emoji := range *encyclopedia {
		f(key, emoji)
	}
	return encyclopedia
}

// Fetch emoji.
func (encyclopedia *Encyclopedia) Fetch(key string) (emoji *Emoji) {
	emoji, _ = encyclopedia.Get(key)
	return emoji
}

// Get emoji.
func (encyclopedia *Encyclopedia) Get(key string) (emoji *Emoji, ok bool) {
	emoji, ok = (*encyclopedia)[key]
	return emoji, ok
}

// Has key in map.
func (encyclopedia *Encyclopedia) Has(key string) (ok bool) {
	_, ok = (*encyclopedia)[key]
	return ok
}

// Keys in map.
func (encyclopedia *Encyclopedia) Keys() (keys Strings) {
	keys = Strings{}
	for key := range *encyclopedia {
		keys.Push(key)
	}
	return keys
}

// Len of map.
func (encyclopedia *Encyclopedia) Len() (number int) {
	number = len(*encyclopedia)
	return number
}

// New Emoji.
func (encyclopedia *Encyclopedia) New(value string) (key string) {
	key = text.Normalize(value)
	(*encyclopedia)[key] = &Emoji{}
	return key
}

// Remove emoji.
func (encyclopedia *Encyclopedia) Remove(key string) (ok bool) {
	delete(*encyclopedia, key)
	ok = (encyclopedia.Has(key) == false)
	return ok
}
