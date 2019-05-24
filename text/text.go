package text

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

var (
	replacements = []string{" ", "-", "(", "", ")", "", "&", "and", ":", "", ",", "", ".", "", "⊛", "", "“", "", "”", "", "’", ""}
	replacer     = strings.NewReplacer(replacements...)
	transformer  = transform.Chain(norm.NFD, transform.RemoveFunc(func(r rune) bool { return unicode.Is(unicode.Mn, r) }), norm.NFC)
)

// Emojize transforms an escaped emoji unicode string to its glyph counterpart.
func Emojize(s string) string {
	r, _ := strconv.ParseInt(strings.TrimPrefix(s, "\\U"), 16, 32)
	return fmt.Sprintf("%s", string(r))
}

// Normalize trims and replaces all non utf-8 characters from the argument string.
func Normalize(s string) string {
	s, _, _ = transform.String(transformer, s)
	s = strings.ToLower(s)
	s = replacer.Replace(strings.TrimSpace(s))
	if len(s) != 0 && strings.HasPrefix(s, "-") {
		s = strings.TrimPrefix(s, "-")
	}
	if len(s) != 0 && strings.HasSuffix(s, "-") {
		s = strings.TrimSuffix(s, "-")
	}
	return s
}
