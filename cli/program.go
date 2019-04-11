package cli

// A Program is a construct of one CLI main function.
// Each Program holds a series of Function structs which represent the available options for the program.
// Similar to a Git prompt each Program attempts to describe a common usage pattern.
// Each Function in the Functions slice is intended to be a unique function.
type Program struct {
	Description string
	Functions   []*Function
	Name        string
	Use         string
}
