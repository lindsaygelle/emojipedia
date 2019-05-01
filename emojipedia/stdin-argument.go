package emojipedia

import "fmt"

// Argument is a program standard-in argument.
type Argument struct {
	Abbreviation string `json:"abbreviation"`
	About        string `json:"about"`
	Key          string `json:"key"`
}

func (argument Argument) String() string {
	return fmt.Sprintf("%s, %s\t%s", argument.Abbreviation, argument.Key, argument.About)
}
