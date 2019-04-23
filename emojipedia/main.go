package emojipedia

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"
	"net/http/httputil"
	"os"
	"path/filepath"
	"runtime"
	"strconv"
	"strings"

	"github.com/PuerkitoBio/goquery"
	"github.com/gellel/emojipedia/text"
)

const (
	// CategorizationFolder is the name for the stored categorization folder.
	CategorizationFolder string = "categorization"
	// EncyclopediaFolder is the name for the stored encyclopedia folder.
	EncyclopediaFolder string = "encyclopedia"
	// KeywordsFolder is the name for the stored keywords folder.
	KeywordsFolder string = "keywords"
	// SubcategorizationFolder is the name for the stored subcategorization folder.
	SubcategorizationFolder string = "subcategorization"
	// UnicodeFolder is the name for the stored unicode HTML folder.
	UnicodeFolder string = "unicode"
	// PackageFolder is the directory of the Emojipedia files.
	PackageFolder string = "emojipedia-storage"
)

const (
	// CategorizationFile is the name for the stored categorization JSON file.
	CategorizationFile string = (CategorizationFolder + ".json")
	// EncyclopediaFile is the name for the stored encyclopedia JSON file.
	EncyclopediaFile string = (EncyclopediaFolder + ".json")
	// KeywordsFile is the name for the stored keywords JSON file.
	KeywordsFile string = (KeywordsFolder + ".json")
	// SubcategorizationFile is the name for the stored subcategorization JSON file.
	SubcategorizationFile string = (SubcategorizationFolder + ".json")
	// UnicodeFile is the name for the stored unicode HTML file.
	UnicodeFile string = (UnicodeFolder + ".html")
)

const (
	// UnicodeOrgHref is the resource URL unicode information is sourced.
	UnicodeOrgHref string = "https://unicode.org/emoji/charts/emoji-list.html"
	// EmojipediaOrgHref is the resource URL emojipedia information is stored.
	EmojipediaOrgHref string = "https://emojipedia.org/"
)

const (
	// FileMode is the setting for which JSON files are stored.
	FileMode os.FileMode = 0644
)

var (
	_, b, _, _ = runtime.Caller(0)
	// Basepath is the directory of the package.
	Basepath = filepath.Dir(b)
	// Rootpath is the direct of the main package.
	Rootpath = filepath.Dir(Basepath)
	// Storagepath is the directory of the program files.
	Storagepath = filepath.Join(Rootpath, PackageFolder)
)

var (
	filenames = map[string]int{
		CategorizationFile:    1,
		EncyclopediaFile:      1,
		KeywordsFile:          1,
		SubcategorizationFile: 1}
)

// NewCategory makes an Emoji category.
func NewCategory(anchor string, emoji *Strings, href string, position int, name string, number int, subcategories *Strings) (category *Category) {
	category = &Category{anchor, emoji, href, position, name, number, subcategories}
	return category
}

// NewCategorization makes a new categorization from a slice of category.
func NewCategorization(categories ...*Category) (categorization *Categorization) {
	categorization = &Categorization{}
	for _, category := range categories {
		categorization.Assign(category)
	}
	return categorization
}

// NewCategorizationFromDocument makes a Emoji categorization from a GoQuery document.
func NewCategorizationFromDocument(document *goquery.Document) (categorization *Categorization) {
	var key string
	categorization = &Categorization{}
	document.Find("tr").Each(func(i int, selection *goquery.Selection) {
		selection.Find("th.bighead a").Each(func(j int, s *goquery.Selection) {
			var (
				anchor, _     = s.Attr("href")
				emoji         = &Strings{}
				href          = (UnicodeOrgHref + anchor)
				position      = i
				name          = text.Normalize(s.Text())
				number        = categorization.Len()
				subcategories = &Strings{}
				category      = NewCategory(anchor, emoji, href, position, name, number, subcategories)
			)
			categorization.Assign(category)
			key = category.Name
		})
		selection.Find("th.mediumhead a").Each(func(j int, s *goquery.Selection) {
			var (
				category, _ = categorization.Get(key)
				subcategory = text.Normalize(s.Text())
			)
			category.Subcategories.Push(subcategory)
		})
		selection.Find("td").Eq(3).Each(func(j int, s *goquery.Selection) {
			var (
				category, _ = categorization.Get(key)
				name        = text.Normalize(s.Text())
			)
			category.Emoji.Push(name)
		})
	})
	return categorization
}

