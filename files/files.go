package files

import (
	"fmt"
	"os"
	"path/filepath"
	"runtime"
	"strings"

	cli "github.com/gellel/emojipedia/cli"
	"github.com/gellel/emojipedia/web"
)

var Options = []interface{}{
	dependencies,
	hasFile}

func checkFile(name string) {
	_, file, _, _ := runtime.Caller(0)
	dir := filepath.Dir(file)
	parent := filepath.Dir(dir)
	filename := filepath.Join(parent, name)
	info, err := os.Stat(filename)
	message := ""
	if err != nil {
		message = fmt.Sprintf("does not have %s on disk. checked directory %s", name, parent)
	} else {
		message = fmt.Sprintf("file info for %s. size: %v", name, info.Size())
	}
	fmt.Println(cli.WrapDescription(message))
}

func dependencies() {
	document := web.HttpToUnicodeOrg()
	emojidex, encyclopedia := web.HttpToEmojiPackage(document)
	fmt.Println(emojidex)
	fmt.Println(encyclopedia.Numeric)
}

func hasFile(name string) {
	name = strings.Replace(strings.ToLower(name), ".json", "", -1)
	names := []string{"categories", "keywords", "names", "numbers", "subcategories"}
	for _, n := range names {
		if n == name {
			checkFile(fmt.Sprintf("%s.json", n))
			break
		}
	}
}

func Main(options []string) {
	substrings := strings.Split(options[0], "=")
	switch strings.ToLower(substrings[0]) {
	case "has-file":
		switch len(substrings) {
		case 1:
			hasFile(options[1])
		default:
			hasFile(substrings[1])
		}
	case "--dependencies":
		dependencies()
	}
}
