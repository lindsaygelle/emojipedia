package emojipedia

// Section is a collection of standard-in arguments that specifically relate.
type Section struct {
	About     string     `json:"about"`
	Arguments []Argument `json:"arguments"`
}
