package cli

type Manifest struct {
	Author      string `json:"author"`
	Description string `json:"description"`
	Programs    map[string]Routine
	Name        string  `json:"name"`
	Version     float64 `json:"version"`
}

type Routine struct {
	Description string `json:"description"`
	Programs    map[string]Routine
	Name        string `json:"name"`
}
