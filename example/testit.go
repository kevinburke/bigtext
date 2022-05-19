package main

import (
	"log"

	"github.com/kevinburke/bigtext"
)

func main() {
	d := bigtext.Client{
		Name: "bigtext",
	}
	log.Fatal(d.Display("why hello there"))
}
