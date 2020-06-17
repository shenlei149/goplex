package main

import (
	"fmt"
	"os"

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

// visit appends to links each link found in n and returns the result.
func visit(links []string, n *html.Node) []string {
	if n == nil {
		return links
	}

	if n.Type == html.ElementNode {
		if n.Data == "a" || n.Data == "link" {
			for _, l := range n.Attr {
				if l.Key == "href" {
					links = append(links, l.Val)
				}
			}
		}

		if n.Data == "script" || n.Data == "img" {
			for _, l := range n.Attr {
				if l.Key == "src" {
					links = append(links, l.Val)
				}
			}
		}
	}

	links = visit(links, n.FirstChild)
	return visit(links, n.NextSibling)
}
