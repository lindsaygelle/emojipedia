package emojipedia

// Emoji is a struct expressing an Emoji and its data.
type Emoji struct {
	Anchor      string   `json:"anchor"`
	Category    string   `json:"category"`
	Codes       *Strings `json:"codes"`
	Description string   `json:"description"`
	Href        string   `json:"href"`
	Image       string   `json:"img"`
	Keywords    *Strings `json:"keywords"`
	Name        string   `json:"name"`
	Number      int      `json:"number"`
	Position    int      `json:"position"`
	Subcategory string   `json:"subcategory"`
	Unicode     string   `json:"unicode"`
}
