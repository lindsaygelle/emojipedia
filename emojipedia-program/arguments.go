package program

type Arguments []*Argument

func (slice *Arguments) Bounds(i int) (ok bool) {
	ok = ((i > -1) && (i < len(*slice)))
	return ok
}

func (slice *Arguments) Drop(i int) *Arguments {
	switch i {
	case 0:
		*slice = (*slice)[1:]
	case (slice.Len() - 1):
		*slice = (*slice)[:(slice.Len() - 1)]
	default:
		*slice = append((*slice)[:i], (*slice)[i:]...)
	}
	return slice
}

func (slice *Arguments) Each(f func(i int, argument *Argument)) *Arguments {
	for i, argument := range *slice {
		f(i, argument)
	}
	return slice
}

func (slice *Arguments) Get(i int) (argument *Argument, ok bool) {
	if ok = slice.Bounds(i); ok {
		argument = (*slice)[i]
	}
	return argument, ok
}

func (slice *Arguments) Has(argument *Argument) (position int, ok bool) {
	for i, arg := range *slice {
		if ok = (argument == arg); ok {
			return i, ok
		}
	}
	return -1, ok
}

func (slice *Arguments) Last() (argument *Argument) {
	argument = slice.Peek(slice.Len() - 1)
	return argument
}

func (slice *Arguments) Len() (length int) {
	length = len(*slice)
	return length
}

func (slice *Arguments) Push(argument *Argument) (i int) {
	*slice = append(*slice, argument)
	i = slice.Len()
	return i
}

func (slice *Arguments) Peek(i int) (argument *Argument) {
	argument, _ = slice.Get(i)
	return argument
}

func (slice *Arguments) New(a ...*Argument) *Arguments {
	for i := range a {
		slice.Push(a[i])
	}
	return slice
}

func (slice *Arguments) Set(s *Arguments) *Arguments {
	*slice = *s
	return slice
}

func (slice *Arguments) Swap(i int, j int) {
	s := *slice
	s[i], s[j] = s[j], s[i]
	*slice = s
}

func (slice *Arguments) Trim(i int, j int) (ok bool) {
	if ok = i < j && slice.Bounds(i) && slice.Bounds(j); ok {
		*slice = (*slice)[i:j]
	}
	return ok
}
