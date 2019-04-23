package emojipedia

// A Subcategory of Emoji.
type Subcategory struct {
	Anchor   string   `json:"anchor"`
	Category string   `json:"category"`
	Emoji    *Strings `json:"emoji"`
	Href     string   `json:"href"`
	Position int      `json:"position"`
	Name     string   `json:"name"`
	Number   int      `json:"number"`
}
