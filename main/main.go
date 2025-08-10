package main

import (
	"fmt"
	"github.com/lmihailovic/link"
	"golang.org/x/net/html"
	"os"
)

func main() {
	htmlPath := os.Args[1]

	htmlBytes, err := os.Open(htmlPath)
	if err != nil {
		panic(err)
	}

	doc, err := html.Parse(htmlBytes)
	if err != nil {
		panic(err)
	}

	links := link.Parse(doc)

	for k, v := range links {
		fmt.Printf("\npath: %v\ntext: %v\n", k, v)
	}
}
