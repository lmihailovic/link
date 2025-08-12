package link

import (
	"golang.org/x/net/html"
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
