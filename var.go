package main

import (
	"fmt"
	"os"
	"strings"
	"text/tabwriter"
)

var (
	writer = new(tabwriter.Writer).Init(os.Stdout, 0, 8, 0, '\t', 0)
)

var (
	starting = fmt.Sprintf(param, "short", "verbose", "[-b|build]")
)

var (
	copt = fmt.Sprintf(param, strings.ToLower(C), strings.ToLower(CATEGORIES), categoriesDescription)
	kopt = fmt.Sprintf(param, strings.ToLower(K), strings.ToLower(KEYWORDS), keywordsDescription)
	eopt = fmt.Sprintf(param, strings.ToLower(E), strings.ToLower(EMOJIPEDIA), emojipediaDescription)
	sopt = fmt.Sprintf(param, strings.ToLower(S), strings.ToLower(SUBCATEGORIES), subcategoriesDescription)
)

var (
	ccopt = fmt.Sprintf(param, strings.ToLower(CC), strings.ToLower(CATEGORY), categoryDescription)
	eeopt = fmt.Sprintf(param, strings.ToLower(EE), strings.ToLower(EMOJI), emojiDescription)
	ssopt = fmt.Sprintf(param, strings.ToLower(SS), strings.ToLower(SUBCATEGORY), subcategoryDescription)
)
