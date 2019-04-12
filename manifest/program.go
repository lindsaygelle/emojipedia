package manifest

type Program struct {
	Description string `json:"description"`
	Programs    map[string]*Program
	Name        string `json:"name"`
}
