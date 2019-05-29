package subcategory

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"strings"
	"text/tabwriter"

	"github.com/gellel/emojipedia/directory"
	"github.com/gellel/emojipedia/slice"
)

var _ subcategory = (*Subcategory)(nil)

var (
	tabs = []string{
		"Category",
		"Emoji",
		"Name",
		"Number",
		"Position"}
)

// New instantiates a new empty Subcategory pointer.
func New() *Subcategory {
	return &Subcategory{Emoji: &slice.Slice{}}
}

// NewSubcategory creates a new Subcategory pointer, requiring all struct features as arguments.
func NewSubcategory(anchor, category, href, name string, number, position int, emoji *slice.Slice) *Subcategory {
	return &Subcategory{
		Anchor:   anchor,
		Category: category,
		Emoji:    emoji,
		Href:     href,
		Name:     name,
		Number:   number,
		Position: position}
}

func Detail(content *[]byte) {
	subcategory, err := Parse(content)
	if err != nil {
		panic(err)
	}
	fields := []string{
		subcategory.Category,
		fmt.Sprintf("%v", subcategory.Emoji.Len()),
		subcategory.Name,
		fmt.Sprintf("%v", subcategory.Number),
		fmt.Sprintf("%v", subcategory.Position)}
	w := new(tabwriter.Writer).Init(os.Stdout, 0, 8, 1, '\t', 0)
	fmt.Fprintln(w, strings.Join(tabs, "\t"))
	fmt.Fprintln(w, strings.Join(fields, "\t"))
	w.Flush()
}

// Get attempts to open a Subcategory from the emojipedia/subcategories folder, but panics if an error occurs.
func Get(name string) *Subcategory {
	subcategory, err := Open(name)
	if err != nil {
		panic(err)
	}
	return subcategory
}

// List formats an interface representation of an Category pointer for os.Stdout.
func List(writer *tabwriter.Writer, i interface{}) {
	subcategory := i.(*Subcategory)
	fields := []interface{}{
		subcategory.Name, "\t",
		subcategory.Number, "\t",
		subcategory.Category, "\t",
		subcategory.Emoji.Len()}
	fmt.Fprintln(writer, fields...)
}

// Open attempts to open a Subcategory from the emojipedia/subcategories folder.
func Open(name string) (*Subcategory, error) {
	filepath := filepath.Join(directory.Subcategory, fmt.Sprintf("%s.json", name))
	reader, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadAll(reader)
	defer reader.Close()
	if err != nil {
		return nil, err
	}
	subcategory := &Subcategory{}
	err = json.Unmarshal(content, subcategory)
	if err != nil {
		return nil, err
	}
	return subcategory, nil
}

func Parse(content *[]byte) (*Subcategory, error) {
	category := &Subcategory{}
	err := json.Unmarshal(*content, category)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func Read(name string) (*[]byte, error) {
	filepath := filepath.Join(directory.Subcategory, fmt.Sprintf("%s.json", name))
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

// Remove deletes the Subcategory data stored in the dependencies folder.
func Remove(name string) error {
	return os.Remove(filepath.Join(directory.Subcategory, fmt.Sprintf("%s.json", name)))
}

// Write stores and Subcategory pointer to the dependencies folder.
func Write(subcategory *Subcategory) error {
	err := os.MkdirAll(directory.Subcategory, 0644)
	if err != nil {
		return err
	}
	content, err := json.Marshal(subcategory)
	if err != nil {
		return err
	}
	filepath := filepath.Join(directory.Subcategory, fmt.Sprintf("%s.json", subcategory.Name))
	return ioutil.WriteFile(filepath, content, 0644)
}

type subcategory interface {
	SetAnchor(anchor string) *Subcategory
	SetCategory(category string) *Subcategory
	SetEmoji(emoji *slice.Slice) *Subcategory
	SetHref(href string) *Subcategory
	SetName(name string) *Subcategory
	SetNumber(number int) *Subcategory
	SetPosition(position int) *Subcategory
}

type Subcategory struct {
	Anchor   string       `json:"anchor"`
	Category string       `json:"category"`
	Emoji    *slice.Slice `json:"emoji"`
	Href     string       `json:"href"`
	Name     string       `json:"name"`
	Number   int          `json:"number"`
	Position int          `json:"position"`
}

// SetAnchor sets the Subcategory.Anchor property.
func (pointer *Subcategory) SetAnchor(anchor string) *Subcategory {
	pointer.Anchor = anchor
	return pointer
}

// SetCategory sets the Subcategory.Category property.
func (pointer *Subcategory) SetCategory(category string) *Subcategory {
	pointer.Category = category
	return pointer
}

// SetEmoji sets the Subcategory.Emoji property.
func (pointer *Subcategory) SetEmoji(emoji *slice.Slice) *Subcategory {
	pointer.Emoji = emoji
	return pointer
}

// SetHref sets the Subcategory.Href property.
func (pointer *Subcategory) SetHref(href string) *Subcategory {
	pointer.Href = href
	return pointer
}

// SetName sets the Subcategory.Name property.
func (pointer *Subcategory) SetName(name string) *Subcategory {
	pointer.Name = name
	return pointer
}

// SetNumber sets the Subcategory.Number property.
func (pointer *Subcategory) SetNumber(number int) *Subcategory {
	pointer.Number = number
	return pointer
}

// SetPosition sets the Subcategory.Position property.
func (pointer *Subcategory) SetPosition(position int) *Subcategory {
	pointer.Position = position
	return pointer
}
