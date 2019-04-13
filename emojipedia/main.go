package emojipedia

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
	"path/filepath"
	"reflect"
	"runtime"
	"strconv"
	"strings"
	"text/tabwriter"
	"unicode"

	"github.com/PuerkitoBio/goquery"

	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

var replacements = []string{
	")", "",
	"(", "",
	"\"", "",
	".", "",
	"'", "",
	":", "",
	",", "",
	"⊛", "",
	"“", "",
	"”", "",
	"_", "-",
	"&", "and",
	" ", "-"}

var replacer = strings.NewReplacer(replacements...)

func NewEmoji(columns []string) *Emoji {
	var unicodes string
	no, columns := columns[0], columns[1:]
	number, _ := strconv.Atoi(no)
	sample, columns := columns[0], columns[1:]
	codes, columns := columns[0], columns[1:]
	for _, code := range strings.Fields(codes) {
		replacement := "000"
		if len(code) == 6 {
			replacement = "0000"
		}
		unicodes = unicodes + strings.Replace(code, "+", replacement, 1)
	}
	unicodes = unicodes + strings.Replace(unicodes, "U", "\\U", 1)
	name, columns := columns[0], columns[1:]
	name = Normalize(name)
	keywords, columns := columns[0], columns[1:]
	keywords = strings.Replace(keywords, "|", "", -1)
	keys := strings.Fields(keywords)
	category, columns := columns[0], columns[1:]
	subcategory, columns := columns[0], columns[1:]
	return &Emoji{
		Category:    category,
		Code:        codes,
		Keywords:    keys,
		Name:        name,
		Number:      number,
		Sample:      sample,
		Subcategory: subcategory,
		Unicode:     unicodes}
}

func NewEmojidex() *Emojidex {
	return &Emojidex{}
}

func NewEmojipedia() *Emojipedia {
	return &Emojipedia{
		Emojidex:     NewEmojidex(),
		Encyclopedia: NewEncyclopedia()}
}

func NewEmojipediaFromDocument(document *goquery.Document) *Emojipedia {
	var category, subcategory string
	emojidex := NewEmojidex()
	encyclopedia := NewEncyclopedia()
	document.Find("tr").Each(func(i int, selection *goquery.Selection) {
		var columns []string
		selection.Find("th.bighead").Each(func(j int, s *goquery.Selection) {
			category = replacer.Replace(encyclopedia.Categories.Set(strings.TrimSpace(s.Text())))
		})
		selection.Find("th.mediumhead").Each(func(j int, s *goquery.Selection) {
			subcategory = replacer.Replace(encyclopedia.Subcategories.Set(strings.TrimSpace(s.Text())))
		})
		selection.Find("td").Each(func(j int, s *goquery.Selection) {
			columns = append(columns, strings.TrimSpace(s.Text()))
		})
		if len(columns) != 5 {
			return
		}
		columns = append(columns, category, subcategory)
		emoji := NewEmoji(columns)
		for _, keyword := range emoji.Keywords {
			encyclopedia.Keywords.Add(keyword, emoji.Name)
		}
		encyclopedia.Categories.Add(category, emoji.Name)
		encyclopedia.Subcategories.Add(subcategory, emoji.Name)
		encyclopedia.Numeric.Add(emoji.Number, emoji.Name)
		emojidex.Add(emoji.Name, emoji)
	})
	return &Emojipedia{
		Emojidex:     emojidex,
		Encyclopedia: encyclopedia}
}

func NewEncyclopedia() *Encyclopedia {
	return &Encyclopedia{
		Categories:    &Associative{},
		Subcategories: &Associative{},
		Keywords:      &Associative{},
		Numeric:       &Set{}}
}

func NewFilepath(filename string) string {
	_, file, _, ok := runtime.Caller(0)
	if ok != true {
		panic(file)
	}
	dir := filepath.Dir(file)
	parent := filepath.Dir(dir)
	filename = strings.Replace(filename, ".json", "", -1)
	return filepath.Join(parent, (filename + ".json"))
}

func Normalize(value string) string {
	f := func(r rune) bool {
		return unicode.Is(unicode.Mn, r)
	}
	t := transform.Chain(norm.NFD, transform.RemoveFunc(f), norm.NFC)
	result, _, _ := transform.String(t, value)
	result = strings.TrimSpace(result)
	result = strings.Replace(replacer.Replace(result), " ", "-", -1)
	result = strings.Replace(strings.ToLower(result), "_", "-", -1)
	return result
}

func MarshallAssociative(filename string, associative *Associative) string {
	filename = NewFilepath(filename)
	contents, err := json.Marshal(associative)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(filename, contents, 0644)
	if err != nil {
		panic(err)
	}
	return filename
}

func MarshallEmojidex(filename string, emojidex *Emojidex) string {
	filename = NewFilepath(filename)
	contents, err := json.Marshal(emojidex)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(filename, contents, 0644)
	if err != nil {
		panic(err)
	}
	return filename
}

func MarshallEmojipedia(emojipedia *Emojipedia) {
	//filename = NewFilepath(filename)
}

func MarshallSet(filename string, set *Set) string {
	filename = NewFilepath(filename)
	contents, err := json.Marshal(set)
	if err != nil {
		panic(err)
	}
	err = ioutil.WriteFile(filename, contents, 0644)
	if err != nil {
		panic(err)
	}
	return filename
}

func PrintEmoji(emoji *Emoji) {
	reflection := reflect.ValueOf(emoji).Elem()
	values := []string{}
	keys := []string{}
	for i := 0; i < reflection.NumField(); i++ {
		in := reflection.Field(i).Interface()
		switch in.(type) {
		case []string:
			values = append(values, strings.Join((in.([]string)), ","))
		case int:
			values = append(values, strconv.Itoa((in.(int))))
		default:
			values = append(values, (in.(string)))
		}
		keys = append(keys, reflection.Type().Field(i).Name)
	}
	writer := new(tabwriter.Writer)
	writer.Init(os.Stdout, 0, 0, 0, ' ', tabwriter.Debug|tabwriter.AlignRight)
	fmt.Fprintln(writer, strings.Join(keys, "\t")+"\t")
	fmt.Fprintln(writer, strings.Join(values, "\t")+"\t")
	writer.Flush()
}

func PrintEmojidex(emojidex *Emojidex) {
	for _, emoji := range *emojidex {
		PrintEmoji(emoji)
	}
}

func UnmarshallEmojidex() *Emojidex {
	filename := NewFilepath("emoji")
	jsonFile, err := os.Open(filename)
	emojidex := &Emojidex{}
	byteValue, err := ioutil.ReadAll(jsonFile)
	if err != nil {
		panic(err)
	}
	err = json.Unmarshal(byteValue, emojidex)
	if err != nil {
		panic(err)
	}
	return emojidex
}
