package lexicon

import (
	"github.com/gellel/emojipedia/slice"
)

var (
	_ lexicon = (*Lexicon)(nil)
)

// New instantiates a new empty Lexicon pointer.
// Lexicon pointers are mutable and hold and arguments as an interface.
// Unlike basic map-like objects, the Lexicon provides safe getters and setters,
// aiming to reduce the likelyhood of an exception being thrown during an operation.
func New() *Lexicon {
	return &Lexicon{}
}

type lexicon interface {
	Add(key string, value interface{}) *Lexicon
	Concatenate(lexicon *Lexicon) *Lexicon
	Each(f func(key string, value interface{})) *Lexicon
	Fetch(key string) interface{}
	Get(key string) (interface{}, bool)
	Has(key string) bool
	Keys() *slice.Slice
	Len() int
	Map(f func(key string, value interface{}) interface{}) *Lexicon
	Remove(key string) bool
	Values() *slice.Slice
}

// Lexicon is a map-like object whose methods are used to perform traversal and mutation operations by key-value pair.
type Lexicon map[string]interface{}

// Add method adds one element to the Lexicon using the key reference and returns the modified Lexicon.
func (pointer *Lexicon) Add(key string, value interface{}) *Lexicon {
	(*pointer)[key] = value
	return pointer
}

// Each method executes a provided function once for each Lexicon element.
func (pointer *Lexicon) Each(f func(key string, value interface{})) *Lexicon {
	for key, value := range *pointer {
		f(key, value)
	}
	return pointer
}

// Concatenate merges two Lexicons.
func (pointer *Lexicon) Concatenate(lexicon *Lexicon) *Lexicon {
	lexicon.Each(func(key string, value interface{}) {
		pointer.Add(key, value)
	})
	return pointer
}

// Fetch retrieves the interface held by the argument key. Returns nil if key does not exist.
func (pointer *Lexicon) Fetch(key string) interface{} {
	return (*pointer)[key]
}

// Get returns the interface held by the argument key and a boolean indicating if it was successfully retrieved.
func (pointer *Lexicon) Get(key string) (interface{}, bool) {
	value, ok := (*pointer)[key]
	return value, ok
}

// Has method checks that a given key exists in the Lexicon.
func (pointer *Lexicon) Has(key string) bool {
	_, ok := (*pointer)[key]
	return ok
}

// Keys method returns a Slice of a given Lexicon's own property names, in the same order as we get with a normal loop.
func (pointer *Lexicon) Keys() *slice.Slice {
	slice := slice.New()
	pointer.Each(func(key string, value interface{}) {
		slice.Append(key)
	})
	return slice
}

// Len method returns the number of elements in the Lexicon.
func (pointer *Lexicon) Len() int {
	return len(*pointer)
}

// Map method executes a provided function once for each Lexicon element and sets the returned value to the current key.
func (pointer *Lexicon) Map(f func(key string, value interface{}) interface{}) *Lexicon {
	pointer.Each(func(key string, value interface{}) {
		pointer.Replace(key, f(key, value))
	})
	return pointer
}

// Missing method checks if a key is not present in the Lexicon.
func (pointer *Lexicon) Missing(key string) bool {
	return pointer.Has(key) == false
}

// Remove method removes a entry from the Lexicon if it exists. Returns a boolean to confirm if it succeeded.
func (pointer *Lexicon) Remove(key string) bool {
	ok := pointer.Has(key)
	if ok == true {
		delete(*pointer, key)
	}
	return ok
}

// Replace method changes the contents of the Lexicon at the argument key if it exists in the Lexicon.
func (pointer *Lexicon) Replace(key string, value interface{}) bool {
	ok := pointer.Has(key)
	if ok == true {
		(*pointer)[key] = value
	}
	return ok
}

// Values method returns a Slice of a given Lexicon's own enumerable property values, in the same order as that provided by a for...in loop.
func (pointer *Lexicon) Values() *slice.Slice {
	slice := slice.New()
	pointer.Each(func(key string, value interface{}) {
		slice.Append(value)
	})
	return slice
}
