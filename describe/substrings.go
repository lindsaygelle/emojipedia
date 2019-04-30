package describe

type Substrings []String

func (substrings Substrings) Bounds(i int) (ok bool) {
	ok = ((i > -1) && (i < substrings.Len()))
	return ok
}

func (substrings Substrings) Each(f func(i int, str String)) Substrings {
	for i, str := range substrings {
		f(i, str)
	}
	return substrings
}

func (substrings Substrings) Len() (length int) {
	length = len(substrings)
	return length
}

func (substrings Substrings) Peek(i int) (str String) {
	if ok := substrings.Bounds(i); ok {
		str = substrings[i]
	}
	return str
}

func (substrings Substrings) Remove(i int) Substrings {
	substrings = append(substrings[:i], substrings[(i+1):]...)
	return substrings
}
