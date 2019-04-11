package emojipedia

type Encyclopedia struct {
	Categories    *Associative
	Subcategories *Associative
	Keywords      *Associative
	Numeric       *Set
}
