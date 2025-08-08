package main

import (
	"golang.org/x/net/html"
	"os"
)

func main() {
	filePath := os.Args[1]

	fileBytes, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	links := make(map[string]string)

	doc, err := html.Parse(fileBytes)
	if err != nil {
		panic(err)
	}
	for n := range doc.Descendants() {
		if n.Type == html.ElementNode && n.Data == "a" {
			links[n.Attr[0].Val] = n.FirstChild.Data
		}
	}

	for k, v := range links {
		println(k, v)
	}

}
