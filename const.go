package main

const (
	ANCHOR        string = "ANCHOR"
	CATEGORIES    string = "CATEGORIES"
	CATEGORY      string = "CATEGORY"
	CODES         string = "CODES"
	DESCRIPTION   string = "DESCRIPTION"
	EMOJIPEDIA    string = "EMOJIPEDIA"
	EMOJI         string = "EMOJI"
	IMAGE         string = "IMAGE"
	HREF          string = "HREF"
	KEYWORDS      string = "KEYWORDS"
	NUMBER        string = "NUMBER"
	SUBCATEGORIES string = "SUBCATEGORIES"
	SUBCATEGORY   string = "SUBCATEGORY"
	UNICODE       string = "UNICODE"
)

const (
	B     string = "-B"
	BUILD string = "BUILD"
)

const (
	G   string = "-G"
	GET string = "GET"
)

const (
	H    string = "-H"
	HELP string = "HELP"
)

const (
	K    string = "-K"
	KEYS string = "KEYS"
)

const (
	L    string = "-L"
	LIST string = "LIST"
)

const (
	P        string = "-P"
	POSITION string = "POSITION"
)

const (
	R      string = "-R"
	REMOVE string = "REMOVE"
)

const (
	T     string = "-T"
	TABLE string = "TABLE"
)

const (
	errorCannotFind    string = "cannot find dependency \"%s\". content either missing or not built"
	errorCannotOpen    string = "cannot open \"%s\"; encountered unexpected error \"%s\""
	errorRemovePackage string = "cannot remove \"%s\"; encountered error \"%s\""
)

const (
	statusBuildPackage  string = "attempting to build \"%s\" package"
	statusRemovePackage string = "attempting to remove \"%s\" package; deleting core packages can affect building!"
)

const (
	successBuildPackage  string = "success! program has built package \"%s\""
	successRemovePackage string = "success! program has removed \"%s\"!"
)

const (
	errorChoiceNotFound string = "Uh-oh. Cannot find content \"%s\" in choice \"$ emojipedia [%s|%s] <choice>\". Please check input and try again."
)
