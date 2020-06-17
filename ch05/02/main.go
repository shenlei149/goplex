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

	counts := make(map[string]int)
	visit(counts, doc)
	for k, v := range counts {
		fmt.Println(k, v)
	}
}

func visit(counts map[string]int, n *html.Node) {
	if n == nil {
		return
	}

	if n.Type == html.ElementNode {
		counts[n.Data]++
	}

	visit(counts, n.FirstChild)
	visit(counts, n.NextSibling)
}
