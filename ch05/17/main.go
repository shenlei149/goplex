package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, _ := html.Parse(os.Stdin)
	images := elementsByTagName(doc, "img")
	for _, node := range images {
		fmt.Println(node)
	}

	headings := elementsByTagName(doc, "h1", "h2", "h3", "h4")
	for _, node := range headings {
		fmt.Println(node)
	}
}

func elementsByTagName(doc *html.Node, name ...string) []*html.Node {
	if doc == nil {
		return nil
	}

	var nodes []*html.Node
	if doc.Type == html.ElementNode {
		for _, n := range name {
			if doc.Data == n {
				nodes = append(nodes, doc)
				break
			}
		}
	}

	nodes = append(nodes, elementsByTagName(doc.FirstChild, name...)...)
	nodes = append(nodes, elementsByTagName(doc.NextSibling, name...)...)
	return nodes
}
