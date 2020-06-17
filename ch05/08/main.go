package main

import (
	"fmt"
	"os"

	"golang.org/x/net/html"
)

func main() {
	doc, _ := html.Parse(os.Stdin)
	fmt.Println(ElementByID(doc, "page"))
}

func ElementByID(doc *html.Node, id string) *html.Node {
	if doc == nil {
		return nil
	}

	pre := func(n *html.Node) bool {
		if n.Type != html.ElementNode {
			return true
		}

		for _, a := range n.Attr {
			if a.Key == "id" && a.Val == id {
				return false
			}
		}

		return true
	}

	return forEachNode(doc, pre, nil)
}

func forEachNode(n *html.Node, pre, post func(n *html.Node) bool) *html.Node {
	if pre != nil {
		if !pre(n) {
			return n
		}
	}

	for c := n.FirstChild; c != nil; c = c.NextSibling {
		ret := forEachNode(c, pre, post)
		if ret != nil {
			return ret
		}
	}

	if post != nil {
		if !post(n) {
			return n
		}
	}

	return nil
}
