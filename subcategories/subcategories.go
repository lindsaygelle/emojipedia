package subcategories

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
	"text/tabwriter"

	"github.com/gellel/emojipedia/directory"

	"github.com/PuerkitoBio/goquery"
	"github.com/gellel/emojipedia/lexicon"
	"github.com/gellel/emojipedia/pkg"
	"github.com/gellel/emojipedia/slice"
	"github.com/gellel/emojipedia/subcategory"
	"github.com/gellel/emojipedia/text"
)

var _ subcategories = (*Subcategories)(nil)

// New instantiates a new empty Subcategories pointer.
func New() *Subcategories {
	return &Subcategories{&lexicon.Lexicon{}}
}

// NewSubcategories creates a new Subcategories pointer, accepting zero or more subcategory.Subcategory pointers as arguments.
func NewSubcategories(subcategory ...*subcategory.Subcategory) *Subcategories {
	subcategories := &Subcategories{&lexicon.Lexicon{}}
	for _, subcategory := range subcategory {
		subcategories.Add(subcategory)
	}
	return subcategories
}

// Get attempts to open all Category data from the emojipedia/subcategories folder, but panics if an error occurs.
func Get() *Subcategories {
	subcategories, err := Open()
	if err != nil {
		panic(err)
	}
	return subcategories
}

func Make(document *goquery.Document) {
	var key, category string
	subcategories := New()
	document.Find("tr").Each(func(i int, selection *goquery.Selection) {
		selection.Find("th.bighead a").Each(func(j int, s *goquery.Selection) {
			category = text.Normalize(s.Text())
		})
		selection.Find("th.mediumhead a").Each(func(j int, s *goquery.Selection) {
			var (
				anchor, _   = s.Attr("href")
				emoji       = &slice.Slice{}
				href        = (pkg.URL + anchor)
				position    = i
				name        = text.Normalize(s.Text())
				number      = subcategories.Len()
				subcategory = subcategory.NewSubcategory(anchor, category, href, name, number, position, emoji)
			)
			subcategories.Add(subcategory)
			key = subcategory.Name
		})
		selection.Find("td").Eq(3).Each(func(j int, s *goquery.Selection) {
			var (
				name           = text.Normalize(s.Text())
				subcategory, _ = subcategories.Get(key)
			)
			subcategory.Emoji.Append(name)
		})
	})
	subcategories.Each(func(s *subcategory.Subcategory) {
		subcategory.Write(s)
	})
}

// Open attempts to open all Category data from the emojipedia/subcategories folder.
func Open() (*Subcategories, error) {
	files, err := ioutil.ReadDir(directory.Subcategory)
	if err != nil {
		return nil, err
	}
	subcategories := New()
	for _, file := range files {
		name := strings.TrimSuffix(file.Name(), ".json")
		subcategory, err := subcategory.Open(name)
		if err != nil {
			return nil, err
		}
		subcategories.Add(subcategory)
	}
	return subcategories, nil
}

// Remove deletes all Subcategory data stored in the dependencies folder.
func Remove() error {
	return os.Remove(directory.Subcategory)
}

type subcategories interface {
	Add(subcategory *subcategory.Subcategory) *Subcategories
	Each(f func(subcategory *subcategory.Subcategory)) *Subcategories
	Fetch(key string) *subcategory.Subcategory
	Get(key string) (*subcategory.Subcategory, bool)
	Has(key string) bool
	Keys() *slice.Slice
	Len() int
	Remove(key string) bool
	Values() *slice.Slice
}

// Subcategories is a map-like struct with methods used to perform traversal and retrieval of subcategory.Subcategory pointers.
type Subcategories struct {
	lexicon *lexicon.Lexicon
}

// Add method adds one subcategory.Subcategory to the Subcategories using the subcategory.Subcategory.Name as the key reference.
func (pointer *Subcategories) Add(subcategory *subcategory.Subcategory) *Subcategories {
	pointer.lexicon.Add(subcategory.Name, subcategory)
	return pointer
}

// Each method executes a provided function once for each subcategory.Subcategory pointer.
func (pointer *Subcategories) Each(f func(subcategory *subcategory.Subcategory)) *Subcategories {
	pointer.lexicon.Each(func(key string, i interface{}) {
		f(i.(*subcategory.Subcategory))
	})
	return pointer
}

// Fetch retrieves the subcategory.Subcategory pointer held by the argument key. Panics if key does not exist.
func (pointer *Subcategories) Fetch(key string) *subcategory.Subcategory {
	property, _ := pointer.Get(key)
	return property
}

// Get returns the subcategory.Subcategory pointer held by the argument key and a boolean indicating if it was successfully retrieved.
// Panics if cannot convert to subcategory.Subcategory pointer.
func (pointer *Subcategories) Get(key string) (*subcategory.Subcategory, bool) {
	property, ok := pointer.lexicon.Get(key)
	return property.(*subcategory.Subcategory), ok
}

// Has method checks that a given key exists in the Subcategories.
func (pointer *Subcategories) Has(key string) bool {
	return pointer.lexicon.Has(key)
}

// Keys method returns a slice.Slice of a given Subcategories' own property names, in the same order as we get with a normal loop.
func (pointer *Subcategories) Keys() *slice.Slice {
	slice := slice.New()
	pointer.lexicon.Each(func(key string, i interface{}) {
		slice.Append(key)
	})
	return slice
}

// Len method returns the number of elements in the Subcategories.
func (pointer *Subcategories) Len() int {
	return pointer.lexicon.Len()
}

// Remove method removes a entry from the Subcategories if it exists. Returns a boolean to confirm if it succeeded.
func (pointer *Subcategories) Remove(key string) bool {
	return pointer.lexicon.Remove(key)
}

// List prints out all subcategories in a tabwriter.
func (pointer *Subcategories) List() {
	writer := new(tabwriter.Writer)
	writer.Init(os.Stdout, 0, 8, 0, '\t', 0)
	keys := []string{"subcategory", "category", "number", "emoji"}
	fmt.Fprintln(writer, strings.Join(keys, "\t"))
	fmt.Fprintln(writer, strings.Join([]string{"-", "-", "-", "-"}, "\t"))
	pointer.Each(func(subcategory *subcategory.Subcategory) {
		values := []string{
			subcategory.Name,
			subcategory.Category,
			fmt.Sprintf("%v", subcategory.Number),
			fmt.Sprintf("%v", subcategory.Emoji.Len())}
		fmt.Fprintln(writer, strings.Join(values, "\t"))
	})
	fmt.Fprintln(writer)
	writer.Flush()
}

// Values method returns a Slice of a given Subcategories's own enumerable property values,
// in the same order as that provided by a for...in loop.
func (pointer *Subcategories) Values() *slice.Slice {
	slice := slice.New()
	pointer.lexicon.Each(func(key string, i interface{}) {
		slice.Append(i)
	})
	return slice
}
