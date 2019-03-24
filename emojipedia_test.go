package emojipedia_test

import (
	"testing"

	"github.com/gellel/emojipedia"
)

func Test(t *testing.T) {
	page, ok := emojipedia.GetPage(emojipedia.URL)

	if ok == nil {
		if categories, ok := emojipedia.GetCategories(page); ok == nil {

			for _, category := range categories {
				emojipedia.GetCategoryPage(category)
				break
			}
		}
	}
}
