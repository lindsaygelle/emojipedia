package emojipedia

func NewEmojidex() *Emojidex {
	return &Emojidex{}
}

func NewEncyclopedia() *Encyclopedia {
	return &Encyclopedia{
		Category:    &Ensemble{},
		Subcategory: &Ensemble{},
		Keywords:    &Ensemble{},
		Numeric:     &Ensemble{}}
}
