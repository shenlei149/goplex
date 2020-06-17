package main

import (
	"context"
	"fmt"
	"log"
	"os"
	"sync"

	"crawl"
)

func c(url string) []string {
	fmt.Println(url)
	list, err := crawl.Extract(url, ctx)
	if err != nil {
		log.Print(err)
	}
	return list
}

var done = make(chan struct{})

func cancelled() bool {
	select {
	case <-done:
		return true
	default:
		return false
	}
}

var ctx context.Context
var cancel context.CancelFunc

func main() {
	worklist := make(chan []string)  // lists of URLs, may have duplicates
	unseenLinks := make(chan string) // de-duplicated URLs

	ctx, cancel = context.WithCancel(context.Background())

	// Add command-line arguments to worklist.
	go func() { worklist <- os.Args[1:] }()

	go func() {
		os.Stdin.Read(make([]byte, 1)) // read a single byte
		close(done)
		cancel()
	}()

	var wg sync.WaitGroup

	// Create 20 crawler goroutines to fetch each unseen link.
	for i := 0; i < 20; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			for link := range unseenLinks {
				if !cancelled() {
					foundLinks := c(link)
					go func() { worklist <- foundLinks }()
				}
			}
		}()
	}

	go func() {
		wg.Wait()
		close(worklist)
	}()

	// The main goroutine de-duplicates worklist items
	// and sends the unseen ones to the crawlers.
	seen := make(map[string]bool)
	for list := range worklist {
		if cancelled() {
			close(unseenLinks)
			break
		} else {
			for _, link := range list {
				if !seen[link] {
					seen[link] = true
					unseenLinks <- link
				}
			}
		}
	}

	for range worklist {
	}
}
