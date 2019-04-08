package emojipedia

type Set map[string][]string

func (set *Set) Add(key string, name string) *Set {
	if s, ok := (*set)[key]; ok {
		(*set)[key] = append(s, name)
	} else {
		(*set)[key] = []string{name}
	}
	return set
}

func (set *Set) Get(key string) []string {
	if emojis, ok := (*set)[key]; ok {
		return emojis
	}
	return []string{}
}

func (set *Set) Has(key string) bool {
	_, ok := (*set)[key]
	return ok
}

func (set *Set) Remove(key string) bool {
	_, ok := (*set)[key]
	if ok {
		delete(*set, key)
	}
	return ok
}
