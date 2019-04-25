package text

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

var replacements = []string{
	" ", "-",
	"(", "",
	")", "",
	"&", "and",
	":", "",
	",", "",
	".", "",
	"⊛", "",
	"“", "",
	"”", "",
	"’", ""}

var replacer = strings.NewReplacer(replacements...)

func Emojize(value string) (emoji string) {
	r, _ := strconv.ParseInt(strings.TrimPrefix(value, "\\U"), 16, 32)
	return fmt.Sprintf("%s", string(r))
}

func Normalizer(r rune) (ok bool) {
	ok = unicode.Is(unicode.Mn, r)
	return ok
}

func Normalize(value string) (result string) {
	transformer := transform.Chain(norm.NFD, transform.RemoveFunc(Normalizer), norm.NFC)
	result, _, _ = transform.String(transformer, value)
	result = replacer.Replace(strings.TrimSpace(result))
	result = strings.ToLower(result)
	if len(result) != 1 && strings.HasPrefix(result, "-") {
		result = strings.TrimPrefix(result, "-")
	}
	if len(result) != 1 && strings.HasSuffix(result, "-") {
		result = strings.TrimSuffix(result, "-")
	}
	return result
}
