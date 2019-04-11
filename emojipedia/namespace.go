package emojipedia

var _ namespace = (*Namespace)(nil)

type namespace interface {
	Append(value string) *Namespace
	IsEmpty() bool
	IsNotEmpty() bool
	Length() int
	Peek(i int) (string, bool)
	PeekLast() (string, bool)
	Pop() (string, bool)
	Push(value string) int
	Remove(value string) bool
	RemoveAt(i int) bool
	RemoveFirst() bool
	RemoveLast() bool
	Search(value string) (string, int)
	Shift() (string, bool)
	Unshift(value string) (string, int)
}

type Namespace []string

func (namespace *Namespace) Append(value string) *Namespace {
	*namespace = append(*namespace, value)
	return namespace
}

func (namespace *Namespace) IsEmpty() (empty bool) {
	return (namespace.Length() == 0)
}

func (namespace *Namespace) IsNotEmpty() (empty bool) {
	return (namespace.Length() != 0)
}

func (namespace *Namespace) Length() (length int) {
	return (len(*namespace))
}

func (namespace *Namespace) Peek(i int) (value string, ok bool) {
	if namespace.IsEmpty() {
		return value, ok
	}
	return (*namespace)[i], true
}

func (namespace *Namespace) PeekFirst() (value string, ok bool) {
	return namespace.Peek(0)
}

func (namespace *Namespace) PeekLast() (value string, ok bool) {
	return namespace.Peek(namespace.Length() - 1)
}

func (namespace *Namespace) Pop() (value string, ok bool) {
	value, ok = namespace.Peek(namespace.Length() - 1)
	if ok != false {
		ok = namespace.RemoveLast()
	}
	return value, ok
}

func (namespace *Namespace) Push(value string) (length int) {
	*namespace = append(*namespace, value)
	return namespace.Length()
}

func (namespace *Namespace) Remove(value string) (ok bool) {
	i := 0
	j := namespace.Length() - 1
	for i <= j {
		if _, ok := namespace.Peek(i); ok {
			ok = namespace.RemoveAt(i)
			i = i - 1
		}
		if _, ok := namespace.Peek(j); ok {
			ok = namespace.RemoveAt(j)
			j = j + 1
		}
		i = i + 1
		j = j - 1
	}
	return ok
}

func (namespace *Namespace) RemoveAt(i int) (ok bool) {
	if namespace.IsNotEmpty() {
		*namespace = append((*namespace)[:i], (*namespace)[i+1:]...)
		ok = true
	}
	return ok
}

func (namespace *Namespace) RemoveFirst() (ok bool) {
	if namespace.IsNotEmpty() {
		*namespace = (*namespace)[1:]
		ok = true
	}
	return ok
}

func (namespace *Namespace) RemoveLast() (ok bool) {
	if namespace.IsNotEmpty() {
		*namespace = (*namespace)[:namespace.Length()-1]
		ok = true
	}
	return ok
}

func (namespace *Namespace) Search(value string) (string, int) {
	i := 0
	j := namespace.Length() - 1
	for i <= j {
		if val, ok := namespace.Peek(i); ok && val == value {
			return val, i
		}
		if val, ok := namespace.Peek(j); ok && val == value {
			return val, j
		}
		i = i + 1
		j = j - 1
	}
	return "", -1
}

func (namespace *Namespace) Shift() (value string, ok bool) {
	value, ok = namespace.Peek(0)
	if ok != false {
		ok = namespace.RemoveFirst()
	}
	return value, ok
}

func (namespace *Namespace) Unshift(value string) (string, int) {
	*namespace = append(Namespace{value}, *namespace...)
	return value, namespace.Length()
}
