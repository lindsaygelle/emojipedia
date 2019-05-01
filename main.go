package main

import "github.com/gellel/emojipedia/emojipedia"

var (
	categories = emojipedia.Section{
		About: "explore bundled emoji information by category",
		Arguments: []emojipedia.Argument{
			emojipedia.Argument{
				Abbreviation: "-c",
				About:        "execute program to interface with emoji-categories",
				Key:          "categories"}}}

	emoji = emojipedia.Section{
		About: "explore emoji specific information",
		Arguments: []emojipedia.Argument{
			emojipedia.Argument{
				Abbreviation: "-e",
				About:        "explore information by emoji-name",
				Key:          "emoji"}}}

	unicode = emojipedia.Section{
		About: "manage content from unicode.org",
		Arguments: []emojipedia.Argument{
			emojipedia.Argument{
				Abbreviation: "-u",
				About:        "execute program to manage data from unicode.org",
				Key:          "unicode"}}}

	feature = emojipedia.Feature{
		Sections: []emojipedia.Section{categories, emoji, unicode}}
)

func main() {

	feature.Describe()
}
