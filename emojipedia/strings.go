package emojipedia

import (
	"sort"
	"strings"
)

type Strings []string

func (slice *Strings) After(i int) (ok bool) {
	if ok = slice.Bounds(i); ok {
		*slice = (*slice)[i:]
	}
	return ok
}

func (slice *Strings) Before(i int) (ok bool) {
	if ok = slice.Bounds(i); ok {
		*slice = (*slice)[:i]
	}
	return ok
}

func (slice *Strings) Bounds(i int) (ok bool) {
	ok = ((i > -1) && (i < len(*slice)))
	return ok
}

func (slice *Strings) Each(f func(i int, value string)) *Strings {
	for i, value := range *slice {
		f(i, value)
	}
	return slice
}

func (slice *Strings) Get(i int) (value string, ok bool) {
	if ok = slice.Bounds(i); ok {
		value = (*slice)[i]
	}
	return value, ok
}

func (slice *Strings) Has(value string) (position int, ok bool) {
	for i, str := range *slice {
		if ok = (value == str); ok {
			return i, ok
		}
	}
	return -1, ok
}

func (slice *Strings) Join(s ...string) (value string) {
	separate := strings.Join(s, "")
	slice.Each(func(i int, s string) {
		value = (value + separate + s)
	})
	return value
}

func (slice *Strings) Last() (value string) {
	value = slice.Peek(slice.Len() - 1)
	return value
}

func (slice *Strings) Len() (length int) {
	length = len(*slice)
	return length
}

func (slice *Strings) Less(i, j int) (ok bool) {
	s := *slice
	a, b := s[i], s[j]
	if ok = (a == b); ok {
		a, b = strings.ToLower(a), strings.ToLower(b)
	}
	ok = a < b
	return ok
}

func (slice *Strings) Push(value string) (i int) {
	*slice = append(*slice, value)
	i = slice.Len()
	return i
}

func (slice *Strings) Peek(i int) (value string) {
	value, _ = slice.Get(i)
	return value
}

func (slice *Strings) Set(s *Strings) *Strings {
	*slice = *s
	return slice
}

func (slice *Strings) Sort() *Strings {
	sort.Sort(slice)
	return slice
}

func (slice *Strings) Swap(i int, j int) {
	s := *slice
	s[i], s[j] = s[j], s[i]
	*slice = s
}

func (slice *Strings) Trim(i int, j int) (ok bool) {
	if ok = i < j && slice.Bounds(i) && slice.Bounds(j); ok {
		*slice = (*slice)[i:j]
	}
	return ok
}
