package directory

import (
	"fmt"
	"path/filepath"
	"runtime"
)

const (
	category    string = "category"
	emoji       string = "emoji"
	keywords    string = "keywords"
	subcategory string = "subcategory"
	unicode     string = "unicode"
)

var (
	_, file, _, _ = runtime.Caller(0)
	rootpath      = filepath.Dir(filepath.Dir(file))
	storagepath   = filepath.Join(rootpath, fmt.Sprintf(".%s", "emojipedia"))
)

var (
	Category    = filepath.Join(storagepath, category)
	Emoji       = filepath.Join(storagepath, emoji)
	Keywords    = filepath.Join(storagepath, keywords)
	Subcategory = filepath.Join(storagepath, subcategory)
	Unicode     = filepath.Join(storagepath, unicode)
)
