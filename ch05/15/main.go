package main

import "fmt"

func main() {
	fmt.Println(min2(3, 1, 2, 4))
	fmt.Println(min(3))
	fmt.Println(max2(3, 1, 2, 4))
	fmt.Println(max2(2))
	fmt.Println(max())
}

func min(numbers ...int) int {
	if len(numbers) == 0 {
		panic("You should pass at least one parameter.")
	}

	return min2(numbers[0], numbers[1:]...)
}

func min2(first int, numbers ...int) int {
	min := first
	for _, n := range numbers {
		if n < min {
			min = n
		}
	}

	return min
}

func max(numbers ...int) int {
	if len(numbers) == 0 {
		panic("You should pass at least one parameter.")
	}

	return max2(numbers[0], numbers[1:]...)
}

func max2(first int, numbers ...int) int {
	max := first
	for _, n := range numbers {
		if n > max {
			max = n
		}
	}

	return max
}
