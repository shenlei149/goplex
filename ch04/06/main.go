package main

import "fmt"

import "unicode"

func main() {
	str := "hello,    世界!    I'll   be back!"
	bytes := []byte(str)

	fmt.Println(str)
	fmt.Println(bytes)

	bytes = squashSpace(bytes)
	str = string(bytes)

	fmt.Println(str)
	fmt.Println(bytes)
}

func squashSpace(data []byte) []byte {
	count := 0
	n := len(data)
	for i := 0; i < n-1-count; {
		if data[i] == data[i+1] && unicode.IsSpace(rune(data[i])) {
			count++
			for j := i + 1; j < n-1; j++ {
				data[j] = data[j+1]
			}
		} else {
			i++
		}
	}

	return data[:n-count]
}
