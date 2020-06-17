package main

import (
	"fmt"
	"sort"

	"crawl"
)

var prereqs = map[string][]string{
	"algorithms": {"data structures"},
	"calculus":   {"linear algebra"},

	"compilers": {
		"data structures",
		"formal languages",
		"computer organization",
	},

	"data structures":       {"discrete math"},
	"databases":             {"data structures"},
	"discrete math":         {"intro to programming"},
	"formal languages":      {"discrete math"},
	"networks":              {"operating systems"},
	"operating systems":     {"data structures", "computer organization"},
	"programming languages": {"data structures", "computer organization"},
}

func main() {
	var keys []string
	for key := range prereqs {
		keys = append(keys, key)
	}

	sort.Strings(keys)

	crawl.BreadthFirst(func(key string) []string {
		fmt.Println(key)
		return prereqs[key]
	}, keys)
}
