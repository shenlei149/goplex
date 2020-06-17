package main

import (
	"fmt"
	"log"
	"os"
	"os/exec"
	"strings"
)

func main() {
	if len(os.Args) < 2 {
		fmt.Println("Usage exe your_wildcard")
	}

	packages := parent(listPackage(os.Args[1:]))
	for _, p := range packages {
		fmt.Println(p)
	}
}

func prependString(x []string, y string) []string {
	x = append(x, "")
	copy(x[1:], x)
	x[0] = y
	return x
}

func listPackage(args []string) []string {
	args = prependString(args, "list")
	cmd := exec.Command("go", args...)
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	return strings.Fields(string(output))
}

func parent(packages []string) []string {
	cmd := exec.Command("go", "list", "-f", "{{.ImportPath}} -> {{join .Imports \" \"}}", "...")
	output, err := cmd.Output()
	if err != nil {
		log.Fatal(err)
	}

	lines := strings.Split(strings.TrimSpace(string(output)), "\n")
	mapping := make(map[string]map[string]bool)
	for _, line := range lines {
		pair := strings.Split(line, " -> ")
		mapping[pair[0]] = make(map[string]bool)
		for _, p := range strings.Fields(pair[1]) {
			mapping[pair[0]][p] = true
		}
	}

	result := []string{}
	for k, v := range mapping {
		for _, p := range packages {
			if _, ok := v[p]; ok {
				result = append(result, k)
			}
		}
	}

	return result
}
