package main

import "fmt"

func main() {
	fmt.Println(f())
}

func f() (result int) {
	defer func() {
		recover()
		result = -1
	}()

	panic("")
}
