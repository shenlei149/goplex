package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	wordfreq := make(map[string]int)
	in := bufio.NewScanner(os.Stdin)
	in.Split(bufio.ScanWords)
	for in.Scan() {
		word := in.Text()
		wordfreq[word]++
	}

	for k, v := range wordfreq {
		fmt.Printf("%s\t%d\n", k, v)
	}
}
