package get

import (
	"fmt"
	"os"
	"runtime"
	"strconv"
	"strings"

	"github.com/gellel/emojipedia/cli"
)

func category(name string) int {
	fmt.Println(name)
	return 0
}

func id(name string) int {
	fmt.Println(name)
	return 0
}

func name(ID int) int {
	fmt.Println(ID)
	return 0
}

func Get(options ...string) int {
	opts := []interface{}{category, id, name}
	manifest := cli.NewManifest(runtime.Caller(0))
	program := cli.NewProgramFromManifest(manifest, opts)
	switch len(options) {
	case 0:
		panic(fmt.Errorf("%s", strings.Join(os.Args, ",")))
	case 1:
		fmt.Println(program.Use)
	default:
		switch strings.ToLower(options[1]) {
		case "category":
			if len(options) > 1 {
				return category(options[2])
			}
			return 2
		case "id":
			if len(options) > 1 {
				return category(options[2])
			}
			return 2
		case "name":
			namespace, err := strconv.Atoi(options[2])
			if err != nil {
				panic(fmt.Errorf("%s", options[2]))
			}
			return name(namespace)
		}
	}
	return 2
}
