package pkg_test

import (
	"fmt"
	"strings"
	"testing"

	"github.com/gellel/emojipedia/pkg"
)

func Test(t *testing.T) {

	if p, err := pkg.Get(); err == nil {
		fmt.Println("Categories")
		for k, v := range *p.Main {
			fmt.Println(strings.ToLower(k), ":", v)
		}
		fmt.Println("-")
		fmt.Println("Subcategories")
		for k, v := range *p.Sub {
			fmt.Println(strings.ToLower(k), ":", v)
		}
	}
}
