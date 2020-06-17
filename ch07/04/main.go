package main

import (
	"fmt"
	"io"
	"os"

	"golang.org/x/net/html"
)

func main() {
	s := "<html><div><a href=\"https://golang.org/doc/\">Documents</a><div/><a href=\"https://golang.org/pkg/\">Packages</a></html>"
	doc, err := html.Parse(NewStringReader(s))
	if err != nil {
		fmt.Fprintf(os.Stderr, "findlinks: %v\n", err)
		os.Exit(1)
	}

	for _, link := range visit(nil, doc) {
		fmt.Println(link)
	}
}

func visit(links []string, n *html.Node) []string {
	if n.Type == html.ElementNode && n.Data == "a" {
		for _, a := range n.Attr {
			if a.Key == "href" {
				links = append(links, a.Val)
			}
		}
	}
	for c := n.FirstChild; c != nil; c = c.NextSibling {
		links = visit(links, c)
	}
	return links
}

type StringReader struct {
	s string
}

func (sr *StringReader) Read(p []byte) (n int, err error) {
	n = copy(p, sr.s)
	sr.s = sr.s[n:]
	if len(sr.s) == 0 {
		err = io.EOF
	}

	return
}

func NewStringReader(s string) *StringReader {
	return &StringReader{s}
}
