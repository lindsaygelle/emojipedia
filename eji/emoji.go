package eji

type Emoji struct {
	Category    string   `json:"Category"`
	Code        string   `json:"Code"`
	Keywords    []string `json:"Keywords"`
	Name        string   `json:"Name"`
	Number      int      `json:"Number"`
	Sample      string   `json:"Sample"`
	SubCategory string   `json:"SubCategory"`
	Unicode     string   `json:"Unicode"`
}