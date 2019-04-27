package emojipedia

// Categorization is a map of category.
type Categorization map[string]*Category

// Assign a new category.
func (categorization *Categorization) Assign(category *Category) (key string) {
	key = category.Name
	(*categorization)[key] = category
	return key
}

// Each key iterates.
func (categorization *Categorization) Each(f func(key string, category *Category)) *Categorization {
	for key, category := range *categorization {
		f(key, category)
	}
	return categorization
}

// Fetch category.
func (categorization *Categorization) Fetch(key string) (category *Category) {
	category, _ = categorization.Get(key)
	return category
}

// Get category.
func (categorization *Categorization) Get(key string) (category *Category, ok bool) {
	category, ok = (*categorization)[key]
	return category, ok
}

// Has key in map.
func (categorization *Categorization) Has(key string) (ok bool) {
	_, ok = (*categorization)[key]
	return ok
}

// Keys in map.
func (categorization *Categorization) Keys() (keys Strings) {
	keys = Strings{}
	for key := range *categorization {
		keys.Push(key)
	}
	return keys
}

// Len of map.
func (categorization *Categorization) Len() (number int) {
	number = len(*categorization)
	return number
}

// New category.
func (categorization *Categorization) New(position int, value string) (key string) {
	key = Normalize(value)
	(*categorization)[key] = &Category{}
	return key
}

// Remove category.
func (categorization *Categorization) Remove(key string) (ok bool) {
	delete(*categorization, key)
	ok = (categorization.Has(key) == false)
	return ok
}
