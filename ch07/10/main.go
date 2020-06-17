package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println(IsPalindrome(myByte([]byte("thissiht"))))
	fmt.Println(IsPalindrome(myByte([]byte("thisAsiht"))))
	fmt.Println(IsPalindrome(myByte([]byte("thissiHt"))))
}

type myByte []byte

func (b myByte) Len() int           { return len(b) }
func (b myByte) Less(i, j int) bool { return b[i] < b[j] }
func (b myByte) Swap(i, j int)      { b[i], b[j] = b[j], b[i] }

func IsPalindrome(s sort.Interface) bool {
	for i := 0; i < s.Len()/2; i++ {
		if s.Less(i, s.Len()-i-1) || s.Less(s.Len()-i-1, i) {
			return false
		}
	}

	return true
}
