package emojipedia

type Cursor int

// Decrease cursor
func (cursor *Cursor) Decrease() *Cursor {
	*cursor = *cursor - 1
	return cursor
}

// Increase cursor.
func (cursor *Cursor) Increase() *Cursor {
	*cursor = *cursor + 1
	return cursor
}

// Value from cursor.
func (cursor *Cursor) Value() (value int) {
	value = int(*cursor)
	return value
}
