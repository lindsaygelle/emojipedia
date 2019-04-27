package x

import "fmt"

const (
	about string = "Create, manage or examine the content collected from unicode.org.\nContents are used to build out supporting emojipedia program dependencies."
)

const (
	cache string = "Request and store the HTML content from unicode.org."
)

const (
	information string = "Get information about the stored HTML content from unicode.org."
)

const (
	remove string = "Remove the locally cached content collected from unicode.org."
)

const (
	append string = "create local dependencies"
)

const (
	details string = "examine the content of the local dependencies"
)

const (
	manage string = "manage the local dependencies."
)

var help = `usage: emojipedia unicode <command> [<args>]

%s

Available "$ emojipedia unicode" options:

%s
   cache       %s

%s
   information %s

%s
   remove      %s
`

func Help() string {
	return fmt.Sprintf(help, about, append, cache, details, information, manage, remove)
}
