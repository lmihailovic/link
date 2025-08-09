package main

import (
	"fmt"
	"golang.org/x/net/html"
	"os"
	"strings"
)

func parseLinkChildrenText(p *html.Node) string {
	var nodeText string

	for c := p.FirstChild; c != nil; c = c.NextSibling {
		if c.Type == html.TextNode {
			nodeText += c.Data
		} else {
			nodeText += parseLinkChildrenText(c)
		}
	}
	return nodeText
}

func ParseLinks(doc *html.Node) map[string]string {
	links := make(map[string]string)

	for n := range doc.Descendants() {
		var nodeText string
		if n.Type == html.ElementNode && n.Data == "a" {
			nodeText = parseLinkChildrenText(n)
			links[n.Attr[0].Val] = strings.TrimSpace(nodeText)
		}
	}
	return links
}

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

	links := ParseLinks(doc)

	for k, v := range links {
		fmt.Printf("\npath: %v\ntext: %v\n", k, v)
	}
}
