package manifest

type Manifest struct {
	Author      string `json:"author"`
	Description string `json:"description"`
	Programs    map[string]*Program
	Name        string  `json:"name"`
	Version     float64 `json:"version"`
}
