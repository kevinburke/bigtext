package bigtext_test

import (
	"fmt"

	"github.com/kevinburke/bigtext"
)

func ExampleDisplay() {
	err := bigtext.Display("Hello World")
	fmt.Println(err) // nil
}
