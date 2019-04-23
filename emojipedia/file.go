package emojipedia

type File struct {
	Directory string   `json:"directory"`
	Format    string   `json:"format"`
	Name      string   `json:"name"`
	Size      FileSize `json:"size"`
}
