package main

import (
	"context"
	"flag"
	"fmt"
	"log"

	"crawl"
)

// tokens is a counting semaphore used to
// enforce a limit of 20 concurrent requests.
var tokens = make(chan struct{}, 20)

func c(url string, depth int) links {
	fmt.Println(url)
	tokens <- struct{}{} // acquire a token
	ctx, _ := context.WithCancel(context.Background())
	list, err := crawl.Extract(url, ctx)
	<-tokens // release the token

	if err != nil {
		log.Print(err)
	}
	return links{list, depth + 1}
}

type links struct {
	link  []string
	depth int
}

func main() {
	var depth = flag.Int("depth", 3, "depth of crawling")
	flag.Parse()
	worklist := make(chan links)
	var n int // number of pending sends to worklist

	// Start with the command-line arguments.
	n++
	go func() { worklist <- links{flag.Args(), 0} }()

	// Crawl the web concurrently.
	seen := make(map[string]bool)
	for ; n > 0; n-- {
		list := <-worklist
		if list.depth > *depth {
			continue
		}

		for _, link := range list.link {
			if !seen[link] {
				seen[link] = true
				n++
				go func(link string, depth int) {
					worklist <- c(link, depth)
				}(link, list.depth)
			}
		}
	}
}
