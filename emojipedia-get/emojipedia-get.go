package get

import (
	"fmt"
	"os"
	"runtime"
	"strings"

	"github.com/gellel/emojipedia/cli"
)

var programs = []interface{}{
	ByID,
	ByName,
	Category,
	Description,
	ID,
	Name,
	Subcategory}

func Main(args []string) {
	manifest := cli.NewManifest(runtime.Caller(0))
	program := cli.NewProgramFromManifest(manifest, programs)
	switch len(args) {
	case 0:
		panic(fmt.Errorf("%s", strings.Join(os.Args, ",")))
	case 1:
		fmt.Println(program.Use)
	default:
		switch strings.ToUpper(args[1]) {
		}
	}
}

func Category(name string) {}

func Description(name string) {}

func Get(options ...string) {}

func ByID(id int) {}

func ByName(name string) {}

func ID(name string) {}

func Name(id int) {}

func Subcategory(name string) {}
