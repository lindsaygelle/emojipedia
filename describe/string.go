package describe

import "strings"

type String string

func (str String) Assign(c String) String {
	str = (str + c)
	return str
}

func (str String) Append(character string) String {
	str = (str + String(character))
	return str
}

func (str String) Len() (length int) {
	length = len(str)
	return length
}

func (str String) Prepend(character string) String {
	str = (String(character) + str)
	return str
}

func (str String) Preassign(c String) String {
	str = (c + str)
	return str
}

func (str String) Separate() (substrings Substrings) {
	substrings = str.Split(" ")
	return substrings
}

func (str String) Split(character string) (substrings Substrings) {
	for _, substring := range strings.Split(string(str), character) {
		substrings = append(substrings, String(substring))
	}
	return substrings
}

func (str String) Surround(character string) String {
	str = (String(character) + str + String(character))
	return str
}
