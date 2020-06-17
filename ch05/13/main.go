package main

import (
	"context"
	"fmt"
	"log"
	"net/url"
	"os"
	"strings"

	"crawl"
)

func main() {
	crawl.BreadthFirst(c, os.Args[1:])
}

func c(urlStr string) []string {
	url, err := url.Parse(urlStr)
	if err != nil {
		log.Print(err)
		return nil
	}

	fmt.Println(urlStr)
	ctx, _ := context.WithCancel(context.Background())
	list, err := crawl.Extract(urlStr, ctx)
	if err != nil {
		log.Print(err)
	}

	var newList []string
	for _, u := range list {
		nu, _ := url.Parse(u)
		if strings.HasSuffix(nu.Host, url.Host) {
			newList = append(newList, u)
		}
	}

	return newList
}
