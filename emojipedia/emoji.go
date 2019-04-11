package emojipedia

type Emoji struct {
	Category    string   `json:"Category"`
	Code        string   `json:"Code"`
	Keywords    []string `json:"Keywords"`
	Name        string   `json:"Name"`
	Number      int      `json:"Number"`
	Sample      string   `json:"Sample"`
	Subcategory string   `json:"Subcategory"`
	Unicode     string   `json:"Unicode"`
}
