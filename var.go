package main

import (
	"os"
	"text/tabwriter"
)

var (
	writer = new(tabwriter.Writer).Init(os.Stdout, 0, 8, 0, '\t', 0)
)

var (
	categoriesCommand = []string{C, CATEGORIES}
	categoryCommand   = []string{CC, CATEGORY}
	keywordsCommand   = []string{K, KEYWORDS}
)
