package emoji

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"

	"github.com/gellel/emojipedia/directory"
	"github.com/gellel/emojipedia/slice"
)

var _ emoji = (*Emoji)(nil)

// New instantiates a new empty Emoji pointer.
func New() *Emoji {
	return &Emoji{
		Codes:    &slice.Slice{},
		Keywords: &slice.Slice{}}
}

// NewEmoji creates a new Emoji pointer, requiring all struct features as arguments.
func NewEmoji(anchor, category, href, image, name, subcategory, unicode string, number, position int, codes, keywords *slice.Slice) *Emoji {
	return &Emoji{
		Anchor:      anchor,
		Category:    category,
		Codes:       codes,
		Description: "NIL",
		Href:        href,
		Image:       image,
		Keywords:    keywords,
		Name:        name,
		Number:      number,
		Position:    position,
		Subcategory: subcategory,
		Unicode:     unicode}
}

// Get attempts to open a Category from the emojipedia/emoji folder, but panics if an error occurs.
func Get(name string) *Emoji {
	emoji, err := Open(name)
	if err != nil {
		panic(err)
	}
	return emoji
}

// Open attempts to open a Emoji from the emojipedia/emoji folder.
func Open(name string) (*Emoji, error) {
	filepath := filepath.Join(directory.Emoji, fmt.Sprintf("%s.json", name))
	reader, err := os.Open(filepath)
	if err != nil {
		return nil, err
	}
	content, err := ioutil.ReadAll(reader)
	defer reader.Close()
	if err != nil {
		return nil, err
	}
	emoji := &Emoji{}
	err = json.Unmarshal(content, emoji)
	if err != nil {
		return nil, err
	}
	return emoji, nil
}

func Parse(content *[]byte) (*Emoji, error) {
	category := &Emoji{}
	err := json.Unmarshal(*content, category)
	if err != nil {
		return nil, err
	}
	return category, nil
}

func Read(name string) (*[]byte, error) {
	filepath := filepath.Join(directory.Emoji, fmt.Sprintf("%s.json", name))
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

// Remove deletes the Emoji data stored in the dependencies folder.
func Remove(name string) error {
	return os.Remove(filepath.Join(directory.Emoji, fmt.Sprintf("%s.json", name)))
}

// Write stores and Emoji pointer to the dependencies folder.
func Write(emoji *Emoji) error {
	err := os.MkdirAll(directory.Emoji,  os.ModePerm)
	if err != nil {
		return err
	}
	content, err := json.Marshal(emoji)
	if err != nil {
		return err
	}
	filepath := filepath.Join(directory.Emoji, fmt.Sprintf("%s.json", emoji.Name))
	return ioutil.WriteFile(filepath, content,  os.ModePerm)
}

type emoji interface {
	SetAnchor(anchor string) *Emoji
	SetCategory(category string) *Emoji
	SetCodes(codes *slice.Slice) *Emoji
	SetDescription(description string) *Emoji
	SetHref(href string) *Emoji
	SetImage(image string) *Emoji
	SetKeywords(keywords *slice.Slice) *Emoji
	SetName(name string) *Emoji
	SetNumber(number int) *Emoji
	SetPosition(position int) *Emoji
	SetSubcategory(subcategory string) *Emoji
	SetUnicode(unicode string) *Emoji
}

// Emoji stores the contents about an emoji scraped from the unicode consortium.
type Emoji struct {
	Anchor      string       `json:"anchor"`
	Category    string       `json:"category"`
	Codes       *slice.Slice `json:"codes"`
	Description string       `json:"description"`
	Href        string       `json:"href"`
	Image       string       `json:"img"`
	Keywords    *slice.Slice `json:"keywords"`
	Name        string       `json:"name"`
	Number      int          `json:"number"`
	Position    int          `json:"position"`
	Subcategory string       `json:"subcategory"`
	Unicode     string       `json:"unicode"`
}

// SetAnchor sets the Emoji.Anchor property.
func (pointer *Emoji) SetAnchor(anchor string) *Emoji {
	pointer.Anchor = anchor
	return pointer
}

// SetCategory sets the Emoji.Category property.
func (pointer *Emoji) SetCategory(category string) *Emoji {
	pointer.Category = category
	return pointer
}

// SetCodes sets the Emoji.Codes property.
func (pointer *Emoji) SetCodes(codes *slice.Slice) *Emoji {
	pointer.Codes = codes
	return pointer
}

// SetDescription sets the Emoji.Description property.
func (pointer *Emoji) SetDescription(description string) *Emoji {
	pointer.Description = description
	return pointer
}

// SetHref sets the Emoji.Href property.
func (pointer *Emoji) SetHref(href string) *Emoji {
	pointer.Href = href
	return pointer
}

// SetImage sets the Emoji.Image property.
func (pointer *Emoji) SetImage(image string) *Emoji {
	pointer.Image = image
	return pointer
}

// SetKeywords sets the Emoji.Keywords property.
func (pointer *Emoji) SetKeywords(keywords *slice.Slice) *Emoji {
	pointer.Keywords = keywords
	return pointer
}

// SetName sets the Emoji.Name property.
func (pointer *Emoji) SetName(name string) *Emoji {
	pointer.Name = name
	return pointer
}

// SetNumber sets the Emoji.Number property.
func (pointer *Emoji) SetNumber(number int) *Emoji {
	pointer.Number = number
	return pointer
}

// SetPosition sets the Emoji.Position property.
func (pointer *Emoji) SetPosition(position int) *Emoji {
	pointer.Position = position
	return pointer
}

// SetSubcategory sets the Emoji.Subcategory property.
func (pointer *Emoji) SetSubcategory(subcategory string) *Emoji {
	pointer.Subcategory = subcategory
	return pointer
}

// SetUnicode sets the Emoji.Unicode property.
func (pointer *Emoji) SetUnicode(unicode string) *Emoji {
	pointer.Unicode = unicode
	return pointer
}
