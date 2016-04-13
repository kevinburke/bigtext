package main

import (
	"log"

	"github.com/kevinburke/bigtext"
)

func main() {
	d := bigtext.Client{
		Name:    "bigtext",
		LogoURL: "http://icons.iconarchive.com/icons/xenatt/the-circle/512/App-Terminal-icon.png",
		OpenURL: "https://google.com",
	}
	log.Fatal(d.Display("why hello there"))
}
