package emojipedia

type associative interface {
	Add(key string, value string)
	Get(key string) (*Namespace, bool)
	GetUnsafely(key string) *Namespace
	Has(key string)
	Keys() *Namespace
	Length() int
	Remove(key string)
	RemoveValue(key string, value string) bool
	Set(key string) *Namespace
}

type Associative map[string]*Namespace

func (associative *Associative) Add(key string, value string) {
	if _, ok := (*associative)[key]; ok != true {
		(*associative)[key] = &Namespace{value}
	} else {
		(*associative)[key].Push(value)
	}
}

func (associative *Associative) Get(key string) (namespace *Namespace, ok bool) {
	namespace, ok = (*associative)[key]
	return namespace, ok
}

func (associative *Associative) GetUnsafely(key string) (namespace *Namespace) {
	namespace, _ = (*associative)[key]
	return namespace
}

func (associative *Associative) Has(key string) (ok bool) {
	_, ok = (*associative)[key]
	return ok
}

func (associative *Associative) Keys() (keys []string) {
	keyset := make([]string, 0, len(*associative))
	for key := range *associative {
		keyset = append(keyset, key)
	}
	return keys
}

func (associative *Associative) Length() (length int) {
	return len(*associative)
}

func (associative *Associative) Set(key string) string {
	(*associative)[key] = &Namespace{}
	return key
}
