package main

import "fmt"

func main() {
	array := [6]int{1, 2, 3, 4, 5, 6}
	fmt.Println(array)
	rotate(array[:], 2)
	fmt.Println(array)

	array2 := [7]int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(array2)
	rotate(array2[:], 3)
	fmt.Println(array2)
}

func rotate(s []int, n int) {
	count := 0
	for start := 0; count < len(s); start++ {
		current := start
		for ok := true; ok; ok = start != current {
			next := (current + n) % len(s)
			s[start], s[next] = s[next], s[start]

			current = next
			count++
		}
	}
}
