package emojipedia

import (
	"io/ioutil"
	"os"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gellel/emojipedia/directory"
	"github.com/gellel/emojipedia/emoji"
	"github.com/gellel/emojipedia/lexicon"
	"github.com/gellel/emojipedia/pkg"
	"github.com/gellel/emojipedia/slice"
	"github.com/gellel/emojipedia/text"
)

var _ emojipedia = (*Emojipedia)(nil)

// New instantiates a new empty Emojipedia pointer.
func New() *Emojipedia {
	return &Emojipedia{&lexicon.Lexicon{}}
}

// NewEmojipedia creates a new Emojipedia pointer, accepting zero or more emoji.Emoji pointers as arguments.
func NewEmojipedia(emoji ...*emoji.Emoji) *Emojipedia {
	emojipedia := &Emojipedia{&lexicon.Lexicon{}}
	for _, emoji := range emoji {
		emojipedia.Add(emoji)
	}
	return emojipedia
}

// Get attempts to open all Emoji data from the emojipedia/emoji folder, but panics if an error occurs.
func Get() *Emojipedia {
	emojipedia, err := Open()
	if err != nil {
		panic(err)
	}
	return emojipedia
}

// Make builds Emoji dependencies from HTML scraped from unicode.org.
func Make(document *goquery.Document) {
	var category, subcategory string
	document.Find("tr").Each(func(i int, selection *goquery.Selection) {
		var (
			anchor   string
			codes    = &slice.Slice{}
			image    string
			keywords = &slice.Slice{}
			name     string
			number   int
			unicodes string
		)
		selection.Find("th.bighead a").Each(func(j int, s *goquery.Selection) {
			category = text.Normalize(s.Text())
		})
		selection.Find("th.mediumhead a").Each(func(j int, s *goquery.Selection) {
			subcategory = text.Normalize(s.Text())
		})
		selection.Find("td.rchars").Each(func(j int, s *goquery.Selection) {
			number, _ = strconv.Atoi(strings.TrimSpace(s.Text()))
		})
		selection.Find("td.code").Each(func(j int, s *goquery.Selection) {
			for _, substring := range strings.Split(s.Text(), " ") {
				codes.Append(strings.TrimSpace(substring))
			}
		})
		selection.Find("td.andr img").Each(func(j int, s *goquery.Selection) {
			image, _ = s.Attr("src")
		})
		selection.Find("td.andr a").Each(func(j int, s *goquery.Selection) {
			anchor, _ = s.Attr("href")
		})
		selection.Find("td.name").First().Each(func(j int, s *goquery.Selection) {
			name = text.Normalize(s.Text())
		})
		selection.Find("td.name").Last().Each(func(j int, s *goquery.Selection) {
			for _, substring := range strings.Split(s.Text(), "|") {
				keywords.Append(text.Normalize(substring))
			}
		})
		if len(name) == 0 {
			return
		}
		codes.Each(func(_ int, i interface{}) {
			code := i.(string)
			replacement := "000"
			if len(code) == 6 {
				replacement = "0000"
			}
			unicodes = unicodes + strings.Replace(code, "+", replacement, 1)
		})
		unicodes = strings.Replace(strings.ToLower(unicodes), "u", "\\U", -1)
		emoji.Write(&emoji.Emoji{
			Anchor:      anchor,
			Category:    category,
			Codes:       codes,
			Href:        (pkg.URL + anchor),
			Image:       image,
			Keywords:    keywords,
			Name:        name,
			Number:      number,
			Position:    i,
			Subcategory: subcategory,
			Unicode:     unicodes})
	})
}

// Open attempts to open all Emoji data from the emojipedia/emoji folder.
func Open() (*Emojipedia, error) {
	files, err := ioutil.ReadDir(directory.Emoji)
	if err != nil {
		return nil, err
	}
	emojipedia := New()
	for _, file := range files {
		name := strings.TrimSuffix(file.Name(), ".json")
		emoji, err := emoji.Open(name)
		if err != nil {
			return nil, err
		}
		emojipedia.Add(emoji)
	}
	return emojipedia, nil
}

// Remove deletes all Emoji data stored in the dependencies folder.
func Remove() error {
	return os.Remove(directory.Emoji)
}

type emojipedia interface {
	Add(emoji *emoji.Emoji) *Emojipedia
	Each(f func(key string, emoji *emoji.Emoji)) *Emojipedia
	Fetch(key string) *emoji.Emoji
	Get(key string) (*emoji.Emoji, bool)
	Has(key string) bool
	Keys() *slice.Slice
	Len() int
	Remove(key string) bool
	Values() *slice.Slice
}

// Emojipedia is a map-like struct with methods used to perform traversal and retrieval of emoji.Emoji pointers.
type Emojipedia struct {
	lexicon *lexicon.Lexicon
}

// Add method adds one emoji.Emoji to the Emojipedia using the emoji.Emoji.Name as the key reference.
func (pointer *Emojipedia) Add(emoji *emoji.Emoji) *Emojipedia {
	pointer.lexicon.Add(emoji.Name, emoji)
	return pointer
}

// Each method executes a provided function once for each emoji.Emoji pointer.
func (pointer *Emojipedia) Each(f func(key string, emoji *emoji.Emoji)) *Emojipedia {
	pointer.lexicon.Each(func(key string, i interface{}) {
		f(key, i.(*emoji.Emoji))
	})
	return pointer
}

// Fetch retrieves the emoji.Emoji pointer held by the argument key. Panics if key does not exist.
func (pointer *Emojipedia) Fetch(key string) *emoji.Emoji {
	property, _ := pointer.Get(key)
	return property
}

// Get returns the emoji.Emoji pointer held by the argument key and a boolean indicating if it was successfully retrieved.
// Panics if cannot convert to emoji.Emoji pointer.
func (pointer *Emojipedia) Get(key string) (*emoji.Emoji, bool) {
	property, ok := pointer.lexicon.Get(key)
	return property.(*emoji.Emoji), ok
}

// Has method checks that a given key exists in the Emojipedia.
func (pointer *Emojipedia) Has(key string) bool {
	return pointer.lexicon.Has(key)
}

// Keys method returns a slice.Slice of a given Emojipedia' own property names, in the same order as we get with a normal loop.
func (pointer *Emojipedia) Keys() *slice.Slice {
	slice := slice.New()
	pointer.lexicon.Each(func(key string, i interface{}) {
		slice.Append(key)
	})
	return slice
}

// Len method returns the number of elements in the Emojipedia.
func (pointer *Emojipedia) Len() int {
	return pointer.lexicon.Len()
}

// Remove method removes a entry from the Emojipedia if it exists. Returns a boolean to confirm if it succeeded.
func (pointer *Emojipedia) Remove(key string) bool {
	return pointer.lexicon.Remove(key)
}

// Values method returns a Slice of a given Emojipedia's own enumerable property values,
// in the same order as that provided by a for...in loop.
func (pointer *Emojipedia) Values() *slice.Slice {
	slice := slice.New()
	pointer.lexicon.Each(func(key string, i interface{}) {
		slice.Append(i)
	})
	return slice
}
