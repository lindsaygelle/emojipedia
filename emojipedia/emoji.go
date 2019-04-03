package emojipedia

type Emoji struct {
	Category    string   `json:"Category"`
	Codes       string   `json:"Codes"`
	Description string   `json:"Description"`
	Keywords    []string `json:"Keywords"`
	Name        string   `json:"Name"`
	Number      int      `json:"Number"`
	Sample      string   `json:"Sample"`
	Subcategory string   `json:"Subcategory"`
	Unicode     string   `json:"Unicode"`
}
