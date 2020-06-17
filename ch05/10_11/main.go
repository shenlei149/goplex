package main

import (
	"fmt"
)

var prereqs = map[string][]string{
	"algorithms":     {"data structures"},
	"calculus":       {"linear algebra"},
	"linear algebra": {"calculus"},

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
	for i, course := range topoSort(prereqs) {
		fmt.Printf("%d:\t%s\n", i+1, course)
	}
}

func topoSort(m map[string][]string) []string {
	var order []string
	seen := make(map[string]bool)
	var visitAll func([]string)

	visitAll = func(items []string) {
		for _, item := range items {
			if !seen[item] {
				seen[item] = true
				visitAll(m[item])
				order = append(order, item)
			}
		}
	}

	var hasCycle func(string, []string) bool
	hasCycle = func(key string, path []string) bool {
		if path == nil {
			path = append(path, key)
		}

		for _, item := range m[key] {
			for _, p := range path {
				if p == item {
					return true
				}
			}

			path = append(path, item)
			return hasCycle(item, path)
		}

		return false
	}

	for key := range m {
		if hasCycle(key, nil) {
			fmt.Println("There is a cycle!")
			return nil
		}

		visitAll([]string{key})
	}

	return order
}
