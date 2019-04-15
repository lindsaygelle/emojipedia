package api

import (
	"github.com/gellel/emojipedia/emojipedia"

	categories "github.com/gellel/emojipedia/emojipedia-files/files-categories"
	emojis "github.com/gellel/emojipedia/emojipedia-files/files-emojis"
	keywords "github.com/gellel/emojipedia/emojipedia-files/files-keywords"
	subcategories "github.com/gellel/emojipedia/emojipedia-files/files-subcategories"
)

func GetCategories() (map[int]string, error) {
	return categories.Open()
}

func GetEmojis() (map[string]*emojipedia.Emoji, error) {
	return emojis.Open()
}

func GetKeywords() (map[string][]string, error) {
	return keywords.Open()
}

func GetSubcategories() (map[int]string, error) {
	return subcategories.Open()
}
