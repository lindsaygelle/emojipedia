package emojipedia

type Article struct {
	Title string
	Aside []Aside
}

type Aside struct {
	Key     string
	Verbose string
}

var a = Article{
	Title: "create local dependencies",
	Aside: []Aside{
		Aside{
			Key:     "create",
			Verbose: "Request and store the HTML content from unicode.org."}}}

var Template = `{{ range $_, $article:= . }}{{ .Title }}
{{ range $j, $article.Aside }}%s{{ .Key }}%s{{ .Verbose }}
{{ end }}
{{ end }}`
