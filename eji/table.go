package eji

type Table map[string]*Emoji

func (table *Table) Add(key string, emoji *Emoji) *Table {
	(*table)[key] = emoji
	return table
}

func (table *Table) Has(key string) bool {
	_, ok := (*table)[key]
	return ok
}

func (table *Table) Remove(key string) bool {
	_, ok := (*table)[key]
	if ok {
		delete(*table, key)
	}
	return ok
}