// NewEmoji makes an Emoji.
func NewEmoji(anchor, category string, codes *Strings, description, href, img string, keywords *Strings, name string, number, position int, subcategory, unicode string) (emoji *Emoji) {
	emoji = &Emoji{anchor, category, codes, description, href, img, keywords, name, number, position, subcategory, unicode}
	return emoji
}

// NewEncyclopedia makes an Emoji encyclopedia.
func NewEncyclopedia(emojis ...*Emoji) (encyclopedia *Encyclopedia) {
	encyclopedia = &Encyclopedia{}
	for _, emoji := range emojis {
		encyclopedia.Assign(emoji)
	}
	return encyclopedia
}

// NewEncyclopediaFromDocument makes an Ecyclopedia from a GoQuery document.
func NewEncyclopediaFromDocument(document *goquery.Document) (encyclopedia *Encyclopedia) {
	var category, subcategory string
	encyclopedia = &Encyclopedia{}
	document.Find("tr").Each(func(i int, selection *goquery.Selection) {
		var (
			anchor   string
			codes    = &Strings{}
			keywords = &Strings{}
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
				codes.Push(strings.TrimSpace(substring))
			}
		})
		selection.Find("td.andr a").Each(func(j int, s *goquery.Selection) {
			anchor, _ = s.Attr("href")
		})
		selection.Find("td.name").First().Each(func(j int, s *goquery.Selection) {
			name = text.Normalize(s.Text())
		})
		selection.Find("td.name").Last().Each(func(j int, s *goquery.Selection) {
			for _, substring := range strings.Split(s.Text(), "|") {
				keywords.Push(strings.TrimSpace(substring))
			}
		})
		if len(name) == 0 {
			return
		}
		codes.Each(func(i int, code string) {
			replacement := "000"
			if len(code) == 6 {
				replacement = "0000"
			}
			unicodes = unicodes + strings.Replace(code, "+", replacement, 1)
		})
		unicodes = strings.Replace(strings.ToLower(unicodes), "u", "\\U", -1)
		emoji := &Emoji{
			Anchor:      anchor,
			Category:    category,
			Codes:       codes,
			Href:        (UnicodeOrgHref + anchor),
			Image:       "NIL",
			Keywords:    keywords,
			Name:        name,
			Number:      number,
			Position:    i,
			Subcategory: subcategory,
			Unicode:     unicodes}
		encyclopedia.Assign(emoji)
	})
	return encyclopedia
}

// NewKeywordsFromDocument makes a Emoji keywords from a GoQuery document.
func NewKeywordsFromDocument(document *goquery.Document) (keywords *Keywords) {
	keywords = &Keywords{}
	document.Find("tr").Each(func(i int, selection *goquery.Selection) {
		s := selection.Find("td.name")
		name := strings.TrimSpace(s.First().Text())
		keys := strings.TrimSpace(s.Last().Text())
		if len(name) == 0 {
			return
		}
		name = text.Normalize(name)
		for _, key := range strings.Split(keys, "|") {
			key = strings.TrimSpace(key)
			key = text.Normalize(key)
			keywords.Append(key, name)
		}
	})
	return keywords
}

// NewSubcategory makes an Emoji subcategory.
func NewSubcategory(anchor string, category string, emoji *Strings, href string, position int, name string, number int) (subcategory *Subcategory) {
	subcategory = &Subcategory{anchor, category, emoji, href, position, name, number}
	return subcategory
}

// NewSubcategorization makes a new subcategorization from slice of subcategory.
func NewSubcategorization(subcategories ...*Subcategory) (subcategorization *Subcategorization) {
	subcategorization = &Subcategorization{}
	for _, subcategory := range subcategories {
		subcategorization.Assign(subcategory)
	}
	return subcategorization
}

