package bigtext

import "fmt"

func ExampleDisplay() {
	err := Display("Hello World")
	fmt.Println(err) // nil
}
