package main

import (
	"fmt"
	"net/http"
	"os"
	"strings"

	"golang.org/x/net/html"
)

func main() {
	fmt.Println(CountWordsAndImages(os.Args[1]))
}

func CountWordsAndImages(url string) (words, images int, err error) {
	resp, err := http.Get(url)
	if err != nil {
		return
	}

	doc, err := html.Parse(resp.Body)
	resp.Body.Close()
	if err != nil {
		err = fmt.Errorf("parsing HTML: %s", err)
		return
	}

	words, images = countWordsAndImages(doc)
	return
}

func countWordsAndImages(n *html.Node) (words, images int) {
	if n == nil {
		return
	}

	if n.Type == html.ElementNode {
		if n.Data == "img" {
			images++
		}
	}

	if n.Type == html.TextNode {
		content := strings.TrimSpace(n.Data)
		if len(content) > 0 {
			words += len(strings.Fields(content))
		}
	}

	w1, i1 := countWordsAndImages(n.FirstChild)
	w2, i2 := countWordsAndImages(n.NextSibling)

	return words + w1 + w2, images + i1 + i2
}
