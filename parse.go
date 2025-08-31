package link

import (
	"golang.org/x/net/html"
	"strings"
)

// parseLinkChildrenText recursively extracts and concatenates text content from all child nodes of the provided HTML node.
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

// Parse extracts all hyperlink references and their associated text content from the provided HTML document.
// It returns a map where keys represent href attribute values, and values are the trimmed text content of the links.
func Parse(doc *html.Node) map[string]string {
	links := make(map[string]string)

	for n := range doc.Descendants() {
		var nodeText string
		if n.Type == html.ElementNode && n.Data == "a" {
			nodeText = parseLinkChildrenText(n)
			for _, attribute := range n.Attr {
				if attribute.Key == "href" {
					links[attribute.Val] = strings.TrimSpace(nodeText)
				}
			}
		}
	}
	return links
}
