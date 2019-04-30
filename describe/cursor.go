package describe

type Cursor int

func (cursor Cursor) Add(number ...int) Cursor {
	cursor = (cursor + Cursor(cursor.Sum(number...)))
	return cursor
}

func (cursor Cursor) Equal(number int) (ok bool) {
	ok = (cursor == Cursor(number))
	return ok
}

func (cursor Cursor) GreaterThan(number int) (ok bool) {
	ok = (cursor > Cursor(number))
	return ok
}

func (cursor Cursor) GreaterThanEqualTo(number int) (ok bool) {
	ok = (cursor >= Cursor(number))
	return ok
}

func (cursor Cursor) Int(number int) int {
	return int(cursor)
}

func (cursor Cursor) LessThan(number int) (ok bool) {
	ok = (cursor < Cursor(number))
	return ok
}

func (cursor Cursor) LessThanEqualTo(number int) (ok bool) {
	ok = (cursor <= Cursor(number))
	return ok
}

func (cursor Cursor) Set(number int) Cursor {
	cursor = Cursor(number)
	return cursor
}

func (cursor Cursor) Subtract(number ...int) Cursor {
	cursor = (cursor - Cursor(cursor.Sum(number...)))
	return cursor
}

func (cursor Cursor) Sum(numbers ...int) (sum int) {
	for _, number := range numbers {
		sum = (sum + number)
	}
	return sum
}
