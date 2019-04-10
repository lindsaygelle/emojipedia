package emojipedia

import "strings"

var keyReplacements = []string{
	"_", "-",
	"&", "and",
	" ", "-"}

var keyReplacer = strings.NewReplacer(keyReplacements...)

type catalogue interface {
	Add(key string, value string)
	Get(key string) (*Namespace, bool)
	GetUnsafely(key string) *Namespace
	Has(key string)
	Keys() *Namespace
	Length() int
	Normalize(key string) string
	Remove(key string)
	RemoveValue(key string, value string) bool
	Set(key string) *Namespace
}

type Catalogue map[string]*Namespace

func (catalogue *Catalogue) Add(key string, value string) {
	if _, ok := (*catalogue)[key]; ok != true {
		(*catalogue)[key] = &Namespace{value}
	} else {
		(*catalogue)[key].Push(value)
	}
}

func (catalogue *Catalogue) Get(key string) (namespace *Namespace, ok bool) {
	namespace, ok = (*catalogue)[key]
	return namespace, ok
}

func (catalogue *Catalogue) GetUnsafely(key string) (namespace *Namespace) {
	namespace, _ = (*catalogue)[key]
	return namespace
}

func (catalogue *Catalogue) Has(key string) (ok bool) {
	_, ok = (*catalogue)[key]
	return ok
}

func (catalogue *Catalogue) Keys() (keys []string) {
	keyset := make([]string, 0, len(*catalogue))
	for key := range *catalogue {
		keyset = append(keyset, key)
	}
	return keys
}

func (catalogue *Catalogue) Length() (length int) {
	return len(*catalogue)
}

func (catalogue *Catalogue) Normalize(key string) (k string) {
	return keyReplacer.Replace(key)
}

func (catalogue *Catalogue) Set(key string) string {
	(*catalogue)[key] = &Namespace{}
	return key
}
