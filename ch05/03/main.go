package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	doc, err := html.Parse(os.Stdin)
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}

	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {
	if n == nil || n.Data == "script" || n.Data == "style" {
		return links
	}

	if n.Type == html.TextNode {
		content := strings.TrimSpace(n.Data)
		if len(content) > 0 {
			links = append(links, content)
		}
	}

	links = visit(links, n.FirstChild)
	return visit(links, n.NextSibling)
}
