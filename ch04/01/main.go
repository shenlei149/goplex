package main

import (
	"crypto/sha256"
	"fmt"
)

var pc [256]byte

func init() {
	for i := range pc {
		pc[i] = pc[i/2] + byte(i&1)
	}
}

func main() {
	c1 := sha256.Sum256([]byte("x"))
	c2 := sha256.Sum256([]byte("X"))
	count := byte(0)
	for i := 0; i < sha256.Size; i++ {
		count += pc[c1[i]^c2[i]]
	}

	fmt.Println(count)
}
