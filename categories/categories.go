package categories

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/PuerkitoBio/goquery"
	"github.com/gellel/emojipedia/category"
	"github.com/gellel/emojipedia/directory"
	"github.com/gellel/emojipedia/lexicon"
	"github.com/gellel/emojipedia/pkg"
	"github.com/gellel/emojipedia/slice"
	"github.com/gellel/emojipedia/text"
)

var _ categories = (*Categories)(nil)

// New instantiates a new empty Categories pointer.
func New() *Categories {
	return &Categories{&lexicon.Lexicon{}}
}

// NewCategories creates a new Categories pointer, accepting zero or more category.Category pointers as arguments.
func NewCategories(category ...*category.Category) *Categories {
	categories := &Categories{&lexicon.Lexicon{}}
	for _, category := range category {
		categories.Add(category)
	}
	return categories
}

// Get attempts to open all Category data from the emojipedia/categories folder, but panics if an error occurs.
func Get() *Categories {
	categories, err := Open()
	if err != nil {
		panic(err)
	}
	return categories
}

func Make(document *goquery.Document) {
	var key string
	categories := New()
	document.Find("tr").Each(func(i int, selection *goquery.Selection) {
		selection.Find("th.bighead a").Each(func(j int, s *goquery.Selection) {
			var (
				anchor, _     = s.Attr("href")
				emoji         = &slice.Slice{}
				href          = (pkg.URL + anchor)
				position      = i
				name          = text.Normalize(s.Text())
				number        = categories.Len()
				subcategories = &slice.Slice{}
				category      = category.NewCategory(anchor, href, name, number, position, emoji, subcategories)
			)
			categories.Add(category)
			key = category.Name
		})
		selection.Find("th.mediumhead a").Each(func(j int, s *goquery.Selection) {
			var (
				category, _ = categories.Get(key)
				subcategory = text.Normalize(s.Text())
			)
			category.Subcategories.Append(subcategory)
		})
		selection.Find("td").Eq(3).Each(func(j int, s *goquery.Selection) {
			var (
				category, _ = categories.Get(key)
				name        = text.Normalize(s.Text())
			)
			category.Emoji.Append(name)
		})
	})
	categories.Each(func(c *category.Category) {
		category.Write(c)
	})
}

// Open attempts to open all Category data from the emojipedia/categories folder.
func Open() (*Categories, error) {
	files, err := ioutil.ReadDir(directory.Category)
	if err != nil {
		return nil, err
	}
	categories := New()
	for _, file := range files {
		name := strings.TrimSuffix(file.Name(), ".json")
		category, err := category.Open(name)
		if err != nil {
			return nil, err
		}
		categories.Add(category)
	}
	return categories, nil
}

// Remove deletes all Category data stored in the dependencies folder.
func Remove() error {
	return os.Remove(directory.Category)
}

type categories interface {
	Add(category *category.Category) *Categories
	Each(f func(category *category.Category)) *Categories
	Fetch(key string) *category.Category
	Get(key string) (*category.Category, bool)
	Has(key string) bool
	Keys() *slice.Slice
	Len() int
	Remove(key string) bool
	Values() *slice.Slice
}

// Categories is a map-like struct with methods used to perform traversal and retrieval of category.Category pointers.
type Categories struct {
	lexicon *lexicon.Lexicon
}

// Add method adds one category.Category to the Categories using the category.Category.Name as the key reference.
func (pointer *Categories) Add(category *category.Category) *Categories {
	pointer.lexicon.Add(category.Name, category)
	return pointer
}

// Each method executes a provided function once for each category.Category pointer.
func (pointer *Categories) Each(f func(category *category.Category)) *Categories {
	pointer.lexicon.Each(func(key string, i interface{}) {
		f(i.(*category.Category))
	})
	return pointer
}

// Fetch retrieves the category.Category pointer held by the argument key. Panics if key does not exist.
func (pointer *Categories) Fetch(key string) *category.Category {
	property, _ := pointer.Get(key)
	return property
}

// Get returns the category.Category pointer held by the argument key and a boolean indicating if it was successfully retrieved.
// Panics if cannot convert to category.Category pointer.
func (pointer *Categories) Get(key string) (*category.Category, bool) {
	property, ok := pointer.lexicon.Get(key)
	return property.(*category.Category), ok
}

// Has method checks that a given key exists in the Categories.
func (pointer *Categories) Has(key string) bool {
	return pointer.lexicon.Has(key)
}

// Keys method returns a slice.Slice of a given Categories' own property names, in the same order as we get with a normal loop.
func (pointer *Categories) Keys() *slice.Slice {
	slice := slice.New()
	pointer.lexicon.Each(func(key string, i interface{}) {
		slice.Append(key)
	})
	return slice
}

// Len method returns the number of elements in the Categories.
func (pointer *Categories) Len() int {
	return pointer.lexicon.Len()
}

// Remove method removes a entry from the Categories if it exists. Returns a boolean to confirm if it succeeded.
func (pointer *Categories) Remove(key string) bool {
	return pointer.lexicon.Remove(key)
}

// List prints out all subcategories in a tabwriter.
func (pointer *Categories) List() {
	writer := new(tabwriter.Writer)
	writer.Init(os.Stdout, 0, 8, 0, '\t', 0)
	keys := []string{"category", "number", "emoji"}
	fmt.Fprintln(writer, strings.Join(keys, "\t"))
	fmt.Fprintln(writer, strings.Join([]string{"-", "-", "-"}, "\t"))
	pointer.Each(func(category *category.Category) {
		values := []string{
			category.Name,
			fmt.Sprintf("%v", category.Number),
			fmt.Sprintf("%v", category.Emoji.Len())}
		fmt.Fprintln(writer, strings.Join(values, "\t"))
	})
	fmt.Fprintln(writer)
	writer.Flush()
}

// Values method returns a Slice of a given Categories's own enumerable property values,
// in the same order as that provided by a for...in loop.
func (pointer *Categories) Values() *slice.Slice {
	slice := slice.New()
	pointer.lexicon.Each(func(key string, i interface{}) {
		slice.Append(i)
	})
	return slice
}
