package main

import "fmt"

func main() {
	slice := []string{"1", "2", "2", "1", "1", "1", "2"}

	slice1 := slice[4:]
	fmt.Println(slice1)

	slice1 = dedup(slice1)
	fmt.Println(slice1)

	slice2 := slice[:5]
	fmt.Println(slice2)

	slice2 = dedup(slice2)
	fmt.Println(slice2)

	fmt.Println(slice)
}

func dedup(slice []string) []string {
	count := 0
	n := len(slice)
	for i := 0; i < n-1-count; {
		if slice[i] == slice[i+1] {
			count++
			for j := i + 1; j < n-1; j++ {
				slice[j] = slice[j+1]
			}
		} else {
			i++
		}
	}

	return slice[:n-count]
}
