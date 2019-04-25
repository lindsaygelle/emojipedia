package emojipedia

type FileInfo struct {
	Directory string   `json:"directory"`
	Format    string   `json:"format"`
	Name      string   `json:"name"`
	Size      FileSize `json:"size"`
}
