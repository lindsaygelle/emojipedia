package emojipedia

type Emojidex map[string]*Emoji

func (emojidex *Emojidex) Add(key string, emoji *Emoji) {
	(*emojidex)[key] = emoji
}

func (emojidex *Emojidex) Get(key string) (emoji *Emoji, ok bool) {
	emoji, ok = (*emojidex)[key]
	return emoji, ok
}

func (emojidex *Emojidex) GetUnsafely(key string) (emoji *Emoji) {
	emoji, _ = (*emojidex)[key]
	return emoji
}

func (emojidex *Emojidex) Has(key string) (ok bool) {
	_, ok = (*emojidex)[key]
	return ok
}

func (emojidex *Emojidex) Length() (length int) {
	return len(*emojidex)
}

func (emojidex *Emojidex) Remove(key string) (ok bool) {
	delete(*emojidex, key)
	if ok = emojidex.Has(key); ok != true {
		return true
	}
	return false
}
