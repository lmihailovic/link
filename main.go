package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
	"strings"
)

func parseLinks(doc *html.Node) map[string]string {
	links := make(map[string]string)

	for n := range doc.Descendants() {
		var nodeText string
		if n.Type == html.ElementNode && n.Data == "a" {
			nodeText = strings.TrimSpace(n.FirstChild.Data)
			links[n.Attr[0].Val] = nodeText
		}
	}
	return links
}

func main() {
	filePath := os.Args[1]

	fileBytes, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}

	doc, err := html.Parse(fileBytes)
	if err != nil {
		panic(err)
	}

	links := parseLinks(doc)

	for k, v := range links {
		fmt.Printf("\npath: %v\ntext: %v\n", k, v)
	}

}