// NewSubcategorizationFromDocument makes a Emoji subcategorization from a GoQuery document.
func NewSubcategorizationFromDocument(document *goquery.Document) (subcategorization *Subcategorization) {
	var key, category string
	subcategorization = &Subcategorization{}
	document.Find("tr").Each(func(i int, selection *goquery.Selection) {
		selection.Find("th.bighead a").Each(func(j int, s *goquery.Selection) {
			category = text.Normalize(s.Text())
		})
		selection.Find("th.mediumhead a").Each(func(j int, s *goquery.Selection) {
			var (
				anchor, _   = s.Attr("href")
				emoji       = &Strings{}
				href        = (UnicodeOrgHref + anchor)
				position    = i
				name        = text.Normalize(s.Text())
				number      = subcategorization.Len()
				subcategory = NewSubcategory(anchor, category, emoji, href, position, name, number)
			)
			subcategorization.Assign(subcategory)
			key = subcategory.Name
		})
		selection.Find("td").Eq(3).Each(func(j int, s *goquery.Selection) {
			var (
				name           = text.Normalize(s.Text())
				subcategory, _ = subcategorization.Get(key)
			)
			subcategory.Emoji.Push(name)
		})
	})
	return subcategorization
}

// NewUnicodeOrgHTMLDump requests the unicode.org data from the net and dumps the HTTP response.
func NewUnicodeOrgHTMLDump() (dump []byte, ok bool) {
	resp, err := http.Get(UnicodeOrgHref)
	ok = (err == nil)
	if ok != true {
		return nil, ok
	}
	defer resp.Body.Close()
	ok = (resp.StatusCode == 200)
	if ok != true {
		return nil, ok
	}
	dump, err = httputil.DumpResponse(resp, true)
	ok = (err == nil)
	if ok != true {
		return nil, ok
	}
	return dump, ok
}

// OpenEmojipediaFile opens a file made by the Emojipedia program.
func OpenEmojipediaFile(name string) (bytes []byte, ok bool) {
	_, ok = filenames[name]
	if ok != true {
		return nil, ok
	}
	filename := filepath.Join(Storagepath, name)
	reader, err := os.Open(filename)
	ok = (err == nil)
	if ok != true {
		return nil, ok
	}
	bytes, err = ioutil.ReadAll(reader)
	reader.Close()
	ok = (err == nil)
	if ok != true {
		return nil, ok
	}
	return bytes, ok
}

// OpenCategorizationFromFile opens a stored categorization file.
func OpenCategorizationFromFile() (categorization *Categorization, ok bool) {
	bytes, ok := OpenEmojipediaFile(CategorizationFile)
	if ok != true {
		return nil, ok
	}
	categorization = &Categorization{}
	ok = (json.Unmarshal(bytes, categorization) == nil)
	if ok != true {
		return nil, ok
	}
	return categorization, ok
}

// OpenEmojiFromFile opens a stored Emoji.
func OpenEmojiFromFile(name string) (emoji *Emoji, ok bool) {
	filename := filepath.Join(Storagepath, EncyclopediaFolder, (name + ".json"))
	reader, err := os.Open(filename)
	ok = (err == nil)
	if ok != true {
		return nil, ok
	}
	bytes, err := ioutil.ReadAll(reader)
	reader.Close()
	ok = (err == nil)
	if ok != true {
		return nil, ok
	}
	emoji = &Emoji{}
	ok = (json.Unmarshal(bytes, emoji) == nil)
	if ok != true {
		return nil, ok
	}
	return emoji, ok
}

// OpenEncyclopediaFromFile opens a stored encyclopedia file.
func OpenEncyclopediaFromFile() (encyclopedia *Encyclopedia, ok bool) {
	bytes, ok := OpenEmojipediaFile(EncyclopediaFile)
	if ok != true {
		return nil, ok
	}
	encyclopedia = &Encyclopedia{}
	ok = (json.Unmarshal(bytes, encyclopedia) == nil)
	if ok != true {
		return nil, ok
	}
	return encyclopedia, ok
}

// OpenSubcategorizationFromFile opens a stored subcategorization file.
func OpenSubcategorizationFromFile() (subcategorization *Subcategorization, ok bool) {
	bytes, ok := OpenEmojipediaFile(SubcategorizationFile)
	if ok != true {
		return nil, ok
	}
	subcategorization = &Subcategorization{}
	ok = (json.Unmarshal(bytes, subcategorization) == nil)
	if ok != true {
		return nil, ok
	}
	return subcategorization, ok
}

