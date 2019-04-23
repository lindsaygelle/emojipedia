package emojipedia

type Numeralization map[int]string

// Each key iterates.
func (numeralization *Numeralization) Each(f func(key int, value string)) *Numeralization {
	for key, value := range *numeralization {
		f(key, value)
	}
	return numeralization
}

// Fetch value.
func (numeralization *Numeralization) Fetch(key int) (value string) {
	value, _ = numeralization.Get(key)
	return value
}

// Get value.
func (numeralization *Numeralization) Get(key int) (value string, ok bool) {
	value, ok = (*numeralization)[key]
	return value, ok
}

// Has key in map.
func (numeralization *Numeralization) Has(key int) (ok bool) {
	_, ok = (*numeralization)[key]
	return ok
}

// Keys in map.
func (numeralization *Numeralization) Keys() (keys []int) {
	keys = []int{}
	for key := range *numeralization {
		keys = append(keys, key)
	}
	return keys
}

// New value.
func (numeralization *Numeralization) New(value string) (key int) {
	key = len(*numeralization)
	(*numeralization)[key] = value
	return key
}

// Remove value.
func (numeralization *Numeralization) Remove(key int) {
	delete(*numeralization, key)
}
