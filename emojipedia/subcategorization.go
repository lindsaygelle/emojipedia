package emojipedia

// A Subcategorization of Emoji.
type Subcategorization map[string]*Subcategory

// Assign a new subcategory.
func (subcategorization *Subcategorization) Assign(subcategory *Subcategory) (key string) {
	key = subcategory.Name
	(*subcategorization)[key] = subcategory
	return key
}

// Each key iterates.
func (subcategorization *Subcategorization) Each(f func(key string, subcategory *Subcategory)) *Subcategorization {
	for key, subcategory := range *subcategorization {
		f(key, subcategory)
	}
	return subcategorization
}

// Fetch subcategory.
func (subcategorization *Subcategorization) Fetch(key string) (subcategory *Subcategory) {
	subcategory, _ = subcategorization.Get(key)
	return subcategory
}

// Get subcategory.
func (subcategorization *Subcategorization) Get(key string) (subcategory *Subcategory, ok bool) {
	subcategory, ok = (*subcategorization)[key]
	return subcategory, ok
}

// Has key in map.
func (subcategorization *Subcategorization) Has(key string) (ok bool) {
	_, ok = (*subcategorization)[key]
	return ok
}

// Keys in map.
func (subcategorization *Subcategorization) Keys() (keys Strings) {
	keys = Strings{}
	for key := range *subcategorization {
		keys.Push(key)
	}
	return keys
}

// Len of map.
func (subcategorization *Subcategorization) Len() (number int) {
	number = len(*subcategorization)
	return number
}

// New subcategory.
func (subcategorization *Subcategorization) New(position int, value string) (key string) {
	key = Normalize(value)
	(*subcategorization)[key] = &Subcategory{}
	return key
}

// Remove subcategory.
func (subcategorization *Subcategorization) Remove(key string) (ok bool) {
	delete(*subcategorization, key)
	ok = (subcategorization.Has(key) == false)
	return ok
}