// RemoveEmojipediaFile removes a file made by the Emojipedia program.
func RemoveEmojipediaFile(name string) (ok bool) {
	_, ok = filenames[name]
	if ok != true {
		return ok
	}
	filename := filepath.Join(Storagepath, name)
	ok = (os.Remove(filename) == nil)
	return ok
}

// RemoveCategorizationJSON removes stored categorization JSON.
func RemoveCategorizationJSON() (ok bool) {
	filename := filepath.Join(Storagepath, CategorizationFolder)
	ok = (os.RemoveAll(filename) == nil)
	if ok != true {
		return ok
	}
	ok = (os.Mkdir(filename, FileMode) == nil)
	return ok
}

// RemoveEmojiJSON removes a stored emoji JSON.
func RemoveEmojiJSON(name string) (ok bool) {
	filename := filepath.Join(Storagepath, EncyclopediaFolder, (name + ".json"))
	fmt.Println(filename)
	err := os.Remove(filename)
	fmt.Println(err)
	ok = (err == nil)
	return ok
}

// RemoveEncyclopediaJSON removes stored encyclopedia JSON.
func RemoveEncyclopediaJSON() (ok bool) {
	filename := filepath.Join(Storagepath, EncyclopediaFolder)
	ok = (os.RemoveAll(filename) == nil)
	if ok != true {
		return ok
	}
	ok = (os.Mkdir(filename, FileMode) == nil)
	return ok
}

// RemoveSubcategorizationJSON removes stored subcategorization JSON.
func RemoveSubcategorizationJSON() (ok bool) {
	filename := filepath.Join(Storagepath, SubcategorizationFolder)
	ok = (os.RemoveAll(filename) == nil)
	if ok != true {
		return ok
	}
	ok = (os.Mkdir(filename, FileMode) == nil)
	return ok
}

// StoreCategorizationAsJSON stores a categorization as JSON.
func StoreCategorizationAsJSON(categorization *Categorization) (ok bool) {
	bytes, err := json.Marshal(categorization)
	ok = (err == nil)
	if ok != true {
		return ok
	}
	filename := filepath.Join(Storagepath, CategorizationFolder, CategorizationFile)
	ok = (ioutil.WriteFile(filename, bytes, FileMode) == nil)
	return ok
}

// StoreKeywordsAsJSON stores an encyclopedia as JSON.
func StoreKeywordsAsJSON(keywords *Keywords) (ok bool) {
	bytes, err := json.Marshal(keywords)
	ok = (err == nil)
	if ok != true {
		return ok
	}
	filename := filepath.Join(Storagepath, KeywordsFolder, KeywordsFile)
	ok = (ioutil.WriteFile(filename, bytes, FileMode) == nil)
	return ok
}

// StoreEmojiAsJSON stores a unique Emoji as JSON.
func StoreEmojiAsJSON(emoji *Emoji) (ok bool) {
	bytes, err := json.Marshal(emoji)
	ok = (err == nil)
	if ok != true {
		return ok
	}
	filename := filepath.Join(Storagepath, EncyclopediaFolder, (emoji.Name + ".json"))
	ok = (ioutil.WriteFile(filename, bytes, FileMode) == nil)
	return ok
}

// StoreEncyclopediaAsJSON stores an encyclopedia as JSON.
func StoreEncyclopediaAsJSON(encyclopedia *Encyclopedia) (ok bool) {
	encyclopedia.Each(func(name string, emoji *Emoji) {
		ok = StoreEmojiAsJSON(emoji)
	})
	return ok
}

// StoreSubcategorizationAsJSON stores a categorization as JSON.
func StoreSubcategorizationAsJSON(subcategorization *Subcategorization) (ok bool) {
	bytes, err := json.Marshal(subcategorization)
	ok = (err == nil)
	if ok != true {
		return ok
	}
	filename := filepath.Join(Storagepath, SubcategorizationFolder, SubcategorizationFile)
	ok = (ioutil.WriteFile(filename, bytes, FileMode) == nil)
	return ok
}

// StoreUnicodeOrgFileAsHTML stores a unicode HTML file requested from the internet.
func StoreUnicodeOrgFileAsHTML(dump *[]byte) (ok bool) {
	filename := filepath.Join(Storagepath, UnicodeFolder, UnicodeFile)
	ok = (ioutil.WriteFile(filename, *dump, FileMode) == nil)
	return ok
}
