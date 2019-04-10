package emojipedia

var _ n = (*N)(nil)

type n interface {
	Append(name string) *N
	IsEmpty() bool
	IsNotEmpty() bool
	Len() int
	Length() int
	Peek(i int) (string, bool)
	Pop() (string, bool)
	Push(name string) int
	Remove(name string) bool
	RemoveAt(i int) bool
	RemoveFirst() bool
	RemoveLast() bool
	Shift() (string, bool)
	Unshift(name string) (string, int)
}

type N []string

func (names *N) Append(name string) *N {
	*names = append(*names, name)
	return names
}

func (names *N) IsEmpty() (empty bool) {
	return (names.Length() == 0)
}

func (names *N) IsNotEmpty() (empty bool) {
	return (names.Length() != 0)
}

func (names *N) Len() (n int) {
	return (len(*names) - 1)
}
func (names *N) Length() (n int) {
	return (len(*names))
}

func (names *N) Peek(i int) (name string, ok bool) {
	if names.IsNotEmpty() {
		return name, ok
	}
	return (*names)[i], true
}

func (names *N) Pop() (name string, ok bool) {
	name, ok = names.Peek(names.Len())
	if ok != false {
		ok = names.RemoveLast()
	}
	return name, ok
}

func (names *N) Push(name string) (n int) {
	*names = append(*names, name)
	n = names.Length()
	return n
}

func (names *N) Remove(name string) (ok bool) {
	i := 0
	j := names.Len()
	for i <= j {
		if _, ok := names.Peek(i); ok {
			ok = names.RemoveAt(i)
			i = i - 1
		}
		if _, ok := names.Peek(j); ok {
			ok = names.RemoveAt(j)
			j = j + 1
		}
		i = i + 1
		j = j - 1
	}
	return ok
}

func (names *N) RemoveAt(i int) (ok bool) {
	if names.IsNotEmpty() {
		*names = append((*names)[:i], (*names)[i+1:]...)
		ok = true
	}
	return ok
}

func (names *N) RemoveFirst() (ok bool) {
	if names.IsNotEmpty() {
		*names = (*names)[1:]
		ok = true
	}
	return ok
}

func (names *N) RemoveLast() (ok bool) {
	if names.IsNotEmpty() {
		*names = (*names)[:names.Len()]
		ok = true
	}
	return ok
}

func (names *N) Shift() (name string, ok bool) {
	name, ok = names.Peek(0)
	if ok != false {
		ok = names.RemoveFirst()
	}
	return name, ok
}

func (names *N) Unshift(name string) (string, int) {
	*names = append(N{name}, *names...)
	return name, names.Length()
}
