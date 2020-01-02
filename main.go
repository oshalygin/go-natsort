package main

import (
	"fmt"

	"github.com/facette/natsort"
)

func main() {
	list := []string{
		"Span",
		"Span 1",
		"Hello Boss",
		"Span 3",
		"Assessment",
		"Foundation 43",
		"Foundation 11",
		"Lolz 1",
		"Lolz",
		"Winning",
		"Hello World",
		"Foundation 2",
		"Foundation 34",
		"Foundation 1",
	}

	// This mutates the underlying array
	natsort.Sort(list)
	fmt.Println(list)
}
