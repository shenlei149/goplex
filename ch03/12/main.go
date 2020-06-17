package main

import "fmt"

func main() {
	fmt.Printf("%t\n", compare("compare", "compare"))
	fmt.Printf("%t\n", compare("compare", "acermp"))
	fmt.Printf("%t\n", compare("compare", "acermop"))
	fmt.Printf("%t\n", compare("compare,世界", "acermop"))
	fmt.Printf("%t\n", compare("compare,世界", "ac界er世mop,"))
}

func compare(s1, s2 string) bool {
	counts := make(map[rune]int)
	for _, r := range s1 {
		counts[r]++
	}

	for _, r := range s2 {
		counts[r]--
	}

	for _, v := range counts {
		if v != 0 {
			return false
		}
	}

	return true
}
