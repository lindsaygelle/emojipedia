package emojipedia

type Set map[int]string

func (set *Set) Add(key int, name string) string {
	(*set)[key] = name
	return name
}

func (set *Set) Has(key int) (ok bool) {
	_, ok = (*set)[key]
	return ok
}

func (set *Set) Get(key int) (name string, ok bool) {
	name, ok = (*set)[key]
	return name, ok
}

func (set *Set) GetUnsafely(key int) (name string) {
	name, _ = (*set)[key]
	return name
}

func (set *Set) Length() (length int) {
	return len(*set)
}

func (set *Set) Remove(key int) (ok bool) {
	delete(*set, key)
	if set.Has(key) != true {
		return true
	}
	return false
}
