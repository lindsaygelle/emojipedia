package emojipedia

type Map map[string]*Emoji

func (m *Map) Add(key string, emoji *Emoji) *Map {
	(*m)[key] = emoji
	return m
}

func (m *Map) Has(key string) bool {
	_, ok := (*m)[key]
	return ok
}

func (m *Map) Remove(key string) bool {
	_, ok := (*m)[key]
	if ok {
		delete(*m, key)
	}
	return ok
}
