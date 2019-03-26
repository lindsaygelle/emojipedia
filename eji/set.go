package eji

type Set map[string][]*Emoji

func (set *Set) Add(key string, emoji *Emoji) {
	if s, ok := (*set)[key]; ok {
		(*set)[key] = append(s, emoji)
	} else {
		(*set)[key] = []*Emoji{emoji}
	}
}

func (set *Set) Has(key string) bool {
	_, ok := (*set)[key]
	return ok
}

func (set *Set) Remove(key string) bool {
	if _, ok := (*set)[key]; ok {
		delete(*set, key)
		return true
	}
	return false
}
