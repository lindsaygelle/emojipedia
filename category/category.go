package category

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"text/tabwriter"

	"github.com/gellel/emojipedia/directory"
	"github.com/gellel/emojipedia/slice"
)

var _ category = (*Category)(nil)

var (
	tabs = []interface{}{
		"Name",
		"\t",
		"Number",
		"\t",
		"Position",
		"\t",
		"Href",
		"\t",
		"Subcategories",
		"\t",
		"Emoji"}
)

// New instantiates a new empty Category pointer.
func New() *Category {
	return &Category{
		Emoji:         &slice.Slice{},
		Subcategories: &slice.Slice{}}
}

// NewCategory creates a new Category pointer, requiring all struct features as arguments.
func NewCategory(anchor, href, name string, number, position int, emoji, subcategories *slice.Slice) *Category {
	return &Category{
		Anchor:        anchor,
		Emoji:         emoji,
		Href:          href,
		Name:          name,
		Number:        number,
		Position:      position,
		Subcategories: subcategories}
}

func Detail(content *[]byte) {
	category, err := Parse(content)
	if err != nil {
		panic(err)
	}
	fields := []interface{}{
		category.Name, "\t",
		category.Number, "\t",
		category.Position, "\t",
		category.Href, "\t",
		category.Subcategories.Len(), "\t",
		category.Emoji.Len()}

	w := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 0, '\t', 0)
	fmt.Fprintln(w, tabs...)
	fmt.Fprintln(w, fields...)
	w.Flush()
}

// Get attempts to open a Category from the emojipedia/categories folder, but panics if an error occurs.
func Get(name string) *Category {
	category, err := Open(name)
	if err != nil {
		panic(err)
	}
	return category
}

// List formats an interface representation of an Category pointer for os.Stdout.
func List(writer *tabwriter.Writer, i interface{}) {
	category := i.(*Category)
	fields := []interface{}{
		category.Name, "\t",
		category.Number, "\t",
		category.Subcategories.Len(), "\t",
		category.Emoji.Len()}

	fmt.Fprintln(writer, fields...)
}

// Open attempts to open a Category from the emojipedia/categories folder.
func Open(name string) (*Category, error) {
	content, err := Read(name)
	if err != nil {
		return nil, err
	}
	return Parse(content)
}

func Parse(content *[]byte) (*Category, error) {
	category := &Category{}
	err := json.Unmarshal(*content, category)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func Read(name string) (*[]byte, error) {
	filepath := filepath.Join(directory.Category, fmt.Sprintf("%s.json", name))
	reader, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadAll(reader)
	defer reader.Close()
	if err != nil {
		return nil, err
	}
	return &content, nil
}

// Remove deletes the Category data stored in the dependencies folder.
func Remove(name string) error {
	return os.Remove(filepath.Join(directory.Category, fmt.Sprintf("%s.json", name)))
}

// Write stores and Category pointer to the dependencies folder.
func Write(category *Category) error {
	err := os.MkdirAll(directory.Category, 0644)
	if err != nil {
		return err
	}
	content, err := json.Marshal(category)
	if err != nil {
		return err
	}
	filepath := filepath.Join(directory.Category, fmt.Sprintf("%s.json", category.Name))
	return ioutil.WriteFile(filepath, content, 0644)
}

type category interface {
	SetAnchor(anchor string) *Category
	SetEmoji(category *slice.Slice) *Category
	SetHref(href string) *Category
	SetName(name string) *Category
	SetNumber(number int) *Category
	SetPosition(position int) *Category
	SetSubcategories(subcategories *slice.Slice) *Category
}

// Category stores the categorical superset of the emoji data.
type Category struct {
	Anchor        string       `json:"anchor"`
	Emoji         *slice.Slice `json:"emoji"`
	Href          string       `json:"href"`
	Name          string       `json:"name"`
	Number        int          `json:"number"`
	Position      int          `json:"position"`
	Subcategories *slice.Slice `json:"subcategories"`
}

// SetAnchor sets the Category.Anchor property.
func (pointer *Category) SetAnchor(anchor string) *Category {
	pointer.Anchor = anchor
	return pointer
}

// SetEmoji sets the Category.Emoji property.
func (pointer *Category) SetEmoji(emoji *slice.Slice) *Category {
	pointer.Emoji = emoji
	return pointer
}

// SetHref sets the Category.Href property.
func (pointer *Category) SetHref(href string) *Category {
	pointer.Href = href
	return pointer
}

// SetName sets the Category.Name property.
func (pointer *Category) SetName(name string) *Category {
	pointer.Name = name
	return pointer
}

// SetNumber sets the Category.Number property.
func (pointer *Category) SetNumber(number int) *Category {
	pointer.Number = number
	return pointer
}

// SetPosition sets the Category.Position property.
func (pointer *Category) SetPosition(position int) *Category {
	pointer.Position = position
	return pointer
}

// SetSubcategories sets the Category.Subcategories property.
func (pointer *Category) SetSubcategories(subcategories *slice.Slice) *Category {
	pointer.Subcategories = subcategories
	return pointer
}
