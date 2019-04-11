package emojipedia

type Set map[int]string

func (set *Set) Add(key int, name string) string {
	(*set)[key] = name
	return name
}

func (set *Set) Has(key int) bool {
	_, ok := (*set)[key]
	return ok
}

func (set *Set) Get(key int) string {
	name, _ := (*set)[key]
	return name
}

func (set *Set) Remove(key int) bool {
	delete(*set, key)
	if set.Has(key) != true {
		return true
	}
	return false
}
