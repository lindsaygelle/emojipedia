package emojipedia

// A Category of Emoji.
type Category struct {
	Anchor        string   `json:"anchor"`
	Emoji         *Strings `json:"emoji"`
	Href          string   `json:"href"`
	Position      int      `json:"position"`
	Name          string   `json:"name"`
	Number        int      `json:"number"`
	Subcategories *Strings `json:"subcategories"`
}
